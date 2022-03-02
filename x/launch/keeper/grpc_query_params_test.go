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
	ctx, tk := testkeeper.NewTestKeepers(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := sample.LaunchParams()
	tk.LaunchKeeper.SetParams(ctx, params)

	response, err := tk.LaunchKeeper.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
