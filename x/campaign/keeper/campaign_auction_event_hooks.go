package keeper

import (
	"cosmossdk.io/math"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"

	"github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

// EmitCampaignAuctionCreated emits EventCampaignAuctionCreated event if an auction is created for a campaign from a coordinator
func (k Keeper) EmitCampaignAuctionCreated(
	ctx sdk.Context,
	auctionID uint64,
	auctioneer string,
	sellingCoin sdk.Coin,
) (bool, error) {
	campaignID, err := types.VoucherCampaign(sellingCoin.Denom)
	if err != nil {
		// not a campaign auction
		return false, nil
	}

	// verify the auctioneer is the coordinator of the campaign
	campaign, found := k.GetCampaign(ctx, campaignID)
	if !found {
		return false, sdkerrors.Wrapf(types.ErrCampaignNotFound,
			"voucher %s is associated to an non-existing campaign %d",
			sellingCoin.Denom,
			campaignID,
		)
	}
	coord, found := k.profileKeeper.GetCoordinator(ctx, campaign.CoordinatorID)
	if !found {
		return false, sdkerrors.Wrapf(profiletypes.ErrCoordInvalid,
			"campaign %d coordinator doesn't exist %d",
			campaignID,
			campaign.CoordinatorID,
		)
	}

	// if the coordinator if the auctioneer, we emit a CampaignAuctionCreated event
	if coord.Address != auctioneer {
		return false, nil
	}

	err = ctx.EventManager().EmitTypedEvents(
		&types.EventCampaignAuctionCreated{
			CampaignID: campaignID,
			AuctionID:  auctionID,
		},
	)
	if err != nil {
		return false, err
	}

	return true, nil
}

// CampaignAuctionEventHooks returns a CampaignAuctionEventHooks associated with the campaign keeper
func (k Keeper) CampaignAuctionEventHooks() CampaignAuctionEventHooks {
	return CampaignAuctionEventHooks{
		campaignKeeper: k,
	}
}

// CampaignAuctionEventHooks implements fundraising hooks and emit events on auction creation
type CampaignAuctionEventHooks struct {
	campaignKeeper Keeper
}

// Implements FundraisingHooks interface
var _ fundraisingtypes.FundraisingHooks = CampaignAuctionEventHooks{}

// AfterFixedPriceAuctionCreated emits a CampaignAuctionCreated event if created for a campaign
func (h CampaignAuctionEventHooks) AfterFixedPriceAuctionCreated(
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
	_, _ = h.campaignKeeper.EmitCampaignAuctionCreated(ctx, auctionID, auctioneer, sellingCoin)
}

// AfterBatchAuctionCreated emits a CampaignAuctionCreated event if created for a campaign
func (h CampaignAuctionEventHooks) AfterBatchAuctionCreated(
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
	_, _ = h.campaignKeeper.EmitCampaignAuctionCreated(ctx, auctionID, auctioneer, sellingCoin)
}

// BeforeFixedPriceAuctionCreated implements FundraisingHooks
func (h CampaignAuctionEventHooks) BeforeFixedPriceAuctionCreated(
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
func (h CampaignAuctionEventHooks) BeforeBatchAuctionCreated(
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
func (h CampaignAuctionEventHooks) BeforeAuctionCanceled(
	_ sdk.Context,
	_ uint64,
	_ string,
) {
}

// BeforeBidPlaced implements FundraisingHooks
func (h CampaignAuctionEventHooks) BeforeBidPlaced(
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
func (h CampaignAuctionEventHooks) BeforeBidModified(
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
func (h CampaignAuctionEventHooks) BeforeAllowedBiddersAdded(
	_ sdk.Context,
	_ []fundraisingtypes.AllowedBidder,
) {
}

// BeforeAllowedBidderUpdated implements FundraisingHooks
func (h CampaignAuctionEventHooks) BeforeAllowedBidderUpdated(
	_ sdk.Context,
	_ uint64,
	_ sdk.AccAddress,
	_ math.Int,
) {
}

// BeforeSellingCoinsAllocated implements FundraisingHooks
func (h CampaignAuctionEventHooks) BeforeSellingCoinsAllocated(
	_ sdk.Context,
	_ uint64,
	_ map[string]math.Int,
	_ map[string]math.Int,
) {
}
