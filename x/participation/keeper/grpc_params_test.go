package keeper_test

import (
	"testing"

	"github.com/tendermint/spn/testutil/sample"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)

	t.Run("should allow params get", func(t *testing.T) {
		params := sample.ParticipationParams(r)
		tk.ParticipationKeeper.SetParams(ctx, params)

		response, err := tk.ParticipationKeeper.Params(wctx, &types.QueryParamsRequest{})
		require.NoError(t, err)
		require.Equal(t, &types.QueryParamsResponse{Params: params}, response)

		_, err = tk.ParticipationKeeper.Params(wctx, nil)
		require.Error(t, err)
	})
}
