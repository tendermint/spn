package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/claim/types"
)

func TestAirdropSupplyQuery(t *testing.T) {
	var (
		ctx, tk, _   = testkeeper.NewTestSetup(t)
		wctx         = sdk.WrapSDKContext(ctx)
		sampleSupply = sample.Coin(r)
	)
	tk.ClaimKeeper.SetAirdropSupply(ctx, sampleSupply)

	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetAirdropSupplyRequest
		response *types.QueryGetAirdropSupplyResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetAirdropSupplyRequest{},
			response: &types.QueryGetAirdropSupplyResponse{AirdropSupply: sampleSupply},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ClaimKeeper.AirdropSupply(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
