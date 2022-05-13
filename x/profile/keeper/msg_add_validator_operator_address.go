package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

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

	// get the current validator to eventually overwrite description and add opAddr
	validatorStore, valFound := k.GetValidator(ctx, valAddr)
	if valFound {
		validator.Description = validatorStore.Description
		validator = validatorStore.AddValidatorOperatorAddress(opAddr)
	}

	// store validator information
	k.SetValidator(ctx, validator)
	k.SetValidatorByOperatorAddress(ctx, types.ValidatorByOperatorAddress{
		OperatorAddress:  opAddr,
		ValidatorAddress: valAddr,
	})

	var err error
	if valFound {
		err = ctx.EventManager().EmitTypedEvent(
			&types.EventValidatorOperatorAddressesUpdated{
				Address:           validator.Address,
				OperatorAddresses: validator.OperatorAddresses})
	} else {
		err = ctx.EventManager().EmitTypedEvent(
			&types.EventValidatorCreated{
				Address:           validator.Address,
				OperatorAddresses: validator.OperatorAddresses,
			})
	}

	return &types.MsgAddValidatorOperatorAddressResponse{}, err
}
