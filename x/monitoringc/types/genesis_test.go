package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "1"},
				},
				LaunchIDFromChannelIDList: []types.LaunchIDFromChannelID{
					{LaunchID: 0, ChannelID: "0"},
					{LaunchID: 1, ChannelID: "1"},
				},
				MonitoringHistoryList: []types.MonitoringHistory{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 1,
					},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated verifiedClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
					{
						LaunchID:  0,
						ClientIDs: []string{"1", "2"},
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated clientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0", "0"},
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated providerClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{
						LaunchID: 0,
						ClientID: "0",
					},
					{
						LaunchID: 0,
						ClientID: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated launchIDFromVerifiedClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{
						ClientID: "0",
						LaunchID: 0,
					},
					{
						ClientID: "0",
						LaunchID: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "provider client id without valid client id",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "3"},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			desc: "launch id from verified client id without valid client id",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDList: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDList: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				LaunchIDFromVerifiedClientIDList: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "1"},
					{LaunchID: 1, ClientID: "1"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			desc: "duplicated launchIDFromChannelID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				LaunchIDFromChannelIDList: []types.LaunchIDFromChannelID{
					{
						ChannelID: "0",
					},
					{
						ChannelID: "0",
					},
				},
			},
			valid: false,
		},
		{
			desc: "duplicated monitoringHistory",
			genState: &types.GenesisState{
				PortId: types.PortID,
				MonitoringHistoryList: []types.MonitoringHistory{
					{
						LaunchID: 0,
					},
					{
						LaunchID: 0,
					},
				},
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
