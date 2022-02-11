package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

var (
	launchID1        = uint64(0)
	launchID2        = uint64(1)
	noExistLaunchID  = uint64(2)
	addr1            = sample.Address()
	addr2            = sample.Address()
	vestingAddress   = sample.Address()
	genesisValidator = sample.GenesisValidator(launchID1, addr1)
	genesisChainID   = sample.GenesisChainID()

	// Those are samples we can use for each fields when they are not the one to test
	sampleChainList = []types.Chain{
		{
			LaunchID:       launchID1,
			GenesisChainID: genesisChainID,
		},
		{
			LaunchID:       launchID2,
			GenesisChainID: genesisChainID,
		},
	}
	sampleGenesisAccountList = []types.GenesisAccount{
		{
			LaunchID: launchID1,
			Address:  addr1,
		},
		{
			LaunchID: launchID1,
			Address:  addr2,
		},
		{
			LaunchID: launchID2,
			Address:  addr1,
		},
		{
			LaunchID: launchID2,
			Address:  addr2,
		},
	}
	sampleVestingAccountList = []types.VestingAccount{
		{
			LaunchID: launchID1,
			Address:  vestingAddress,
		},
		{
			LaunchID: launchID2,
			Address:  vestingAddress,
		},
	}
	sampleGenesisValidatorList = []types.GenesisValidator{genesisValidator}
	sampleRequestList          = []types.Request{
		{
			LaunchID:  launchID1,
			RequestID: 0,
		},
		{
			LaunchID:  launchID1,
			RequestID: 1,
		},
	}
	sampleRequestCounterList = []types.RequestCounter{
		{
			LaunchID: launchID1,
			Counter:  10,
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
				ChainCounter:         10,
				GenesisAccountList:   sampleGenesisAccountList,
				VestingAccountList:   sampleVestingAccountList,
				GenesisValidatorList: sampleGenesisValidatorList,
				RequestList:          sampleRequestList,
				RequestCounterList:   sampleRequestCounterList,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			shouldBeValid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
		{
			desc: "invalid chain",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
					{
						LaunchID:       launchID1,
						GenesisChainID: "invalid_chain_id",
					},
				},
				ChainCounter: 10,
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated chains",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
					{
						LaunchID:       launchID1,
						GenesisChainID: genesisChainID,
					},
					{
						LaunchID:       launchID1,
						GenesisChainID: genesisChainID,
					},
				},
				ChainCounter: 10,
			},
			shouldBeValid: false,
		},
		{
			desc: "chain with a chain id number above the chain counter",
			genState: &types.GenesisState{
				ChainList: []types.Chain{
					{
						LaunchID:       12,
						GenesisChainID: genesisChainID,
					},
				},
				ChainCounter: 10,
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated accounts",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				GenesisAccountList: []types.GenesisAccount{
					{
						LaunchID: launchID1,
						Address:  addr1,
					},
					{
						LaunchID: launchID1,
						Address:  addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "account not associated with chain",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				GenesisAccountList: []types.GenesisAccount{
					{
						LaunchID: noExistLaunchID,
						Address:  addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated vesting accounts",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				VestingAccountList: []types.VestingAccount{
					{
						LaunchID: launchID1,
						Address:  vestingAddress,
					},
					{
						LaunchID: launchID1,
						Address:  vestingAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "vesting account not associated with chain",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				VestingAccountList: []types.VestingAccount{
					{
						LaunchID: noExistLaunchID,
						Address:  vestingAddress,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "address as genesis account and vesting account",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				GenesisAccountList: []types.GenesisAccount{
					{
						LaunchID: launchID1,
						Address:  addr1,
					},
				},
				VestingAccountList: []types.VestingAccount{
					{
						LaunchID: launchID1,
						Address:  addr1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "genesis validator not associated to a chain",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				GenesisValidatorList: []types.GenesisValidator{
					sample.GenesisValidator(noExistLaunchID, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated genesis validator",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				GenesisValidatorList: []types.GenesisValidator{
					sample.GenesisValidator(launchID1, addr1),
					sample.GenesisValidator(launchID1, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "validator address not associated to a chain",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				GenesisValidatorList: []types.GenesisValidator{
					sample.GenesisValidator(noExistLaunchID, addr1),
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated requests",
			genState: &types.GenesisState{
				ChainList:          sampleChainList,
				ChainCounter:       10,
				RequestCounterList: sampleRequestCounterList,
				RequestList: []types.Request{
					{
						LaunchID:  launchID1,
						RequestID: 0,
					},
					{
						LaunchID:  launchID1,
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
				ChainCounter:       10,
				RequestCounterList: sampleRequestCounterList,
				RequestList: []types.Request{
					{
						LaunchID:  noExistLaunchID,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request while no request count for the chain",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				RequestCounterList: []types.RequestCounter{
					{
						LaunchID: launchID2,
						Counter:  1,
					},
				},
				RequestList: []types.Request{
					{
						LaunchID:  launchID1,
						RequestID: 0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "duplicated request counter",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				RequestCounterList: []types.RequestCounter{
					{
						LaunchID: launchID1,
						Counter:  0,
					},
					{
						LaunchID: launchID1,
						Counter:  1,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request counter not associated with chain",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				RequestCounterList: []types.RequestCounter{
					{
						LaunchID: noExistLaunchID,
						Counter:  0,
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "request counter below a request id",
			genState: &types.GenesisState{
				ChainList:    sampleChainList,
				ChainCounter: 10,
				RequestCounterList: []types.RequestCounter{
					{
						LaunchID: launchID1,
						Counter:  5,
					},
				},
				RequestList: []types.Request{
					{
						LaunchID:  launchID1,
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
			if !tc.shouldBeValid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			launchIDMap := make(map[uint64]struct{})
			for _, elem := range tc.genState.ChainList {
				launchIDMap[elem.LaunchID] = struct{}{}
			}

			for _, acc := range tc.genState.RequestList {
				// check if the chain exist for requests
				_, ok := launchIDMap[acc.LaunchID]
				require.True(t, ok)
			}

			for _, acc := range tc.genState.GenesisValidatorList {
				// check if the chain exist for validators
				_, ok := launchIDMap[acc.LaunchID]
				require.True(t, ok)
			}

			for _, acc := range tc.genState.GenesisAccountList {
				// check if the chain exist for genesis accounts
				_, ok := launchIDMap[acc.LaunchID]
				require.True(t, ok)
			}

			for _, acc := range tc.genState.VestingAccountList {
				// check if the chain exist for vesting accounts
				_, ok := launchIDMap[acc.LaunchID]
				require.True(t, ok)
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
