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
	coordByAddress, found := k.profileKeeper.GetCoordinatorByAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	coord, found := k.profileKeeper.GetCoordinator(ctx, coordByAddress.CoordinatorID)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	if !coord.Active {
		return nil, sdkerrors.Wrapf(profiletypes.ErrCoordInactive,
			"coordinator %d inactive", coord.CoordinatorID)
	}

	if campaign.CoordinatorID != coord.CoordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the campaign is %d",
			campaign.CoordinatorID,
		))
	}

	if campaign.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%d", msg.CampaignID)
	}

	campaign.TotalSupply = types.UpdateTotalSupply(campaign.TotalSupply, msg.TotalSupplyUpdate)
	k.SetCampaign(ctx, campaign)

	return &types.MsgUpdateTotalSupplyResponse{}, nil
}
