package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc string
		genState   *types.GenesisState
		shouldBeValid bool
	}{
		{
			desc: "default is valid",
			genState: types.DefaultGenesis(),
			shouldBeValid: true,
		},
		{
			desc: "duplicated chains",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: "foo",
					},
					{
						ChainID: "foo",
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
						ChainID: "foo",
						Address: "bar",
					},
					{
						ChainID: "foo",
						Address: "bar",
					},
				},
			},
			shouldBeValid: false,
		},
		{
			desc: "account associated with chain",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: "foo",
					},
				},
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: "foo",
						Address: "bar",
					},
				},
			},
			shouldBeValid: true,
		},
		{
			desc: "account not associated with chain",
			genState: &types.GenesisState{
				ChainList: []*types.Chain{
					{
						ChainID: "foo",
					},
				},
				GenesisAccountList: []*types.GenesisAccount{
					{
						ChainID: "nonexistent",
						Address: "bar",
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