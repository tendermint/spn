package keeper

import (
	"context"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	ignterrors "github.com/ignite/modules/pkg/errors"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	"github.com/tendermint/spn/x/participation/types"
)

func (k msgServer) WithdrawAllocations(goCtx context.Context, msg *types.MsgWithdrawAllocations) (*types.MsgWithdrawAllocationsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	blockTime := ctx.BlockTime()

	auction, found := k.fundraisingKeeper.GetAuction(ctx, msg.AuctionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAuctionNotFound, "auction %d not found", msg.AuctionID)
	}

	// only prevent time-based restrictions on withdrawals if the auction's status is not `CANCELLED`
	if auction.GetStatus() != fundraisingtypes.AuctionStatusCancelled {
		withdrawalDelay := k.WithdrawalDelay(ctx)
		if !blockTime.After(auction.GetStartTime().Add(withdrawalDelay)) {
			return nil, sdkerrors.Wrapf(types.ErrAllocationWithdrawalTimeNotReached, "withdrawal for auction %d not yet allowed", msg.AuctionID)
		}
	}

	auctionUsedAllocations, found := k.GetAuctionUsedAllocations(ctx, msg.Participant, msg.AuctionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrUsedAllocationsNotFound, "used allocations for auction %d not found", msg.AuctionID)
	}
	if auctionUsedAllocations.Withdrawn {
		return nil, sdkerrors.Wrapf(types.ErrAllocationsAlreadyWithdrawn, "allocations for auction %d already claimed", msg.AuctionID)
	}

	totalUsedAllocations, found := k.GetUsedAllocations(ctx, msg.Participant)
	if !found {
		return nil, ignterrors.Criticalf("unable to find total used allocations entry for address %s", msg.Participant)
	}

	// decrease totalUsedAllocations making sure subtraction is feasible
	if totalUsedAllocations.NumAllocations.LT(auctionUsedAllocations.NumAllocations) {
		return nil, ignterrors.Critical("number of total used allocations cannot become negative")
	}
	totalUsedAllocations.NumAllocations = totalUsedAllocations.NumAllocations.Sub(auctionUsedAllocations.NumAllocations)

	auctionUsedAllocations.Withdrawn = true
	k.SetAuctionUsedAllocations(ctx, auctionUsedAllocations)
	k.SetUsedAllocations(ctx, totalUsedAllocations)

	return &types.MsgWithdrawAllocationsResponse{}, ctx.EventManager().EmitTypedEvent(&types.EventAllocationsWithdrawn{
		Participant: msg.Participant,
		AuctionID:   msg.AuctionID,
	})
}
