package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func (k msgServer) MintVouchers(goCtx context.Context, msg *types.MsgMintVouchers) (*types.MsgMintVouchersResponse, error) {
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

	// Increase the campaign shares
	campaign.AllocatedShares = types.IncreaseShares(campaign.AllocatedShares, msg.Shares)
	reached, err := types.IsTotalSharesReached(campaign.AllocatedShares, k.GetTotalShares(ctx))
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if reached {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.CampaignID)
	}

	// Mint vouchers to the coordinator account
	vouchers, err := types.SharesToVouchers(msg.Shares, msg.CampaignID)
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, vouchers); err != nil {
		return nil, sdkerrors.Wrap(types.ErrVouchersMinting, err.Error())
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return nil, ignterrors.Criticalf("can't parse coordinator address %s", err.Error())
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, vouchers); err != nil {
		return nil, ignterrors.Criticalf("can't send minted coins %s", err.Error())
	}

	k.SetCampaign(ctx, campaign)

	err = ctx.EventManager().EmitTypedEvent(
		&types.EventCampaignSharesUpdated{
			CampaignID:         campaign.CampaignID,
			CoordinatorAddress: msg.Coordinator,
			AllocatedShares:    campaign.AllocatedShares,
		})

	return &types.MsgMintVouchersResponse{}, err
}
