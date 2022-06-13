package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
	"github.com/tendermint/spn/x/campaign/types"
	"time"
)

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
	auctionId uint64,
	auctioneer string,
	_ sdk.Dec,
	sellingCoin sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ time.Time,
	_ time.Time,
) {
	campaignID, err := types.VoucherCampaign(sellingCoin.Denom)
	if err != nil {
		// not a campaign auction
		return
	}

	// verify the auctioneer is the coordinator of the campaign
	campaign, found := h.campaignKeeper.GetCampaign(ctx, campaignID)
	if !found {
		return
	}
	coord, found := h.campaignKeeper.profileKeeper.GetCoordinator(ctx, campaign.CoordinatorID)
	if !found {
		return
	}

	// if the coordinator if the auctioneer, we emit a CampaignAuctionCreated event
	if coord.Address == auctioneer {
		_ = ctx.EventManager().EmitTypedEvents(
			&types.EventCampaignAuctionCreated{
				CampaignID: campaignID,
				AuctionID:  auctionId,
			},
		)
	}

	return
}

// AfterBatchAuctionCreated emits a CampaignAuctionCreated event if created for a campaign
func (h CampaignAuctionEventHooks) AfterBatchAuctionCreated(
	ctx sdk.Context,
	auctionId uint64,
	auctioneer string,
	startPrice sdk.Dec,
	minBidPrice sdk.Dec,
	sellingCoin sdk.Coin,
	payingCoinDenom string,
	vestingSchedules []fundraisingtypes.VestingSchedule,
	maxExtendedRound uint32,
	extendedRoundRate sdk.Dec,
	startTime time.Time,
	endTime time.Time,
) {
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
	_ sdk.Int,
) {
}

// BeforeSellingCoinsAllocated implements FundraisingHooks
func (h CampaignAuctionEventHooks) BeforeSellingCoinsAllocated(
	_ sdk.Context,
	_ uint64,
	_ map[string]sdk.Int,
	_ map[string]sdk.Int,
) {
}
