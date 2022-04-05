package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
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
		if errors.Is(err, types.ErrInvalidSupplyRange) {
			return nil, spnerrors.Critical(err.Error())
		}

		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	// Deduct campaign creation fee if set
	creationFee := k.CampaignCreationFee(ctx)
	if !creationFee.Empty() {
		coordAddr, err := sdk.AccAddressFromBech32(msg.Coordinator)
		if err != nil {
			return nil, spnerrors.Criticalf("invalid coordinator bech32 address %s", err.Error())
		}
		if err = k.distrKeeper.FundCommunityPool(ctx, creationFee, coordAddr); err != nil {
			return nil, err
		}
	}

	// Append the new campaign
	campaign := types.NewCampaign(0, msg.CampaignName, coordID, msg.TotalSupply, msg.Metadata)
	campaignID := k.AppendCampaign(ctx, campaign)

	// Initialize the list of campaign chains
	k.SetCampaignChains(ctx, types.CampaignChains{
		CampaignID: campaignID,
		Chains:     []uint64{},
	})

	return &types.MsgCreateCampaignResponse{CampaignID: campaignID}, nil
}
