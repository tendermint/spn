package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateValidatorDescription(
	goCtx context.Context,
	msg *types.MsgUpdateValidatorDescription,
) (*types.MsgUpdateValidatorDescriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the validator address is already in the store
	validator, valfound := k.GetValidator(ctx, msg.Address)
	if !valfound {
		validator = types.Validator{
			Address:     msg.Address,
			Description: types.ValidatorDescription{},
		}
	}

	if len(msg.Description.Identity) > 0 {
		validator.Description.Identity = msg.Description.Identity
	}
	if len(msg.Description.Website) > 0 {
		validator.Description.Website = msg.Description.Website
	}
	if len(msg.Description.Details) > 0 {
		validator.Description.Details = msg.Description.Details
	}
	if len(msg.Description.Moniker) > 0 {
		validator.Description.Moniker = msg.Description.Moniker
	}
	if len(msg.Description.SecurityContact) > 0 {
		validator.Description.SecurityContact = msg.Description.SecurityContact
	}

	k.SetValidator(ctx, validator)
	var err error
	if !valfound {
		err = ctx.EventManager().EmitTypedEvent(
			&types.EventValidatorCreated{
				Address:           validator.Address,
				OperatorAddresses: validator.OperatorAddresses,
			})
	}

	return &types.MsgUpdateValidatorDescriptionResponse{}, err
}
