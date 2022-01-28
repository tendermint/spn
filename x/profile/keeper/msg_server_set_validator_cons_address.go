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

	valPubKey, err := valtypes.NewValidatorConsPubKey(msg.ValidatorConsPubKey, msg.ValidatorKeyType)
	if err != nil {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrap(types.ErrInvalidValidatorKey, string(msg.ValidatorConsPubKey))
	}
	consAddress := valPubKey.GetConsAddress().String()

	// check signature
	currentNonce := uint64(0)
	consensusNonce, found := k.GetConsensusKeyNonce(ctx, consAddress)
	if found {
		currentNonce = consensusNonce.Nonce
	}
	if currentNonce != msg.Nonce {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrapf(types.ErrInvalidValidatorNonce, "%d", msg.Nonce)
	}
	if ctx.ChainID() != msg.ChainID {
		return &types.MsgSetValidatorConsAddressResponse{},
			sdkerrors.Wrap(types.ErrInvalidValidatorChainID, msg.ChainID)
	}

	validator := types.Validator{
		Address:          valPubKey.Address().String(),
		ConsensusAddress: consAddress,
		Description:      types.ValidatorDescription{},
	}

	// get the current validator to eventually overwrite description
	validatorStore, found := k.GetValidator(ctx, msg.ValidatorAddress)
	if found {
		validator.Description = validatorStore.Description
		k.RemoveValidatorByConsAddress(ctx, validatorStore.ConsensusAddress)
	}

	// store validator information
	k.SetValidator(ctx, validator)
	k.SetValidatorByConsAddress(ctx, types.ValidatorByConsAddress{
		ConsensusAddress: consAddress,
		ValidatorAddress: valPubKey.Address().String(),
	})
	k.SetConsensusKeyNonce(ctx, types.ConsensusKeyNonce{
		ConsensusAddress: consAddress,
		Nonce:            currentNonce + 1,
	})

	return &types.MsgSetValidatorConsAddressResponse{}, nil
}
