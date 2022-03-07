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
	valAddr := msg.ValidatorAddress
	opAddr := msg.OperatorAddress

	validator := types.Validator{
		Address:           valAddr,
		OperatorAddresses: []string{opAddr},
		Description:       types.ValidatorDescription{},
	}

	// remove the operator address from previous address
	if previousValByOpAddr, found := k.GetValidatorByOperatorAddress(ctx, opAddr); found {
		lastValidator, found := k.GetValidator(ctx, previousValByOpAddr.ValidatorAddress)
		if !found {
			return nil, spnerrors.Criticalf(
				"validator should exist for operator address %s",
				previousValByOpAddr.ValidatorAddress,
			)
		}
		lastValidator = lastValidator.RemoveValidatorOperatorAddress(opAddr)
		k.SetValidator(ctx, lastValidator)
	}

	// get the current validator to eventually overwrite description and remove existing operator address
	if validatorStore, found := k.GetValidator(ctx, valAddr); found {
		validator.Description = validatorStore.Description
		validator = validatorStore.AddValidatorOperatorAddress(opAddr)
	}

	// store validator information
	k.SetValidator(ctx, validator)
	k.SetValidatorByOperatorAddress(ctx, types.ValidatorByOperatorAddress{
		OperatorAddress:  opAddr,
		ValidatorAddress: valAddr,
	})

	return &types.MsgAddValidatorOperatorAddressResponse{}, nil
}
