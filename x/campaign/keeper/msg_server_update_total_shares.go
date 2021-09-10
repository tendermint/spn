package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateTotalShares(goCtx context.Context, msg *types.MsgUpdateTotalShares) (*types.MsgUpdateTotalSharesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%v", msg.CampaignID)
	}

	if !campaign.DynamicShares {
		return nil, sdkerrors.Wrap(types.ErrNoDynamicShares, "campaign doesn't has dynamic shares option set")
	}

	// Get the coordinator ID
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}
	if campaign.CoordinatorID != coordinatorID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the campaign is %v",
			campaign.CoordinatorID,
		))
	}

	if campaign.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%v", msg.CampaignID)
	}

	// Check the new total shares are not overflown by the currently allocated shares
	if types.IsTotalSharesReached(campaign.AllocatedShares, msg.TotalShares) {
		return nil, sdkerrors.Wrap(types.ErrInvalidShares, "more allocated shares than total shares")
	}
	campaign.TotalShares = msg.TotalShares
	k.SetCampaign(ctx, campaign)

	return &types.MsgUpdateTotalSharesResponse{}, nil
}
