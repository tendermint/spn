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

	availableAlloc, err := k.GetAvailableAllocations(ctx, msg.Participant)
	if err != nil {
		return nil, err
	}

	// check if auction exists
	auction, found := k.fundraisingKeeper.GetAuction(ctx, msg.AuctionID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrAuctionNotFound, "auction %d not found", msg.AuctionID)
	}

	// check if auction is not in standby
	if auction.GetStatus() != fundraisingtypes.AuctionStatusStandBy {
		return nil, sdkerrors.Wrapf(types.ErrParticipationNotAllowed, "auction %d is not in standby", msg.AuctionID)
	}

	// check if auction allows participation at this time
	if !k.IsRegistrationEnabled(ctx, auction.GetStartTime()) {
		return nil, sdkerrors.Wrapf(types.ErrParticipationNotAllowed, "participation period for auction %d not yet started", msg.AuctionID)
	}

	// check if the user is already added as an allowed bidder for the auction
	_, found = k.GetAuctionUsedAllocations(ctx, msg.Participant, msg.AuctionID)
	if found {
		return nil, sdkerrors.Wrapf(types.ErrAlreadyParticipating,
			"address %s is already a participant for auction %d",
			msg.Participant, msg.AuctionID)
	}

	tiers := k.ParticipationTierList(ctx)
	tier, found := types.GetTierFromID(tiers, msg.TierID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrTierNotFound, "tier %d not found", msg.TierID)
	}

	// check if user has enough available allocations to cover tier
	if tier.RequiredAllocations.GT(availableAlloc) {
		return nil, sdkerrors.Wrapf(types.ErrInsufficientAllocations,
			"available allocations %d is less than required allocations %d for tier %d",
			availableAlloc, tier.RequiredAllocations, tier.TierID)
	}

	allowedBidder := fundraisingtypes.AllowedBidder{
		Bidder:       msg.Participant,
		MaxBidAmount: tier.Benefits.MaxBidAmount,
	}
	if err := k.fundraisingKeeper.AddAllowedBidders(
		ctx, msg.AuctionID,
		[]fundraisingtypes.AllowedBidder{allowedBidder},
	); err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidBidder, err.Error())
	}

	// set used allocations
	numUsedAllocations := sdkmath.ZeroInt()
	used, found := k.GetUsedAllocations(ctx, msg.Participant)
	if found {
		numUsedAllocations = used.NumAllocations
	}

	numUsedAllocations = numUsedAllocations.Add(tier.RequiredAllocations)
	k.SetUsedAllocations(ctx, types.UsedAllocations{
		Address:        msg.Participant,
		NumAllocations: numUsedAllocations,
	})

	// set auction used allocations
	k.SetAuctionUsedAllocations(ctx, types.AuctionUsedAllocations{
		Address:        msg.Participant,
		AuctionID:      msg.AuctionID,
		NumAllocations: tier.RequiredAllocations,
		Withdrawn:      false,
	})

	return &types.MsgParticipateResponse{}, ctx.EventManager().EmitTypedEvent(&types.EventAllocationsUsed{
		Participant:    msg.Participant,
		AuctionID:      msg.AuctionID,
		NumAllocations: tier.RequiredAllocations,
	})
}
