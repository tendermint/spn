package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/fundraising/x/fundraising/types"
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
var _ types.FundraisingHooks = CampaignAuctionEventHooks{}

// BeforeFixedPriceAuctionCreated emits a CampaignAuctionCreated event
func (h CampaignAuctionEventHooks) BeforeFixedPriceAuctionCreated(
	_ sdk.Context,
	_ string,
	_ sdk.Dec,
	_ sdk.Coin,
	_ string,
	_ []types.VestingSchedule,
	_ time.Time,
	_ time.Time,
) {
}

// BeforeBatchAuctionCreated emits a CampaignAuctionCreated event
func (h CampaignAuctionEventHooks) BeforeBatchAuctionCreated(
	_ sdk.Context,
	_ string,
	_ sdk.Dec,
	_ sdk.Dec,
	_ sdk.Coin,
	_ string,
	_ []types.VestingSchedule,
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
	_ types.BidType,
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
	_ types.BidType,
	_ sdk.Dec,
	_ sdk.Coin,
) {
}

// BeforeAllowedBiddersAdded implements FundraisingHooks
func (h CampaignAuctionEventHooks) BeforeAllowedBiddersAdded(
	_ sdk.Context,
	_ []types.AllowedBidder,
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
