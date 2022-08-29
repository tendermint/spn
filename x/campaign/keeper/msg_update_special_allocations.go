package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) UpdateSpecialAllocations(goCtx context.Context, msg *types.MsgUpdateSpecialAllocations) (*types.MsgUpdateSpecialAllocationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	// get the coordinator ID associated to the sender address
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

	// verify mainnet launch is not triggered
	mainnetLaunched, err := k.IsCampaignMainnetLaunchTriggered(ctx, campaign.CampaignID)
	if err != nil {
		return nil, spnerrors.Critical(err.Error())
	}
	if mainnetLaunched {
		return nil, sdkerrors.Wrap(types.ErrMainnetLaunchTriggered, fmt.Sprintf(
			"mainnet %d launch is already triggered",
			campaign.MainnetID,
		))
	}

	// decrease allocated shares from current special allocations
	campaign.AllocatedShares, err = types.DecreaseShares(campaign.AllocatedShares, campaign.SpecialAllocations.TotalShares())
	if err != nil {
		return nil, spnerrors.Critical("campaign allocated shares should be bigger than current special allocations" + err.Error())
	}

	// increase with new special allocations
	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, msg.SpecialAllocations.TotalShares())

	// increase the campaign shares
	reached, err := types.IsTotalSharesReached(campaign.AllocatedShares, k.GetTotalShares(ctx))
	if err != nil {
		return nil, spnerrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if reached {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.CampaignID)
	}

	campaign.SpecialAllocations = msg.SpecialAllocations
	k.SetCampaign(ctx, campaign)
	err = ctx.EventManager().EmitTypedEvents(
		&types.EventCampaignSharesUpdated{
			CampaignID:         campaign.CampaignID,
			CoordinatorAddress: msg.Coordinator,
			AllocatedShares:    campaign.AllocatedShares,
		})
	return &types.MsgUpdateSpecialAllocationsResponse{}, err
}
