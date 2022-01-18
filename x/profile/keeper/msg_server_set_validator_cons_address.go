package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) SetValidatorConsAddress(goCtx context.Context, msg *types.MsgSetValidatorConsAddress) (*types.MsgSetValidatorConsAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// cannot set the consensus key if it's used for another validator
	validatorByConsAddr, found := k.GetValidatorByConsAddress(ctx, msg.ConsAddress)
	if !found {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrap(types.ErrValidatorConsAddressNotFound, msg.ConsAddress)
	}

	// check signature
	currentNonce := uint64(0)
	consensusNonce, found := k.GetConsensusKeyNonce(ctx, msg.ConsAddress)
	if found {
		currentNonce = consensusNonce.Nonce
	}

	if !checkValidatorSignature(msg.Signature, msg.ConsAddress, currentNonce) {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrapf(types.ErrInvalidValidatorSignature, "consensus address: %s / signature: %s", msg.ConsAddress, msg.Signature)
	}

	validator := types.Validator{
		Address:          msg.Address,
		ConsensusAddress: msg.ConsAddress,
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
		ConsensusAddress: msg.ConsAddress,
		ValidatorAddress: msg.Address,
	})
	k.SetConsensusKeyNonce(ctx, types.ConsensusKeyNonce{
		ConsensusAddress: consensusNonce.ConsensusAddress,
		Nonce:            currentNonce + 1,
	})

	return &types.MsgSetValidatorConsAddressResponse{}, nil
}
