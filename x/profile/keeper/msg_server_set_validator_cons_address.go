package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
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
	// consensusNonce, found := k.GetConsensusKeyNonce(ctx, msg.ConsAddress)
	// if found {
	// 	currentNonce = consensusNonce.Nonce
	// }

	if !checkValidatorSignature(msg.Signature, msg.ConsAddress, currentNonce) {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrapf(types.ErrInvalidValidatorSignature, "consensus address: %s / signature: %s", msg.ConsAddress, msg.Signature)
	}

	// k.SetConsensusKeyNonce(ctx, types.ConsensusKeyNonce{
	//	ConsensusAddress: consensusNonce.ConsensusAddress,
	//	Nonce: currentNonce+1},
	// })

	// get the current validator to eventually overwrite description
	validator, found := k.GetValidator(ctx, validatorByConsAddr.ValidatorAddress)
	if !found {
		return &types.MsgSetValidatorConsAddressResponse{},
			spnerrors.Criticalf("a validator consensus address %s is associated to a non-existent validator address %s",
				validatorByConsAddr.ConsensusAddress,
				validatorByConsAddr.ValidatorAddress,
			)
	}

	//validator := load(ValidatorsByAddress, msg.Address)
	//
	//newValidator := newValidator()
	//newValidator.Address = msg.Address
	//newValidator.ConsAddress = msg.ConsAddress
	//if validator != nil
	//	newValidator.Description = validator.Description
	//
	//// store validator information
	//store(ValidatorsByAddress, msg.Address, validator)
	//store(ValidatorsByConsKey, msg.ConsAddress, msg.Address)

	return &types.MsgSetValidatorConsAddressResponse{}, nil
}
