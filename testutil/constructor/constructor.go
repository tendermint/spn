// Package constructor provides constructors to easily initialize objects for test purpose with automatic error handling
package constructor

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	campaigntypes "github.com/tendermint/spn/x/project/types"
	monitoringptypes "github.com/tendermint/spn/x/monitoringp/types"
)

// Vote is a simplified type for abci.VoteInfo for testing purpose
type Vote struct {
	Address []byte
	Signed  bool
}

// LastCommitInfo creates a ABCI LastCommitInfo object for test purpose from a list of vote
func LastCommitInfo(votes ...Vote) abci.LastCommitInfo {
	var lci abci.LastCommitInfo

	// add votes
	for _, vote := range votes {
		lci.Votes = append(lci.Votes, abci.VoteInfo{
			Validator: abci.Validator{
				Address: vote.Address,
			},
			SignedLastBlock: vote.Signed,
		})
	}
	return lci
}

// Coin returns a sdk.Coin from a string
func Coin(t testing.TB, str string) sdk.Coin {
	coin, err := sdk.ParseCoinNormalized(str)
	require.NoError(t, err)
	return coin
}

// Coins returns a sdk.Coins from a string
func Coins(t testing.TB, str string) sdk.Coins {
	coins, err := sdk.ParseCoinsNormalized(str)
	require.NoError(t, err)
	return coins
}

// Dec returns a sdk.Dec from a string
func Dec(t testing.TB, str string) sdk.Dec {
	dec, err := sdk.NewDecFromStr(str)
	require.NoError(t, err)
	return dec
}

// SignatureCount returns a signature count object for test from a operator address and a decimal string for relative signatures
func SignatureCount(t testing.TB, opAddr string, relSig string) spntypes.SignatureCount {
	return spntypes.SignatureCount{
		OpAddress:          opAddr,
		RelativeSignatures: Dec(t, relSig),
	}
}

// SignatureCounts returns a signature counts object for tests from a a block count and list of signature counts
func SignatureCounts(blockCount uint64, sc ...spntypes.SignatureCount) spntypes.SignatureCounts {
	return spntypes.SignatureCounts{
		BlockCount: blockCount,
		Counts:     sc,
	}
}

// MonitoringInfo returns a monitoring info object for tests from a a block count and list of signature counts
func MonitoringInfo(blockCount uint64, sc ...spntypes.SignatureCount) (mi monitoringptypes.MonitoringInfo) {
	mi.SignatureCounts = SignatureCounts(blockCount, sc...)
	return
}

// Shares returns a Shares object from a string of coin inputs
func Shares(t testing.TB, coinStr string) campaigntypes.Shares {
	shares := campaigntypes.NewSharesFromCoins(Coins(t, coinStr))
	return shares
}

// Vouchers returns a Vouchers object from a string of coin inputs
func Vouchers(t testing.TB, coinStr string, campaignID uint64) sdk.Coins {
	coins := Coins(t, coinStr)
	vouchers := make(sdk.Coins, len(coins))
	for i, coin := range coins {
		coin.Denom = campaigntypes.VoucherDenom(campaignID, coin.Denom)
		vouchers[i] = coin
	}
	return vouchers
}
