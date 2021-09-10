package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	profiletypes "github.com/tendermint/spn/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) UpdateTotalSupply(goCtx context.Context, msg *types.MsgUpdateTotalSupply) (*types.MsgUpdateTotalSupplyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%v", msg.CampaignID)
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

	campaign.TotalSupply = msg.TotalSupply
	k.SetCampaign(ctx, campaign)

	return &types.MsgUpdateTotalSupplyResponse{}, nil
}
