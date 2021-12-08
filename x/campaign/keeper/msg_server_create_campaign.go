package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) CreateCampaign(goCtx context.Context, msg *types.MsgCreateCampaign) (*types.MsgCreateCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID
	coordinatorID, found := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if !found {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordAddressNotFound, msg.Coordinator)
	}

	// Append the new campaign
	campaign := types.NewCampaign(0, msg.CampaignName, coordinatorID, msg.TotalSupply, false)
	campaignID := k.AppendCampaign(ctx, campaign)

	// Initialize the list of campaign chains
	k.SetCampaignChains(ctx, types.CampaignChains{
		CampaignID: campaignID,
		Chains:     []uint64{},
	})

	return &types.MsgCreateCampaignResponse{CampaignID: campaignID}, nil
}
