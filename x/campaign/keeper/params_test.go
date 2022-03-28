package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func Test_GetParams(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	params := sample.CampaignParams(r)

	tk.CampaignKeeper.SetParams(ctx, params)

	require.EqualValues(t, params, tk.CampaignKeeper.GetParams(ctx))
	require.EqualValues(t, params.TotalSupplyRange.MinTotalSupply, tk.CampaignKeeper.TotalSupplyRange(ctx).MinTotalSupply)
	require.EqualValues(t, params.TotalSupplyRange.MaxTotalSupply, tk.CampaignKeeper.TotalSupplyRange(ctx).MaxTotalSupply)
	require.EqualValues(t, params.CampaignCreationFee, tk.CampaignKeeper.CampaignCreationFee(ctx))
}
