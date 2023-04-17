package keeper

import (
	"context"
	"fmt"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"

	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/types"
)

func (k msgServer) MintVouchers(goCtx context.Context, msg *types.MsgMintVouchers) (*types.MsgMintVouchersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	project, found := k.GetProject(ctx, msg.ProjectID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrProjectNotFound, "%d", msg.ProjectID)
	}

	// Get the coordinator ID associated to the sender address
	coordID, err := k.profileKeeper.CoordinatorIDFromAddress(ctx, msg.Coordinator)
	if err != nil {
		return nil, err
	}

	if project.CoordinatorID != coordID {
		return nil, sdkerrors.Wrap(profiletypes.ErrCoordInvalid, fmt.Sprintf(
			"coordinator of the project is %d",
			project.CoordinatorID,
		))
	}

	// Increase the project shares
	project.AllocatedShares = types.IncreaseShares(project.AllocatedShares, msg.Shares)
	reached, err := types.IsTotalSharesReached(project.AllocatedShares, k.GetTotalShares(ctx))
	if err != nil {
		return nil, ignterrors.Criticalf("verified shares are invalid %s", err.Error())
	}
	if reached {
		return nil, sdkerrors.Wrapf(types.ErrTotalSharesLimit, "%d", msg.ProjectID)
	}

	// Mint vouchers to the coordinator account
	vouchers, err := types.SharesToVouchers(msg.Shares, msg.ProjectID)
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

	k.SetProject(ctx, project)

	err = ctx.EventManager().EmitTypedEvent(
		&types.EventProjectSharesUpdated{
			ProjectID:          project.ProjectID,
			CoordinatorAddress: msg.Coordinator,
			AllocatedShares:    project.AllocatedShares,
		})

	return &types.MsgMintVouchersResponse{}, err
}
