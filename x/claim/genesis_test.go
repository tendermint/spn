package claim_test

import (
	"github.com/tendermint/spn/testutil/sample"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/claim"
	"github.com/tendermint/spn/x/claim/types"
)

var (
	r *rand.Rand
)

// initialize random generator
func init() {
	s := rand.NewSource(1)
	r = rand.New(s)
}

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ClaimRecordList: []types.ClaimRecord{
			{
				Address: sample.Address(r),
			},
			{
				Address: sample.Address(r),
			},
		},
		MissionList: []types.Mission{
			{
				ID: 0,
			},
			{
				ID: 1,
			},
		},
		MissionCount:  2,
		AirdropSupply: sample.Coin(r),
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ClaimKeeper(t)
	claim.InitGenesis(ctx, *k, genesisState)
	got := claim.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ClaimRecordList, got.ClaimRecordList)
	require.ElementsMatch(t, genesisState.MissionList, got.MissionList)
	require.Equal(t, genesisState.MissionCount, got.MissionCount)
	require.Equal(t, genesisState.AirdropSupply, got.AirdropSupply)
	// this line is used by starport scaffolding # genesis/test/assert
}
