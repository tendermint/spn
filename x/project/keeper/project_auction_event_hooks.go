package keeper

import (
	"time"

	sdkerrors "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	profiletypes "github.com/tendermint/spn/x/profile/types"
	"github.com/tendermint/spn/x/project/types"
)

// EmitProjectAuctionCreated emits EventProjectAuctionCreated event if an auction is created for a project from a coordinator
func (k Keeper) EmitProjectAuctionCreated(
	ctx sdk.Context,
	auctionID uint64,
	auctioneer string,
	sellingCoin sdk.Coin,
) (bool, error) {
	projectID, err := types.VoucherProject(sellingCoin.Denom)
	if err != nil {
		// not a project auction
		return false, nil
	}

	// verify the auctioneer is the coordinator of the project
	project, found := k.GetProject(ctx, projectID)
	if !found {
		return false, sdkerrors.Wrapf(types.ErrProjectNotFound,
			"voucher %s is associated to an non-existing project %d",
			sellingCoin.Denom,
			projectID,
		)
	}
	coord, found := k.profileKeeper.GetCoordinator(ctx, project.CoordinatorID)
	if !found {
		return false, sdkerrors.Wrapf(profiletypes.ErrCoordInvalid,
			"project %d coordinator doesn't exist %d",
			projectID,
			project.CoordinatorID,
		)
	}

	// if the coordinator if the auctioneer, we emit a ProjectAuctionCreated event
	if coord.Address != auctioneer {
		return false, nil
	}

	err = ctx.EventManager().EmitTypedEvents(
		&types.EventProjectAuctionCreated{
			ProjectID: projectID,
			AuctionID: auctionID,
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

// ProjectAuctionEventHooks returns a ProjectAuctionEventHooks associated with the project keeper
func (k Keeper) ProjectAuctionEventHooks() ProjectAuctionEventHooks {
	return ProjectAuctionEventHooks{
		projectKeeper: k,
	}
}

// ProjectAuctionEventHooks implements fundraising hooks and emit events on auction creation
type ProjectAuctionEventHooks struct {
	projectKeeper Keeper
}

// Implements FundraisingHooks interface
var _ fundraisingtypes.FundraisingHooks = ProjectAuctionEventHooks{}

// AfterFixedPriceAuctionCreated emits a ProjectAuctionCreated event if created for a project
func (h ProjectAuctionEventHooks) AfterFixedPriceAuctionCreated(
	ctx sdk.Context,
	auctionID uint64,
	auctioneer string,
	_ sdk.Dec,
	sellingCoin sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ time.Time,
	_ time.Time,
) {
	// TODO: investigate error handling for hooks
	// https://github.com/tendermint/spn/issues/869
	_, _ = h.projectKeeper.EmitProjectAuctionCreated(ctx, auctionID, auctioneer, sellingCoin)
}

// AfterBatchAuctionCreated emits a ProjectAuctionCreated event if created for a project
func (h ProjectAuctionEventHooks) AfterBatchAuctionCreated(
	ctx sdk.Context,
	auctionID uint64,
	auctioneer string,
	_ sdk.Dec,
	_ sdk.Dec,
	sellingCoin sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ uint32,
	_ sdk.Dec,
	_ time.Time,
	_ time.Time,
) {
	// TODO: investigate error handling for hooks
	// https://github.com/tendermint/spn/issues/869
	_, _ = h.projectKeeper.EmitProjectAuctionCreated(ctx, auctionID, auctioneer, sellingCoin)
}

// BeforeFixedPriceAuctionCreated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeFixedPriceAuctionCreated(
	_ sdk.Context,
	_ string,
	_ sdk.Dec,
	_ sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ time.Time,
	_ time.Time,
) {
}

// BeforeBatchAuctionCreated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeBatchAuctionCreated(
	_ sdk.Context,
	_ string,
	_ sdk.Dec,
	_ sdk.Dec,
	_ sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ uint32,
	_ sdk.Dec,
	_ time.Time,
	_ time.Time,
) {
}

// BeforeAuctionCanceled implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeAuctionCanceled(
	_ sdk.Context,
	_ uint64,
	_ string,
) {
}

// BeforeBidPlaced implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeBidPlaced(
	_ sdk.Context,
	_ uint64,
	_ uint64,
	_ string,
	_ fundraisingtypes.BidType,
	_ sdk.Dec,
	_ sdk.Coin,
) {
}

// BeforeBidModified implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeBidModified(
	_ sdk.Context,
	_ uint64,
	_ uint64,
	_ string,
	_ fundraisingtypes.BidType,
	_ sdk.Dec,
	_ sdk.Coin,
) {
}

// BeforeAllowedBiddersAdded implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeAllowedBiddersAdded(
	_ sdk.Context,
	_ []fundraisingtypes.AllowedBidder,
) {
}

// BeforeAllowedBidderUpdated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeAllowedBidderUpdated(
	_ sdk.Context,
	_ uint64,
	_ sdk.AccAddress,
	_ sdkmath.Int,
) {
}

// BeforeSellingCoinsAllocated implements FundraisingHooks
func (h ProjectAuctionEventHooks) BeforeSellingCoinsAllocated(
	_ sdk.Context,
	_ uint64,
	_ map[string]sdkmath.Int,
	_ map[string]sdkmath.Int,
) {
}
