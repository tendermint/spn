package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) AddMainnetVestingAccount(goCtx context.Context, msg *types.MsgAddMainnetVestingAccount) (*types.MsgAddMainnetVestingAccountResponse, error) {
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

	if campaign.MainnetInitialized {
		return nil, sdkerrors.Wrapf(types.ErrMainnetInitialized, "%v", msg.CampaignID)
	}

	account := types.MainnetVestingAccount{
		CampaignID:     campaign.Id,
		Address:        msg.Address,
		Shares:         msg.Shares,
		VestingOptions: msg.VestingOptions,
	}

	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, msg.Shares)
	if types.IsTotalSharesReached(campaign.AllocatedShares, campaign.TotalShares) {
		return nil, sdkerrors.Wrapf(types.ErrTotalShareLimit, "%v", msg.CampaignID)
	}

	k.SetMainnetVestingAccount(ctx, account)
	k.SetCampaign(ctx, campaign)

	return &types.MsgAddMainnetVestingAccountResponse{}, nil
}
