package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

var (
	chainId1, chainName1 = sample.ChainID(0)
	chainId2, chainName2 = sample.ChainID(0)
	addr1                = sample.AccAddress()
	addr2                = sample.AccAddress()
	vestedAddress        = sample.AccAddress()

	// Those are samples we can use for each fields when they are not the one to test
	sampleChainList = []*types.Chain{
		{
			ChainID: chainId1,
		},
		{
			ChainID: chainId2,
		},
	}
	sampleChainNameCountList = []*types.ChainNameCount{
		{
			ChainName: chainName1,
			Count:     10,
		},
		{
			ChainName: chainName2,
			Count:     10,
		},
	}
	sampleGenesisAccountList = []*types.GenesisAccount{
		{
			ChainID: chainId1,
			Address: addr1,
		},
		{
			ChainID: chainId1,
			Address: addr2,
		},
		{
			ChainID: chainId2,
			Address: addr1,
		},
		{
			ChainID: chainId2,
			Address: addr2,
		},
	}
	sampleVestedAccountList = []*types.VestedAccount{
		{
			ChainID: chainId1,
			Address: vestedAddress,
		},
		{
			ChainID: chainId2,
			Address: vestedAddress,
		},
	}
	sampleRequestList = []*types.Request{
		{
			ChainID:   chainId1,
			RequestID: 0,
		},
		{
			ChainID:   chainId1,
			RequestID: 1,
		},
	}
	sampleRequestCountList = []*types.RequestCount{
		{
			ChainID: chainId1,
			Count:   10,
		},
	}
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc          string
		genState      *types.GenesisState
		shouldBeValid bool
	}{
		{
			desc:          "default is valid",
			genState:      types.DefaultGenesis(),
			shouldBeValid: true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				GenesisAccountList: sampleGenesisAccountList,
				VestedAccountList:  sampleVestedAccountList,
				RequestList:        sampleRequestList,
				RequestCountList:   sampleRequestCountList,
			},
			shouldBeValid: true,
		},
		{
			desc: "duplicated chains",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
					{
						ChainID: chainId1,
					},
				},
				ChainNameCountList: []*types.ChainNameCount{
					{
						ChainName: chainName1,
						Count:     10,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated accounts",
			genState: &types.GenesisState{
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: chainId1,
						Address: addr1,
					},
					{
						ChainID: chainId1,
						Address: addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "account not associated with chain",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: chainId2,
						Address: addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated vested accounts",
			genState: &types.GenesisState{
				VestedAccountList: []*types.VestedAccount{
					{
						ChainID: chainId1,
						Address: vestedAddress,
					},
					{
						ChainID: chainId1,
						Address: vestedAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "vested account not associated with chain",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				VestedAccountList: []*types.VestedAccount{
					{
						ChainID: chainId2,
						Address: vestedAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "address as genesis account and vested account",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: chainId1,
						Address: addr1,
					},
				},
				VestedAccountList: []*types.VestedAccount{
					{
						ChainID: chainId1,
						Address: addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated requests",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainId1,
						Count:   2,
					},
				},
				RequestList: []*types.Request{
					{
						ChainID:   chainId1,
						RequestID: 0,
					},
					{
						ChainID:   chainId1,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request not associated with chain",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainId1,
						Count:   1,
					},
				},
				RequestList: []*types.Request{
					{
						ChainID:   chainId2,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request while no request count for the chain",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
					{
						ChainID: chainId2,
					},
				},
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainId2,
						Count:   1,
					},
				},
				RequestList: []*types.Request{
					{
						ChainID:   chainId1,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated request count",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainId1,
						Count:   0,
					},
					{
						ChainID: chainId1,
						Count:   1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request count not associated with chain",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainId2,
						Count:   0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request count below a request id",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainId1,
					},
				},
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainId2,
						Count:   5,
					},
				},
				RequestList: []*types.Request{
					{
						ChainID:   chainId1,
						RequestID: 10,
					},
				},
			},
			shouldBeValid: false,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.shouldBeValid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
