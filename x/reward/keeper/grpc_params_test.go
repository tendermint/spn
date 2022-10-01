package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()

	t.Run("should allow query for params", func(t *testing.T) {
		tk.RewardKeeper.SetParams(ctx, params)
		response, err := tk.RewardKeeper.Params(wctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, &types.QueryParamsResponse{Params: params}, response)

		_, err = tk.RewardKeeper.Params(wctx, nil)
		require.Error(t, err)
	})
}
