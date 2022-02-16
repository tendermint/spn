package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateTotalSupply(goCtx context.Context, msg *types.MsgUpdateTotalSupply) (*types.MsgUpdateTotalSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if campaign.CoordinatorID != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the campaign is %d",
			campaign.CoordinatorID,
		))
	}

	if campaign.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%d", msg.CampaignID)
	}

	// Validate provided totalSupply
	totalSupplyRange := k.TotalSupplyRange(ctx)
	if err := types.ValidateTotalSupply(msg.TotalSupplyUpdate, totalSupplyRange); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	campaign.TotalSupply = types.UpdateTotalSupply(campaign.TotalSupply, msg.TotalSupplyUpdate)
	k.SetCampaign(ctx, campaign)

	return &types.MsgUpdateTotalSupplyResponse{}, nil
}
