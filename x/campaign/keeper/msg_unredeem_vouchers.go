package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) UnredeemVouchers(goCtx context.Context, msg *types.MsgUnredeemVouchers) (*types.MsgUnredeemVouchersResponse, error) {
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

	account, found := k.GetMainnetAccount(ctx, msg.CampaignID, msg.Sender)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrAccountNotFound, msg.Sender)
	}

	// Update the shares of the account
	account.Shares, err = types.DecreaseShares(account.Shares, msg.Shares)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrSharesDecrease, err.Error())
	}

	// If the account no longer has shares, it can be removed from the store
	if types.IsEqualShares(account.Shares, types.EmptyShares()) {
		k.RemoveMainnetAccount(ctx, msg.CampaignID, msg.Sender)
		if err := ctx.EventManager().EmitTypedEvent(&types.EventMainnetAccountRemoved{
			CampaignID: campaign.CampaignID,
			Address:    account.Address,
		}); err != nil {
			return nil, err
		}
	} else {
		k.SetMainnetAccount(ctx, account)
		if err := ctx.EventManager().EmitTypedEvent(&types.EventMainnetAccountUpdated{
			CampaignID: account.CampaignID,
			Address:    account.Address,
			Shares:     account.Shares,
		}); err != nil {
			return nil, err
		}
	}

	// Mint vouchers from the removed shares and send them to sender balance
	vouchers, err := types.SharesToVouchers(msg.Shares, msg.CampaignID)
	if err != nil {
		return nil, spnerrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, vouchers); err != nil {
		return nil, sdkerrors.Wrap(types.ErrVouchersMinting, err.Error())
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, spnerrors.Criticalf("can't parse sender address %s", err.Error())
	}

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, vouchers); err != nil {
		return nil, spnerrors.Criticalf("can't send minted coins %s", err.Error())
	}

	return &types.MsgUnredeemVouchersResponse{}, nil
}
