package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	profiletypes "github.com/tendermint/spn/x/profile/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) CreateCampaign(goCtx context.Context, msg *types.MsgCreateCampaign) (*types.MsgCreateCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	// Append the new campaign
	campaign := types.NewCampaign(0, msg.CampaignName, coordinatorID, msg.TotalSupply, msg.DynamicShares)
	campaignID := k.AppendCampaign(ctx, campaign)

	return &types.MsgCreateCampaignResponse{CampaignID: campaignID}, nil
}
