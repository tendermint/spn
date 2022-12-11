package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/project/types"
)

func (k msgServer) RedeemVouchers(goCtx context.Context, msg *types.MsgRedeemVouchers) (*types.MsgRedeemVouchersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	project, found := k.GetProject(ctx, msg.ProjectID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProjectNotFound, "%d", msg.ProjectID)
	}

	mainnetLaunched, err := k.IsProjectMainnetLaunchTriggered(ctx, project.ProjectID)
	if err != nil {
		return nil, ignterrors.Critical(err.Error())
	}
	if mainnetLaunched {
		return nil, sdkerrors.Wrap(types.ErrMainnetLaunchTriggered, fmt.Sprintf(
			"mainnet %d launch is already triggered",
			project.MainnetID,
		))
	}

	// Convert and validate vouchers first
	shares, err := types.VouchersToShares(msg.Vouchers, msg.ProjectID)
	if err != nil {
		return nil, ignterrors.Criticalf("verified voucher are invalid %s", err.Error())
	}

	// Send coins and burn them
	creatorAddr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, ignterrors.Criticalf("can't parse sender address %s", err.Error())
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddr, types.ModuleName, msg.Vouchers); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientVouchers, "%s", creatorAddr.String())
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Vouchers); err != nil {
		return nil, ignterrors.Criticalf("can't burn coins %s", err.Error())
	}

	// Check if the account already exists
	account, found := k.GetMainnetAccount(ctx, msg.ProjectID, msg.Account)
	if !found {
		// If not, create the account
		account = types.MainnetAccount{
			ProjectID: project.ProjectID,
			Address:   msg.Account,
			Shares:    types.EmptyShares(),
		}
	}

	// Increase the account shares
	account.Shares = types.IncreaseShares(account.Shares, shares)
	k.SetMainnetAccount(ctx, account)

	if !found {
		err = ctx.EventManager().EmitTypedEvent(&types.EventMainnetAccountCreated{
			ProjectID: account.ProjectID,
			Address:   account.Address,
			Shares:    account.Shares,
		})
	} else {
		err = ctx.EventManager().EmitTypedEvent(&types.EventMainnetAccountUpdated{
			ProjectID: account.ProjectID,
			Address:   account.Address,
			Shares:    account.Shares,
		})
	}

	return &types.MsgRedeemVouchersResponse{}, err
}
