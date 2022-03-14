package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	"github.com/tendermint/spn/x/participation/types"
)

func (k msgServer) Participate(goCtx context.Context, msg *types.MsgParticipate) (*types.MsgParticipateResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Read available allocation

	_, found := k.fundraisingKeeper.GetAuction(ctx, msg.AuctionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAuctionNotFound, "auction %d not found", msg.AuctionID)
	}

	allowedBidder := fundraisingtypes.AllowedBidder{
		Bidder:       msg.Participant,
		MaxBidAmount: sdk.NewIntFromUint64(1000),
	}
	if err := k.fundraisingKeeper.AddAllowedBidders(
		ctx, msg.AuctionID,
		[]fundraisingtypes.AllowedBidder{allowedBidder},
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBidder, err.Error())
	}

	return &types.MsgParticipateResponse{}, nil
}
