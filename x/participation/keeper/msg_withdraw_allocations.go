package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/participation/types"
)

func (k msgServer) WithdrawAllocations(goCtx context.Context, msg *types.MsgWithdrawAllocations) (*types.MsgWithdrawAllocationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	auction, found := k.fundraisingKeeper.GetAuction(ctx, msg.AuctionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAuctionNotFound, "auction %d not found", msg.AuctionID)
	}

	blockTime := ctx.BlockTime()
	if !auction.IsAuctionStarted(blockTime) {
		return nil, sdkerrors.Wrapf(types.ErrAllocationsLocked, "auction %d not yet started", msg.AuctionID)
	}

	// TODO check delay is reached

	auctionUsedAllocations, found := k.GetAuctionUsedAllocations(ctx, msg.Participant, msg.AuctionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUsedAllocationsNotFound, "used allocations for auction %d not found", msg.AuctionID)
	}
	if auctionUsedAllocations.Withdrawn {
		return nil, sdkerrors.Wrapf(types.ErrAllocationsAlreadyWithdrawn, "allocations for auction %d already claimed", msg.AuctionID)
	}

	totalUsedAllocations, found := k.GetUsedAllocations(ctx, msg.Participant)
	if !found {
		return nil, spnerrors.Critical(fmt.Sprintf("unable to find total used allocations entry for address %s", msg.Participant))
	}

	// decrease totalUsedAllocations making sure subtraction is feasible
	if totalUsedAllocations.NumAllocations < auctionUsedAllocations.NumAllocations {
		return nil, spnerrors.Critical("number of total used allocations cannot become negative")
	}
	totalUsedAllocations.NumAllocations -= auctionUsedAllocations.NumAllocations

	auctionUsedAllocations.Withdrawn = true
	k.SetAuctionUsedAllocations(ctx, auctionUsedAllocations)
	k.SetUsedAllocations(ctx, totalUsedAllocations)

	return &types.MsgWithdrawAllocationsResponse{}, nil
}
