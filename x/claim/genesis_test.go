package claim_test

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim"
	"github.com/tendermint/spn/x/claim/types"
)

var r *rand.Rand

// initialize random generator
func init() {
	s := rand.NewSource(1)
	r = rand.New(s)
}

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClaimRecords: []types.ClaimRecord{
			{
				Address: sample.Address(r),
			},
			{
				Address: sample.Address(r),
			},
		},
		Missions: []types.Mission{
			{
				MissionID: 0,
			},
			{
				MissionID: 1,
			},
		},
		AirdropSupply: sample.Coin(r),
		InitialClaim: &types.InitialClaim{
			Enabled:   true,
			MissionID: 35,
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	ctx, tk, _ := testkeeper.NewTestSetup(t)
	claim.InitGenesis(ctx, *tk.ClaimKeeper, genesisState)
	got := claim.ExportGenesis(ctx, *tk.ClaimKeeper)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClaimRecords, got.ClaimRecords)
	require.ElementsMatch(t, genesisState.Missions, got.Missions)
	require.Equal(t, genesisState.AirdropSupply, got.AirdropSupply)
	require.Equal(t, genesisState.InitialClaim, got.InitialClaim)
	// this line is used by starport scaffolding # genesis/test/assert
}
