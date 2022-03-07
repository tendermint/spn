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

func TestValidatorByConsAddressQuerySingle(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetup(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNValidatorByConsAddress(tk.ProfileKeeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetValidatorByConsAddressRequest
		response *types.QueryGetValidatorByConsAddressResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetValidatorByConsAddressRequest{
				ConsensusAddress: msgs[0].ConsensusAddress,
			},
			response: &types.QueryGetValidatorByConsAddressResponse{ValidatorByConsAddress: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetValidatorByConsAddressRequest{
				ConsensusAddress: msgs[1].ConsensusAddress,
			},
			response: &types.QueryGetValidatorByConsAddressResponse{ValidatorByConsAddress: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetValidatorByConsAddressRequest{
				ConsensusAddress: sample.ConsAddress().Bytes(),
			},
			err: status.Error(codes.InvalidArgument, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := tk.ProfileKeeper.ValidatorByConsAddress(wctx, tc.request)
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