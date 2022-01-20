package keeper_test

import (
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgSetValidatorConsAddress(t *testing.T) {
	var (
		validatorAddr           = sample.Address()
		consensusAddr           = sample.Address()
		consensusAddrWithoutAcc = sample.Address()
		ctx, k, srv             = setupMsgServer(t)
		wCtx                    = sdk.WrapSDKContext(ctx)
	)
	k.SetValidator(ctx, types.Validator{
		Address:     validatorAddr,
		Description: types.ValidatorDescription{},
	})
	k.SetValidatorByConsAddress(ctx, types.ValidatorByConsAddress{
		ValidatorAddress: validatorAddr,
		ConsensusAddress: consensusAddr,
	})
	// TODO add account to Starport

	k.SetValidatorByConsAddress(ctx, types.ValidatorByConsAddress{
		ValidatorAddress: sample.Address(),
		ConsensusAddress: consensusAddrWithoutAcc,
	})
	signature := ""
	tests := []struct {
		name string
		msg  *types.MsgSetValidatorConsAddress
		err  error
	}{
		{
			name: "valid message",
			msg: &types.MsgSetValidatorConsAddress{
				Creator:     validatorAddr,
				Address:     consensusAddr,
				ConsAddress: consensusAddr,
				Signature:   signature,
			},
		},
		{
			name: "consensus address not found",
			msg: &types.MsgSetValidatorConsAddress{
				Creator:     validatorAddr,
				Address:     consensusAddr,
				ConsAddress: sample.Address(),
				Signature:   signature,
			},
			err: types.ErrValidatorConsAddressNotFound,
		},
		{
			name: "invalid consensus address",
			msg: &types.MsgSetValidatorConsAddress{
				Creator:     validatorAddr,
				Address:     consensusAddr,
				ConsAddress: "invalid_address",
				Signature:   signature,
			},
			err: spnerrors.ErrCritical,
		},
		{
			name: "consensus account not found",
			msg: &types.MsgSetValidatorConsAddress{
				Creator:     validatorAddr,
				Address:     consensusAddr,
				ConsAddress: sample.Address(),
				Signature:   signature,
			},
			err: types.ErrConsdAccNotFound,
		},
		{
			name: "invalid signature",
			msg: &types.MsgSetValidatorConsAddress{
				Creator:     validatorAddr,
				Address:     consensusAddr,
				ConsAddress: consensusAddr,
				Signature:   "invalid_signature",
			},
			err: types.ErrInvalidValidatorSignature,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.SetValidatorConsAddress(wCtx, tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			_, found := k.GetValidator(ctx, tt.msg.Address)
			require.False(t, found, "validator was not removed")
		})
	}
}
