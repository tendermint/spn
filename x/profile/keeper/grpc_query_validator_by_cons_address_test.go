package keeper

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/x/profile/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestValidatorByConsAddressQuerySingle(t *testing.T) {
	keeper, ctx := setupKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidatorByConsAddress(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorByConsAddressRequest
		response *types.QueryGetValidatorByConsAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetValidatorByConsAddressRequest{
				ConsAddress: msgs[0].ConsAddress,
			},
			response: &types.QueryGetValidatorByConsAddressResponse{ValidatorByConsAddress: &msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetValidatorByConsAddressRequest{
				ConsAddress: msgs[1].ConsAddress,
			},
			response: &types.QueryGetValidatorByConsAddressResponse{ValidatorByConsAddress: &msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetValidatorByConsAddressRequest{
				ConsAddress: strconv.Itoa(100000),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ValidatorByConsAddress(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.Equal(t, tc.response, response)
			}
		})
	}
}
