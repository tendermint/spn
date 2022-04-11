package keeper

import (
	"context"
	"fmt"

	spnerrors "github.com/tendermint/spn/pkg/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) AddShares(goCtx context.Context, msg *types.MsgAddShares) (*types.MsgAddSharesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	if campaign.MainnetInitialized {
		mainnetChain, found := k.launchKeeper.GetChain(ctx, campaign.MainnetID)
		if !found {
			return nil, spnerrors.Criticalf("cannot find mainnet chain %d for campaign %d", campaign.MainnetID, campaign.CampaignID)
		}
		if mainnetChain.LaunchTriggered {
			return nil, sdkerrors.Wrap(types.ErrMainnetLaunchTriggered, fmt.Sprintf(
				"mainnet %d is already launched, action prohibited",
				campaign.MainnetID,
			))
		}
	}

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
	account, found := k.GetMainnetAccount(ctx, campaign.CampaignID, msg.Address)
	if !found {
		// if not, create the account
		account = types.MainnetAccount{
			CampaignID: campaign.CampaignID,
			Address:    msg.Address,
			Shares:     types.EmptyShares(),
		}
	}
	// increase the account shares
	account.Shares = types.IncreaseShares(account.Shares, msg.Shares)

	// increase the campaign shares
	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, msg.Shares)
	if types.IsTotalSharesReached(campaign.AllocatedShares, k.GetTotalShares(ctx)) {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.CampaignID)
	}

	k.SetCampaign(ctx, campaign)
	k.SetMainnetAccount(ctx, account)

	return &types.MsgAddSharesResponse{}, nil
}
