package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) CreateCampaign(goCtx context.Context, msg *types.MsgCreateCampaign) (*types.MsgCreateCampaignResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	// Append the new campaign
	campaign := types.NewCampaign(0, msg.CampaignName, coordID, msg.TotalSupply, false, msg.Metadata)
	campaignID := k.AppendCampaign(ctx, campaign)

	// Initialize the list of campaign chains
	k.SetCampaignChains(ctx, types.CampaignChains{
		CampaignID: campaignID,
		Chains:     []uint64{},
	})

	return &types.MsgCreateCampaignResponse{CampaignID: campaignID}, nil
}
