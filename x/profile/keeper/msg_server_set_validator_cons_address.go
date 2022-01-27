package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	valtypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) SetValidatorConsAddress(
	goCtx context.Context,
	msg *types.MsgSetValidatorConsAddress,
) (*types.MsgSetValidatorConsAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valKey, err := valtypes.LoadValidatorKey(msg.ValidatorKey)
	if err != nil {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrap(types.ErrInvalidValidatorKey, string(msg.ValidatorKey))
	}
	consAddress := valKey.GetConsAddress().String()

	// cannot set the consensus key if key is used for another validator
	validatorByConsAddr, found := k.GetValidatorByConsAddress(ctx, consAddress)
	if !found {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrap(types.ErrValidatorConsAddressNotFound, consAddress)
	}

	// check signature
	currentNonce := uint64(0)
	consensusNonce, found := k.GetConsensusKeyNonce(ctx, consAddress)
	if found {
		currentNonce = consensusNonce.Nonce
	}

	if !valKey.VerifySignature(currentNonce, []byte(msg.Signature)) {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrapf(types.ErrInvalidValidatorSignature,
				"invalid signature for consensus address: %s - %s",
				consAddress,
				msg.Signature,
			)
	}

	validator := types.Validator{
		Address:          valKey.Address.String(),
		ConsensusAddress: consAddress,
		Description:      types.ValidatorDescription{},
	}

	// get the current validator to eventually overwrite description
	validatorStore, found := k.GetValidator(ctx, validatorByConsAddr.ValidatorAddress)
	if found {
		validator.Description = validatorStore.Description
	}

	// store validator information
	k.SetValidator(ctx, validator)
	k.SetValidatorByConsAddress(ctx, types.ValidatorByConsAddress{
		ConsensusAddress: consAddress,
		ValidatorAddress: valKey.Address.String(),
	})
	k.SetConsensusKeyNonce(ctx, types.ConsensusKeyNonce{
		ConsensusAddress: consensusNonce.ConsensusAddress,
		Nonce:            currentNonce + 1,
	})

	return &types.MsgSetValidatorConsAddressResponse{}, nil
}
