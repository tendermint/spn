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

func (k msgServer) AddVestingOptions(goCtx context.Context, msg *types.MsgAddVestingOptions) (*types.MsgAddVestingOptionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	// Get the coordinator ID associated to the sender address
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

	// check if the account already exists
	oldAccount, foundAcc := k.GetMainnetVestingAccount(ctx, campaign.CampaignID, msg.Address)
	if foundAcc {
		// if yes, remove the allocated share
		totalShares, err := oldAccount.GetTotalShares()
		if err != nil {
			return nil, spnerrors.Critical(err.Error())
		}
		campaign.AllocatedShares, err = types.DecreaseShares(campaign.AllocatedShares, totalShares)
		if err != nil {
			return nil, spnerrors.Critical("campaign allocated shares is lower than all shares")
		}
	}

	account := types.MainnetVestingAccount{
		CampaignID:     campaign.CampaignID,
		Address:        msg.Address,
		VestingOptions: msg.VestingOptions,
	}

	// get the VestingOption and account shares
	totalShares, err := account.GetTotalShares()
	if err != nil {
		return nil, spnerrors.Criticalf(
			"fail to get the account and vesting option shares: %s",
			err.Error(),
		)
	}

	// increase the campaign shares
	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, totalShares)
	if types.IsTotalSharesReached(campaign.AllocatedShares, k.GetTotalShares(ctx)) {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.CampaignID)
	}

	k.SetCampaign(ctx, campaign)
	k.SetMainnetVestingAccount(ctx, account)

	return &types.MsgAddVestingOptionsResponse{}, nil
}
