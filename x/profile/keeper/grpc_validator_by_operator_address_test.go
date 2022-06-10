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
	"github.com/tendermint/spn/x/profile/types"
)

func TestValidatorByOperatorAddressQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidatorByOperatorAddress(tk.ProfileKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorByOperatorAddressRequest
		response *types.QueryGetValidatorByOperatorAddressResponse
		err      error
	}{
		{
			desc: "should allow querying first validator by operator address",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				OperatorAddress: msgs[0].OperatorAddress,
			},
			response: &types.QueryGetValidatorByOperatorAddressResponse{ValidatorByOperatorAddress: msgs[0]},
		},
		{
			desc: "should allow querying second validator by operator address",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				OperatorAddress: msgs[1].OperatorAddress,
			},
			response: &types.QueryGetValidatorByOperatorAddressResponse{ValidatorByOperatorAddress: msgs[1]},
		},
		{
			desc: "should prevent querying non existing validator by operator address",
			request: &types.QueryGetValidatorByOperatorAddressRequest{
				OperatorAddress: sample.Address(r),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "should prevent querying with invalid request",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProfileKeeper.ValidatorByOperatorAddress(wctx, tc.request)
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
