package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) RedeemVouchers(goCtx context.Context, msg *types.MsgRedeemVouchers) (*types.MsgRedeemVouchersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	if campaign.MainnetInitialized {
		mainnetLaunch, found := k.launchKeeper.GetChain(ctx, campaign.MainnetID)
		if !found {
			return nil, spnerrors.Criticalf("cannot find mainnet chain %d for campaign %d", campaign.MainnetID, campaign.CampaignID)
		}
		if mainnetLaunch.LaunchTriggered {
			return nil, sdkerrors.Wrap(types.ErrMainnetLaunchTriggered, fmt.Sprintf(
				"mainnet %d launched, cannot  add shares",
				campaign.MainnetID,
			))
		}
	}

	// Convert and validate vouchers first
	shares, err := types.VouchersToShares(msg.Vouchers, msg.CampaignID)
	if err != nil {
		return nil, spnerrors.Criticalf("verified voucher are invalid %s", err.Error())
	}

	// Send coins and burn them
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, spnerrors.Criticalf("can't parse sender address %s", err.Error())
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, msg.Vouchers); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientVouchers, "%s", creatorAddr.String())
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Vouchers); err != nil {
		return nil, spnerrors.Criticalf("can't burn coins %s", err.Error())
	}

	// Check if the account already exists
	account, found := k.GetMainnetAccount(ctx, msg.CampaignID, msg.Account)
	if !found {
		// If not, create the account
		account = types.MainnetAccount{
			CampaignID: campaign.CampaignID,
			Address:    msg.Account,
			Shares:     types.EmptyShares(),
		}
	}
	// Increase the account shares
	account.Shares = types.IncreaseShares(account.Shares, shares)
	k.SetMainnetAccount(ctx, account)

	return &types.MsgRedeemVouchersResponse{}, nil
}
