package simulation_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	profilekeeper "github.com/tendermint/spn/x/profile/keeper"
	"math/rand"
	"testing"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	simcampaign "github.com/tendermint/spn/x/campaign/simulation"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

// populateCoordinators populates the profile keeper with some coordinators from simulation accounts
func populateCoordinators(
	t *testing.T,
	r *rand.Rand,
	ctx sdk.Context,
	pk profilekeeper.Keeper,
	accs []simtypes.Account,
	coordNb int,
) (coordIDs []uint64) {
	require.LessOrEqual(t, coordNb, len(accs))

	// Create coordinator from random sim accounts
	created := make(map[int]struct{})
	for i := 0; i < coordNb; i++ {
		// Find randomly an index never used
		var nb int
		for {
			nb = r.Intn(coordNb)
			if _, ok := created[nb]; ok {
				continue
			} else {
				created[nb] = struct{}{}
				break
			}
		}

		coordID := pk.AppendCoordinator(ctx, profiletypes.Coordinator{
			Address:     accs[nb].Address.String(),
			Description: sample.CoordinatorDescription(),
		})

		coordIDs = append(coordIDs, coordID)
	}

	return
}

func TestGetCoordSimAccount(t *testing.T) {
	pk, ctx := testkeeper.Profile(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	// No coordinator account
	_, _, found := simcampaign.GetCoordSimAccount(r, ctx, pk, accs)
	require.False(t, found)

	// Find a coordinator
	populateCoordinators(t, r, ctx, *pk, accs, 10)
	acc, coordID, found := simcampaign.GetCoordSimAccount(r, ctx, pk, accs)
	require.True(t, found)
	require.Contains(t, accs, acc)
	require.True(t, pk.HasCoordinator(ctx, coordID))
}

func TestGetCoordSimAccountWithCampaignID(t *testing.T) {
	ck, _, pk, _, ctx := testkeeper.AllKeepers(t)
	r := sample.Rand()
	accs := sample.SimAccounts()

	// No campaign
	_, _, found := simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, false, false)
	require.False(t, found)

	coords := populateCoordinators(t, r, ctx, *pk, accs, 10)

	// Find a campaign
	camp := campaigntypes.NewCampaign(
		0,
		sample.AlphaString(5),
		coords[0],
		sample.Coins(),
		false,
	)
	camp.MainnetInitialized = true
	ck.AppendCampaign(ctx, camp)
	acc, id, found := simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, false, false)
	require.True(t, found)
	require.Contains(t, accs, acc)
	_, found = ck.GetCampaign(ctx, id)
	require.True(t, found)
	require.EqualValues(t, id, camp.Id)

	// Find a campaign with dynamic shares
	camp = campaigntypes.NewCampaign(
		1,
		sample.AlphaString(5),
		coords[1],
		sample.Coins(),
		true,
	)
	camp.MainnetInitialized = true
	idDynamicShares := ck.AppendCampaign(ctx, camp)
	acc, id, found = simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, true, false)
	require.True(t, found)
	require.Contains(t, accs, acc)
	camp, found = ck.GetCampaign(ctx, id)
	require.True(t, found)
	require.EqualValues(t, idDynamicShares, id)
	require.EqualValues(t, id, camp.Id)
	require.True(t, camp.DynamicShares)

	// Find a campaign with no mainnet initialized
	camp = campaigntypes.NewCampaign(
		2,
		sample.AlphaString(5),
		coords[2],
		sample.Coins(),
		false,
	)
	idNoMainnet := ck.AppendCampaign(ctx, camp)
	acc, id, found = simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, false, true)
	require.True(t, found)
	require.Contains(t, accs, acc)
	_, found = ck.GetCampaign(ctx, id)
	require.True(t, found)
	require.EqualValues(t, idNoMainnet, id)
	require.EqualValues(t, id, camp.Id)
	require.False(t, camp.MainnetInitialized)

	// Find a campaign with no mainnet initialized and with dynamic shares
	camp = campaigntypes.NewCampaign(
		3,
		sample.AlphaString(5),
		coords[3],
		sample.Coins(),
		true,
	)
	idNoMainnetDynamicShares := ck.AppendCampaign(ctx, camp)
	acc, id, found = simcampaign.GetCoordSimAccountWithCampaignID(r, ctx, pk, *ck, accs, true, true)
	require.True(t, found)
	require.Contains(t, accs, acc)
	_, found = ck.GetCampaign(ctx, id)
	require.True(t, found)
	require.EqualValues(t, idNoMainnetDynamicShares, id)
	require.EqualValues(t, id, camp.Id)
	require.False(t, camp.MainnetInitialized)
	require.True(t, camp.DynamicShares)
}
