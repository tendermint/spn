package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

var (
	chainID1, chainName1             = sample.ChainID(0)
	chainID2, chainName2             = sample.ChainID(0)
	noExistChainID, noExistChainName = sample.ChainID(0)
	addr1                            = sample.AccAddress()
	addr2                            = sample.AccAddress()
	vestedAddress                    = sample.AccAddress()
	genesisValidator = sample.GenesisValidator(chainID1, addr1)

	// Those are samples we can use for each fields when they are not the one to test
	sampleChainList = []*types.Chain{
		{
			ChainID: chainID1,
		},
		{
			ChainID: chainID2,
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
			ChainID: chainID1,
			Address: addr1,
		},
		{
			ChainID: chainID1,
			Address: addr2,
		},
		{
			ChainID: chainID2,
			Address: addr1,
		},
		{
			ChainID: chainID2,
			Address: addr2,
		},
	}
	sampleVestedAccountList = []*types.VestedAccount{
		{
			ChainID: chainID1,
			Address: vestedAddress,
		},
		{
			ChainID: chainID2,
			Address: vestedAddress,
		},
	}
	sampleGenesisValidatorList = []*types.GenesisValidator{genesisValidator}
	sampleRequestList = []*types.Request{
		{
			ChainID:   chainID1,
			RequestID: 0,
		},
		{
			ChainID:   chainID1,
			RequestID: 1,
		},
	}
	sampleRequestCountList = []*types.RequestCount{
		{
			ChainID: chainID1,
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
				GenesisValidatorList:  sampleGenesisValidatorList,
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
						ChainID: chainID1,
					},
					{
						ChainID: chainID1,
					},
				},
				ChainNameCountList: sampleChainNameCountList,
			},
			shouldBeValid: false,
		},
		{
			desc: "invalid chain id",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: "foo",
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated chain name counts",
			genState: &types.GenesisState{
				ChainList: sampleChainList,
				ChainNameCountList: []*types.ChainNameCount{
					{
						ChainName: chainName1,
						Count:     10,
					},
					{
						ChainName: chainName1,
						Count:     10,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "chain with a name not associated to a chain name counts",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: chainID1,
					},
				},
				ChainNameCountList: []*types.ChainNameCount{
					{
						ChainName: chainName2,
						Count:     10,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "chain with a chain id number above the chain name counts",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: types.ChainIDFromChainName(chainName1, 20),
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
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: chainID1,
						Address: addr1,
					},
					{
						ChainID: chainID1,
						Address: addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "account not associated with chain",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: noExistChainID,
						Address: addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated vested accounts",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				VestedAccountList: []*types.VestedAccount{
					{
						ChainID: chainID1,
						Address: vestedAddress,
					},
					{
						ChainID: chainID1,
						Address: vestedAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "vested account not associated with chain",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				VestedAccountList: []*types.VestedAccount{
					{
						ChainID: noExistChainID,
						Address: vestedAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "address as genesis account and vested account",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: chainID1,
						Address: addr1,
					},
				},
				VestedAccountList: []*types.VestedAccount{
					{
						ChainID: chainID1,
						Address: addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "genesis validator not associated to a chain",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				GenesisValidatorList: []*types.GenesisValidator{
					sample.GenesisValidator(noExistChainID, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated genesis validator",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				GenesisValidatorList: []*types.GenesisValidator{
					sample.GenesisValidator(chainID1, addr1),
					sample.GenesisValidator(chainID1, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "validator address not associated to a chain",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				GenesisValidatorList: []*types.GenesisValidator{
					sample.GenesisValidator(noExistChainID, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated requests",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				RequestCountList:   sampleRequestCountList,
				RequestList: []*types.Request{
					{
						ChainID:   chainID1,
						RequestID: 0,
					},
					{
						ChainID:   chainID1,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request not associated with chain",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				RequestCountList:   sampleRequestCountList,
				RequestList: []*types.Request{
					{
						ChainID:   noExistChainID,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request while no request count for the chain",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainID2,
						Count:   1,
					},
				},
				RequestList: []*types.Request{
					{
						ChainID:   chainID1,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated request count",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainID1,
						Count:   0,
					},
					{
						ChainID: chainID1,
						Count:   1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request count not associated with chain",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				RequestCountList: []*types.RequestCount{
					{
						ChainID: noExistChainID,
						Count:   0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request count below a request id",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainNameCountList: sampleChainNameCountList,
				RequestCountList: []*types.RequestCount{
					{
						ChainID: chainID1,
						Count:   5,
					},
				},
				RequestList: []*types.Request{
					{
						ChainID:   chainID1,
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
