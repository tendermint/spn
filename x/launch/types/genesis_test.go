package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

var (
	chainID1         = uint64(0)
	chainID2         = uint64(1)
	noExistChainID   = uint64(2)
	addr1            = sample.AccAddress()
	addr2            = sample.AccAddress()
	vestingAddress   = sample.AccAddress()
	genesisValidator = *sample.GenesisValidator(chainID1, addr1)
	genesisChainID   = sample.GenesisChainID()

	// Those are samples we can use for each fields when they are not the one to test
	sampleChainList = []types.Chain{
		{
			Id:             chainID1,
			GenesisChainID: genesisChainID,
		},
		{
			Id:             chainID2,
			GenesisChainID: genesisChainID,
		},
	}
	sampleGenesisAccountList = []types.GenesisAccount{
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
	sampleVestingAccountList = []types.VestingAccount{
		{
			ChainID: chainID1,
			Address: vestingAddress,
		},
		{
			ChainID: chainID2,
			Address: vestingAddress,
		},
	}
	sampleGenesisValidatorList = []types.GenesisValidator{genesisValidator}
	sampleRequestList          = []types.Request{
		{
			ChainID:   chainID1,
			RequestID: 0,
		},
		{
			ChainID:   chainID1,
			RequestID: 1,
		},
	}
	sampleRequestCountList = []types.RequestCount{
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
				ChainList:            sampleChainList,
				ChainCount:           10,
				GenesisAccountList:   sampleGenesisAccountList,
				VestingAccountList:   sampleVestingAccountList,
				GenesisValidatorList: sampleGenesisValidatorList,
				RequestList:          sampleRequestList,
				RequestCountList:     sampleRequestCountList,
			},
			shouldBeValid: true,
		},
		{
			desc: "invalid chain",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
					{
						Id:             chainID1,
						GenesisChainID: "invalid_chain_id",
					},
				},
				ChainCount: 10,
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated chains",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
					{
						Id:             chainID1,
						GenesisChainID: genesisChainID,
					},
					{
						Id:             chainID1,
						GenesisChainID: genesisChainID,
					},
				},
				ChainCount: 10,
			},
			shouldBeValid: false,
		},
		{
			desc: "chain with a chain id number above the chain count",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
					{
						Id:             12,
						GenesisChainID: genesisChainID,
					},
				},
				ChainCount: 10,
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated accounts",
			genState: &types.GenesisState{
				ChainList:  sampleChainList,
				ChainCount: 10,
				GenesisAccountList: []types.GenesisAccount{
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
				ChainList:  sampleChainList,
				ChainCount: 10,
				GenesisAccountList: []types.GenesisAccount{
					{
						ChainID: noExistChainID,
						Address: addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated vesting accounts",
			genState: &types.GenesisState{
				ChainList:  sampleChainList,
				ChainCount: 10,
				VestingAccountList: []types.VestingAccount{
					{
						ChainID: chainID1,
						Address: vestingAddress,
					},
					{
						ChainID: chainID1,
						Address: vestingAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "vesting account not associated with chain",
			genState: &types.GenesisState{
				ChainList:  sampleChainList,
				ChainCount: 10,
				VestingAccountList: []types.VestingAccount{
					{
						ChainID: noExistChainID,
						Address: vestingAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "address as genesis account and vesting account",
			genState: &types.GenesisState{
				ChainList:  sampleChainList,
				ChainCount: 10,
				GenesisAccountList: []types.GenesisAccount{
					{
						ChainID: chainID1,
						Address: addr1,
					},
				},
				VestingAccountList: []types.VestingAccount{
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
				ChainList:  sampleChainList,
				ChainCount: 10,
				GenesisValidatorList: []types.GenesisValidator{
					*sample.GenesisValidator(noExistChainID, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated genesis validator",
			genState: &types.GenesisState{
				ChainList:  sampleChainList,
				ChainCount: 10,
				GenesisValidatorList: []types.GenesisValidator{
					*sample.GenesisValidator(chainID1, addr1),
					*sample.GenesisValidator(chainID1, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "validator address not associated to a chain",
			genState: &types.GenesisState{
				ChainList:  sampleChainList,
				ChainCount: 10,
				GenesisValidatorList: []types.GenesisValidator{
					*sample.GenesisValidator(noExistChainID, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated requests",
			genState: &types.GenesisState{
				ChainList:        sampleChainList,
				ChainCount:       10,
				RequestCountList: sampleRequestCountList,
				RequestList: []types.Request{
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
				ChainList:        sampleChainList,
				ChainCount:       10,
				RequestCountList: sampleRequestCountList,
				RequestList: []types.Request{
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
				ChainList:  sampleChainList,
				ChainCount: 10,
				RequestCountList: []types.RequestCount{
					{
						ChainID: chainID2,
						Count:   1,
					},
				},
				RequestList: []types.Request{
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
				ChainList:  sampleChainList,
				ChainCount: 10,
				RequestCountList: []types.RequestCount{
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
				ChainList:  sampleChainList,
				ChainCount: 10,
				RequestCountList: []types.RequestCount{
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
				ChainList:  sampleChainList,
				ChainCount: 10,
				RequestCountList: []types.RequestCount{
					{
						ChainID: chainID1,
						Count:   5,
					},
				},
				RequestList: []types.Request{
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

func TestGenesisState_ValidateParams(t *testing.T) {
	for _, tc := range []struct {
		desc          string
		genState      types.GenesisState
		shouldBeValid bool
	}{
		{
			desc: "max launch time above the max parametrable launch time",
			genState: types.GenesisState{
				Params: types.NewParams(types.DefaultMinLaunchTime, types.MaxParametrableLaunchTime+1),
			},
			shouldBeValid: false,
		},
		{
			desc: "min launch time above max launch time",
			genState: types.GenesisState{
				Params: types.NewParams(types.DefaultMinLaunchTime+1, types.DefaultMinLaunchTime),
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
