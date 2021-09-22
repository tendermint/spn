package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) AddMainnetAccount(goCtx context.Context, msg *types.MsgAddMainnetAccount) (*types.MsgAddMainnetAccountResponse, error) {
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

	// check if the account already exists
	account, found := k.GetMainnetAccount(ctx, campaign.Id, msg.Address)
	if !found {
		// if not, create the account
		account = types.MainnetAccount{
			CampaignID: campaign.Id,
			Address:    msg.Address,
		}
	}
	// increase the account shares
	account.Shares = types.IncreaseShares(account.Shares, msg.Shares)

	// increase the campaign shares
	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, msg.Shares)
	if types.IsTotalSharesReached(campaign.AllocatedShares, campaign.TotalShares) {
		return nil, sdkerrors.Wrapf(types.ErrTotalShareLimit, "%v", msg.CampaignID)
	}

	k.SetCampaign(ctx, campaign)
	k.SetMainnetAccount(ctx, account)

	return &types.MsgAddMainnetAccountResponse{}, nil
}
