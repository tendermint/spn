package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := sample.LaunchParams(r)
	tk.LaunchKeeper.SetParams(ctx, params)

	t.Run("should allow querying params", func(t *testing.T) {
		response, err := tk.LaunchKeeper.Params(wctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
	})

	t.Run("should prevent querying params with invalid request", func(t *testing.T) {
		_, err := tk.LaunchKeeper.Params(wctx, nil)
		require.Error(t, err)
	})
}
