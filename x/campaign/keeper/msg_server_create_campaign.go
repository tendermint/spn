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

	// Validate provided totalSupply
	totalSupplyRange := k.TotalSupplyRange(ctx)
	if err := types.ValidateTotalSupply(msg.TotalSupply, totalSupplyRange); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	// Append the new campaign
	campaign := types.NewCampaign(0, msg.CampaignName, coordID, msg.TotalSupply, false)
	campaignID := k.AppendCampaign(ctx, campaign)

	// Initialize the list of campaign chains
	k.SetCampaignChains(ctx, types.CampaignChains{
		CampaignID: campaignID,
		Chains:     []uint64{},
	})

	return &types.MsgCreateCampaignResponse{CampaignID: campaignID}, nil
}
