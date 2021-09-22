package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
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

	// check if the account already exists
	oldAccount, foundAcc := k.GetMainnetVestingAccount(ctx, campaign.Id, msg.Address)
	if foundAcc {
		// if yes, remove the allocated share
		totalShares, err := oldAccount.GetTotalShares()
		if err != nil {
			return nil, spnerrors.Critical(err.Error())
		}
		campaign.AllocatedShares, err = types.DecreaseShares(campaign.AllocatedShares, totalShares)
		if err != nil {
			return nil, sdkerrors.Wrapf(types.ErrTotalShareLimit, "%v", err.Error())
		}
	}

	account := types.MainnetVestingAccount{
		CampaignID:     campaign.Id,
		Address:        msg.Address,
		Shares:         msg.Shares,
		VestingOptions: msg.VestingOptions,
	}

	// get the VestingOption and account shares
	totalShares, err := account.GetTotalShares()
	if err != nil {
		return nil, spnerrors.Critical(err.Error())
	}

	// increase the campaign shares
	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, totalShares)
	if types.IsTotalSharesReached(campaign.AllocatedShares, campaign.TotalShares) {
		return nil, sdkerrors.Wrapf(types.ErrTotalShareLimit, "%v", msg.CampaignID)
	}

	k.SetCampaign(ctx, campaign)
	k.SetMainnetVestingAccount(ctx, account)

	return &types.MsgAddMainnetVestingAccountResponse{}, nil
}
