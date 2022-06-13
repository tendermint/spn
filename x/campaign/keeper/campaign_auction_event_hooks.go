package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
	"github.com/tendermint/spn/x/campaign/types"
	"time"
)

// EmitCampaignAuctionCreated emits EventCampaignAuctionCreated event if an auction is created for a campaign from a coordinator
func (k Keeper) EmitCampaignAuctionCreated(
	ctx sdk.Context,
	auctionId uint64,
	auctioneer string,
	sellingCoin sdk.Coin,
) bool {
	campaignID, err := types.VoucherCampaign(sellingCoin.Denom)
	if err != nil {
		// not a campaign auction
		return false
	}

	// verify the auctioneer is the coordinator of the campaign
	campaign, found := k.GetCampaign(ctx, campaignID)
	if !found {
		return false
	}
	coord, found := k.profileKeeper.GetCoordinator(ctx, campaign.CoordinatorID)
	if !found {
		return false
	}

	// if the coordinator if the auctioneer, we emit a CampaignAuctionCreated event
	if coord.Address != auctioneer {
		return false
	}

	err = ctx.EventManager().EmitTypedEvents(
		&types.EventCampaignAuctionCreated{
			CampaignID: campaignID,
			AuctionID:  auctionId,
		},
	)
	if err != nil {
		return false
	}

	return true
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
	auctionId uint64,
	auctioneer string,
	_ sdk.Dec,
	sellingCoin sdk.Coin,
	_ string,
	_ []fundraisingtypes.VestingSchedule,
	_ time.Time,
	_ time.Time,
) {
	_ = h.campaignKeeper.EmitCampaignAuctionCreated(ctx, auctionId, auctioneer, sellingCoin)
}

// AfterBatchAuctionCreated emits a CampaignAuctionCreated event if created for a campaign
func (h CampaignAuctionEventHooks) AfterBatchAuctionCreated(
	ctx sdk.Context,
	auctionId uint64,
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
	_ = h.campaignKeeper.EmitCampaignAuctionCreated(ctx, auctionId, auctioneer, sellingCoin)
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
