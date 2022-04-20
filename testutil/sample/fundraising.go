package sample

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	fundraisingtypes "github.com/tendermint/fundraising/x/fundraising/types"
)

// MsgCreateFixedAuction create a sample MsgCreateFixedAuction message
func MsgCreateFixedAuction(
	r *rand.Rand,
	auctioneer string,
	sellingCoin sdk.Coin,
	startTime,
	endTime time.Time,
) *fundraisingtypes.MsgCreateFixedPriceAuction {
	sellingPrice := int64(r.Intn(10000)) + 10000 // 10000 - 20000

	return &fundraisingtypes.MsgCreateFixedPriceAuction{
		Auctioneer:       auctioneer,
		StartPrice:       sdk.NewDec(sellingPrice),
		SellingCoin:      sellingCoin,
		PayingCoinDenom:  stakingtypes.DefaultParams().BondDenom,
		VestingSchedules: []fundraisingtypes.VestingSchedule{},
		StartTime:        startTime,
		EndTime:          endTime,
	}
}
