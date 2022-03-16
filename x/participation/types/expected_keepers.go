package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
)

type FundraisingKeeper interface {
	GetAuction(ctx sdk.Context, id uint64) (auction fundraisingtypes.AuctionI, found bool)
	AddAllowedBidders(ctx sdk.Context, auctionID uint64, bidders []fundraisingtypes.AllowedBidder) error
}

type StakingKeeper interface {
	GetDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress,
		maxRetrieve uint16) []stakingtypes.Delegation
}
