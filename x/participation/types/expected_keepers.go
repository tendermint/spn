package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
)

type FundraisingKeeper interface {
	GetAuction(ctx sdk.Context, id uint64) (auction fundraisingtypes.AuctionI, found bool)
	AddAllowedBidders(ctx sdk.Context, auctionId uint64, bidders []fundraisingtypes.AllowedBidder) error
}

type StakingKeeper interface {
	// Methods imported from staking should be defined here
}
