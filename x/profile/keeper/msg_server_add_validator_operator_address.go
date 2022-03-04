package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) AddValidatorOperatorAddress(
	goCtx context.Context,
	msg *types.MsgAddValidatorOperatorAddress,
) (*types.MsgAddValidatorOperatorAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	opAddr := msg.OperatorAddress

	validator := types.Validator{
		Address:           msg.ValidatorAddress,
		OperatorAddresses: []string{opAddr},
		Description:       types.ValidatorDescription{},
	}

	// remove the operator address from previous address
	if valByOpAddr, found := k.GetValidatorByOperatorAddress(ctx, opAddr); found {
		lastValidator, found := k.GetValidator(ctx, valByOpAddr.ValidatorAddress)
		if !found {
			return nil, spnerrors.Criticalf(
				"validator should exist for operator address %s",
				valByOpAddr.ValidatorAddress,
			)
		}
		lastValidator = lastValidator.RemoveValidatorOperatorAddress(opAddr)
		k.SetValidator(ctx, lastValidator)
	}

	// get the current validator to eventually overwrite description and remove existing operator address
	if validatorStore, found := k.GetValidator(ctx, msg.ValidatorAddress); found {
		validator.Description = validatorStore.Description
		validator = validatorStore.AddValidatorOperatorAddress(opAddr)
	}

	// store validator information
	k.SetValidator(ctx, validator)
	k.SetValidatorByOperatorAddress(ctx, types.ValidatorByOperatorAddress{
		OperatorAddress:  opAddr,
		ValidatorAddress: msg.ValidatorAddress,
	})

	return &types.MsgAddValidatorOperatorAddressResponse{}, nil
}
