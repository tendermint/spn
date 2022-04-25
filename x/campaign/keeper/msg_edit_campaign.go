package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) EditCampaign(goCtx context.Context, msg *types.MsgEditCampaign) (*types.MsgEditCampaignResponse, error) {
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

	if len(msg.Name) > 0 {
		campaign.CampaignName = msg.Name
	}

	if len(msg.Metadata) > 0 {
		campaign.Metadata = msg.Metadata
	}

	k.SetCampaign(ctx, campaign)

	err = ctx.EventManager().EmitTypedEvent(&types.EventCampaignUpdated{Campaign: campaign})

	return &types.MsgEditCampaignResponse{}, err
}
