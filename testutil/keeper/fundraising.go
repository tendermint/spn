package keeper

import (
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
)

// CreateFixedPriceAuction makes the provided address create a fixed price auction with the specified selling coin and
// start time. Returns the ID of the created auction.
func (tk TestKeepers) CreateFixedPriceAuction(
	ctx sdk.Context,
	r *rand.Rand,
	auctioneer string,
	sellingCoin sdk.Coin,
	startTime,
	endTime time.Time,
) uint64 {
	res, err := tk.FundraisingKeeper.CreateFixedPriceAuction(ctx, sample.MsgCreateFixedAuction(r,
		auctioneer,
		sellingCoin,
		startTime,
		endTime,
	))
	require.NoError(tk.T, err)
	require.NotNil(tk.T, res)
	require.NotNil(tk.T, res.BaseAuction)
	return res.BaseAuction.Id
}
