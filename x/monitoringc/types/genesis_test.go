package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/x/monitoringc/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		name     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			name:     "should allow valid default genesis",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			name: "should allow valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDs: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDs: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				LaunchIDsFromVerifiedClientID: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "1"},
				},
				LaunchIDsFromChannelID: []types.LaunchIDFromChannelID{
					{LaunchID: 0, ChannelID: "0"},
					{LaunchID: 1, ChannelID: "1"},
				},
				MonitoringHistories: []types.MonitoringHistory{
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
			name: "should prevent invalid portID",
			genState: &types.GenesisState{
				PortId: "",
			},
			valid: false,
		},
		{
			name: "should prevent duplicated verifiedClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDs: []types.VerifiedClientID{
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
			name: "should prevent duplicated clientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDs: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0", "0"},
					},
				},
			},
			valid: false,
		},
		{
			name: "should prevent duplicated providerClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDs: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
				},
				ProviderClientIDs: []types.ProviderClientID{
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
			name: "should prevent duplicated launchIDFromVerifiedClientID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDs: []types.VerifiedClientID{
					{
						LaunchID:  0,
						ClientIDs: []string{"0"},
					},
				},
				LaunchIDsFromVerifiedClientID: []types.LaunchIDFromVerifiedClientID{
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
			name: "should prevent provider client id without valid client id",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDs: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDs: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "3"},
				},
				LaunchIDsFromVerifiedClientID: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			name: "should prevent launch id from verified client id without valid client id",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VerifiedClientIDs: []types.VerifiedClientID{
					{LaunchID: 0, ClientIDs: []string{"0"}},
					{LaunchID: 1, ClientIDs: []string{"1", "2"}},
				},
				ProviderClientIDs: []types.ProviderClientID{
					{LaunchID: 0, ClientID: "0"},
					{LaunchID: 1, ClientID: "2"},
				},
				LaunchIDsFromVerifiedClientID: []types.LaunchIDFromVerifiedClientID{
					{LaunchID: 0, ClientID: "1"},
					{LaunchID: 1, ClientID: "1"},
				},
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: false,
		},
		{
			name: "should prevent duplicated launchIDFromChannelID",
			genState: &types.GenesisState{
				PortId: types.PortID,
				LaunchIDsFromChannelID: []types.LaunchIDFromChannelID{
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
			name: "should prevent duplicated monitoringHistory",
			genState: &types.GenesisState{
				PortId: types.PortID,
				MonitoringHistories: []types.MonitoringHistory{
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
		t.Run(tc.name, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
