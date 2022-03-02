package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestValidatorByConsAddressQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.Profile(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidatorByOperatorAddress(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorByOperatorAddressRequest
		response *types.QueryGetValidatorByOperatorAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				OperatorAddress: msgs[0].OperatorAddress,
			},
			response: &types.QueryGetValidatorByOperatorAddressResponse{ValidatorByOperatorAddress: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				OperatorAddress: msgs[1].OperatorAddress,
			},
			response: &types.QueryGetValidatorByOperatorAddressResponse{ValidatorByOperatorAddress: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				OperatorAddress: sample.Address(),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ValidatorByConsAddress(wctx, tc.request)
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
