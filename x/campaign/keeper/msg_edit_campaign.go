package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) EditCampaign(goCtx context.Context, msg *types.MsgEditCampaign) (*types.MsgEditCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// check if the metadata length is valid
	maxMetadataLength := k.MaxMetadataLength(ctx)
	if uint64(len(msg.Metadata)) > maxMetadataLength {
		return nil, sdkerrors.Wrapf(types.ErrInvalidMetadataLength,
			"metadata length %d is greater than maximum %d",
			len(msg.Metadata),
			maxMetadataLength,
		)
	}

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

	err = ctx.EventManager().EmitTypedEvent(&types.EventCampaignInfoUpdated{
		CampaignID:         campaign.CampaignID,
		CoordinatorAddress: msg.Coordinator,
		CampaignName:       campaign.CampaignName,
		Metadata:           campaign.Metadata,
	})

	return &types.MsgEditCampaignResponse{}, err
}
