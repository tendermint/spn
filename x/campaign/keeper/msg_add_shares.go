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

func (k msgServer) AddShares(goCtx context.Context, msg *types.MsgAddShares) (*types.MsgAddSharesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

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
	account, accountFound := k.GetMainnetAccount(ctx, campaign.CampaignID, msg.Address)
	if !accountFound {
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

	if !accountFound {
		err = ctx.EventManager().EmitTypedEvents(
			&types.EventCampaignSharesUpdated{
				CampaignID:         campaign.CampaignID,
				CoordinatorAddress: msg.Coordinator,
				AllocatedShares:    campaign.AllocatedShares,
			}, &types.EventMainnetAccountCreated{
				CampaignID: campaign.CampaignID,
				Address:    msg.Address,
				Shares:     msg.Shares,
			},
		)
	} else {
		err = ctx.EventManager().EmitTypedEvents(
			&types.EventCampaignSharesUpdated{
				CampaignID:         campaign.CampaignID,
				CoordinatorAddress: msg.Coordinator,
				AllocatedShares:    campaign.AllocatedShares,
			}, &types.EventMainnetAccountUpdated{
				CampaignID: campaign.CampaignID,
				Address:    msg.Address,
				Shares:     msg.Shares,
			},
		)
	}

	return &types.MsgAddSharesResponse{}, err
}
