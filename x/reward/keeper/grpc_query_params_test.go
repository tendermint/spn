package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk := testkeeper.NewTestKeepers(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	tk.RewardKeeper.SetParams(ctx, params)

	response, err := tk.RewardKeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
