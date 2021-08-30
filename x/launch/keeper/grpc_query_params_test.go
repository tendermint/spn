package keeper_test

import (
	"github.com/tendermint/spn/testutil/sample"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/launch/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, _, ctx, _ := testkeeper.Launch(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := sample.LaunchParams()
	keeper.SetParams(ctx, params)

	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
