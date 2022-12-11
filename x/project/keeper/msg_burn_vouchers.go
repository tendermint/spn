package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	"github.com/tendermint/spn/x/project/types"
)

func (k msgServer) BurnVouchers(goCtx context.Context, msg *types.MsgBurnVouchers) (*types.MsgBurnVouchersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	project, found := k.GetProject(ctx, msg.ProjectID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProjectNotFound, "%d", msg.ProjectID)
	}

	// Convert and validate vouchers first
	shares, err := types.VouchersToShares(msg.Vouchers, msg.ProjectID)
	if err != nil {
		return nil, ignterrors.Criticalf("verified voucher are invalid %s", err.Error())
	}

	// Send coins and burn them
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, ignterrors.Criticalf("can't parse sender address %s", err.Error())
	}
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, msg.Vouchers); err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientVouchers, "%s", err.Error())
	}
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Vouchers); err != nil {
		return nil, ignterrors.Criticalf("can't burn coins %s", err.Error())
	}

	// Decrease the project shares
	project.AllocatedShares, err = types.DecreaseShares(project.AllocatedShares, shares)
	if err != nil {
		return nil, ignterrors.Criticalf("invalid allocated share amount %s", err.Error())
	}
	k.SetProject(ctx, project)

	err = ctx.EventManager().EmitTypedEvent(&types.EventProjectSharesUpdated{
		ProjectID:      project.ProjectID,
		AllocatedShares: project.AllocatedShares,
	})

	return &types.MsgBurnVouchersResponse{}, err
}
