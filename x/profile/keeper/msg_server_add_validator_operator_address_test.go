package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgAddValidatorOperatorAddress(t *testing.T) {
	var (
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
		valAddr     = sample.Address()
	)

	k.SetValidator(ctx, types.Validator{
		Address:     valAddr,
		Description: types.ValidatorDescription{},
	})

	tests := []struct {
		name string
		msg  *types.MsgAddValidatorOperatorAddress
	}{
		{
			name: "should allow associating a new operator address to a validator",
			msg: &types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: valAddr,
				OperatorAddress:  sample.Address(),
			},
		},
		{
			name: "should allow to associate the same address",
			msg: &types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: valAddr,
				OperatorAddress:  valAddr,
			},
		},
		{
			name: "should create a validator is it doesn't exist",
			msg: &types.MsgAddValidatorOperatorAddress{
				ValidatorAddress: sample.Address(),
				OperatorAddress:  sample.Address(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.AddValidatorOperatorAddress(wCtx, tt.msg)
			require.NoError(t, err)

			validator, found := k.GetValidator(ctx, tt.msg.ValidatorAddress)
			require.True(t, found, "validator was not saved")
			require.Equal(t, tt.msg.ValidatorAddress, validator.Address)
			require.True(t, validator.HasOperatorAddress(tt.msg.OperatorAddress))

			valByOpAddr, found := k.GetValidatorByOperatorAddress(ctx, tt.msg.OperatorAddress)
			require.True(t, found, "validator by operator address was not saved")
			require.Equal(t, tt.msg.ValidatorAddress, valByOpAddr.ValidatorAddress)
			require.Equal(t, tt.msg.OperatorAddress, valByOpAddr.OperatorAddress)
		})
	}
}
