package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

func TestParamsQuery(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)

	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	tk.ClaimKeeper.SetParams(ctx, params)

	response, err := tk.ClaimKeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
