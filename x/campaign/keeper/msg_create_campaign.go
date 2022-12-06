package keeper

import (
	"context"
	"errors"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) CreateCampaign(goCtx context.Context, msg *types.MsgCreateCampaign) (*types.MsgCreateCampaignResponse, error) {
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

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	// Validate provided totalSupply
	totalSupplyRange := k.TotalSupplyRange(ctx)
	if err := types.ValidateTotalSupply(msg.TotalSupply, totalSupplyRange); err != nil {
		if errors.Is(err, types.ErrInvalidSupplyRange) {
			return nil, ignterrors.Critical(err.Error())
		}

		return nil, sdkerrors.Wrap(types.ErrInvalidTotalSupply, err.Error())
	}

	// Deduct campaign creation fee if set
	creationFee := k.CampaignCreationFee(ctx)
	if !creationFee.Empty() {
		coordAddr, err := sdk.AccAddressFromBech32(msg.Coordinator)
		if err != nil {
			return nil, ignterrors.Criticalf("invalid coordinator bech32 address %s", err.Error())
		}
		if err = k.distrKeeper.FundCommunityPool(ctx, creationFee, coordAddr); err != nil {
			return nil, sdkerrors.Wrap(types.ErrFundCommunityPool, err.Error())
		}
	}

	// Append the new campaign
	campaign := types.NewCampaign(
		0,
		msg.CampaignName,
		coordID,
		msg.TotalSupply,
		msg.Metadata,
		ctx.BlockTime().Unix(),
	)
	campaignID := k.AppendCampaign(ctx, campaign)

	// Initialize the list of campaign chains
	k.SetCampaignChains(ctx, types.CampaignChains{
		CampaignID: campaignID,
		Chains:     []uint64{},
	})

	err = ctx.EventManager().EmitTypedEvent(&types.EventCampaignCreated{
		CampaignID:         campaignID,
		CoordinatorAddress: msg.Coordinator,
		CoordinatorID:      coordID,
	})

	return &types.MsgCreateCampaignResponse{CampaignID: campaignID}, err
}
