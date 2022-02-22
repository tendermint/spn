package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func Test_GetParams(t *testing.T) {
	k, ctx := testkeeper.Campaign(t)
	params := sample.CampaignParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.TotalSupplyRange.MinTotalSupply, k.TotalSupplyRange(ctx).MinTotalSupply)
	require.EqualValues(t, params.TotalSupplyRange.MaxTotalSupply, k.TotalSupplyRange(ctx).MaxTotalSupply)
	require.EqualValues(t, params.CampaignCreationFee, k.CampaignCreationFee(ctx))
}
