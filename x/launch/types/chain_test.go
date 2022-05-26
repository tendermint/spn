package types_test

import (
	"testing"

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

	for _, tc := range []struct {
		desc  string
		chain types.Chain
		valid bool
	}{
		{
			desc:  "valid chain",
			chain: sample.Chain(r, 0, 0),
			valid: true,
		},
		{
			desc:  "invalid genesis chain ID",
			chain: invalidGenesisChainID,
			valid: false,
		},
		{
			desc:  "invalid launch timestamp",
			chain: invalidLaunchTimestamp,
			valid: false,
		},
		{
			desc:  "mainnet without campaign",
			chain: mainnetWithoutCampaign,
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.chain.Validate()
			require.EqualValues(t, tc.valid, err == nil)
		})
	}
}
