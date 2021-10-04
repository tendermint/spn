package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/campaign/types"
)

func (k msgServer) BurnVouchers(goCtx context.Context, msg *types.MsgBurnVouchers) (*types.MsgBurnVouchersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	campaign, found := k.GetCampaign(ctx, msg.CampaignID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrCampaignNotFound, "%d", msg.CampaignID)
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, spnerrors.Criticalf("can't parse creator address %s", err.Error())
	}

	// transfer vouchers
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, msg.Vouchers); err != nil {
		return nil, spnerrors.Criticalf("can't send burned coins %s", err.Error())
	}

	// Burn vouchers
	if err := k.bankKeeper.BurnCoins(ctx, types.ModuleName, msg.Vouchers); err != nil {
		return nil, sdkerrors.Wrap(types.ErrVouchersBurn, err.Error())
	}

	shares, err := types.VouchersToShares(msg.Vouchers, msg.CampaignID)
	if err != nil {
		return nil, spnerrors.Criticalf("verified voucher are invalid %s", err.Error())
	}

	// Decrease the campaign shares
	campaign.AllocatedShares, err = types.DecreaseShares(campaign.AllocatedShares, shares)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidShares, "%d", msg.CampaignID)
	}
	k.SetCampaign(ctx, campaign)

	return &types.MsgBurnVouchersResponse{}, nil
}
