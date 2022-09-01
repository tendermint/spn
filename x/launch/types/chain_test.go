package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestChain_Validate(t *testing.T) {
	invalidGenesisChainID := sample.Chain(r, 0, 0)
	invalidGenesisChainID.GenesisChainID = "invalid"

	invalidLaunchTimestamp := sample.Chain(r, 0, 0)
	invalidLaunchTimestamp.LaunchTriggered = true

	mainnetWithoutCampaign := sample.Chain(r, 0, 0)
	mainnetWithoutCampaign.IsMainnet = true

	invalidCoins := sample.Chain(r, 0, 0)
	// add invalid coin amount
	invalidCoins.AccountBalance = sdk.Coins{sdk.Coin{Denom: "invalid", Amount: sdk.NewInt(-1)}}

	for _, tc := range []struct {
		desc  string
		chain types.Chain
		valid bool
	}{
		{
			desc:  "should validate valid chain",
			chain: sample.Chain(r, 0, 0),
			valid: true,
		},
		{
			desc:  "should prevent validate invalid genesis chain ID",
			chain: invalidGenesisChainID,
			valid: false,
		},
		{
			desc:  "should prevent validate invalid launch timestamp",
			chain: invalidLaunchTimestamp,
			valid: false,
		},
		{
			desc:  "should prevent validate mainnet chain without associated campaign ID",
			chain: mainnetWithoutCampaign,
			valid: false,
		},
		{
			desc:  "should prevent chain with invalid coins structure",
			chain: invalidCoins,
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.chain.Validate()
			require.EqualValues(t, tc.valid, err == nil)
		})
	}
}
