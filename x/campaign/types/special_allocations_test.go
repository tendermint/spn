package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/campaign/types"
)

func TestNewSpecialAllocation(t *testing.T) {
	genesisDistribution := sample.Shares(r)
	claimableAirdrop := sample.Shares(r)

	sa := types.NewSpecialAllocations(genesisDistribution, claimableAirdrop)
	require.True(t, types.IsEqualShares(sa.GenesisDistribution, genesisDistribution))
	require.True(t, types.IsEqualShares(sa.ClaimableAirdrop, claimableAirdrop))
}

func TestEmptySpecialAllocations(t *testing.T) {
	sa := types.EmptySpecialAllocations()
	require.True(t, sa.GenesisDistribution.Empty())
	require.True(t, sa.ClaimableAirdrop.Empty())
}

func TestSpecialAllocations_Validate(t *testing.T) {
	invalidShares := types.Shares(sdk.NewCoins(
		sdk.NewCoin("foo", sdk.NewInt(100)),
		sdk.NewCoin("s/bar", sdk.NewInt(200)),
	))
	require.Error(t, types.CheckShares(invalidShares))

	for _, tc := range []struct {
		desc  string
		sa    types.SpecialAllocations
		valid bool
	}{
		{
			desc:  "should validate empty",
			sa:    types.EmptySpecialAllocations(),
			valid: true,
		},
		{
			desc:  "should validate valid special allocations",
			sa:    types.NewSpecialAllocations(sample.Shares(r), sample.Shares(r)),
			valid: true,
		},
		{
			desc:  "should prevent invalid genesis distribution",
			sa:    types.NewSpecialAllocations(invalidShares, sample.Shares(r)),
			valid: false,
		},
		{
			desc:  "should prevent invalid claimable airdrop",
			sa:    types.NewSpecialAllocations(sample.Shares(r), invalidShares),
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.sa.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestSpecialAllocations_TotalShares(t *testing.T) {
	// No share for empty special allocations
	sa := types.EmptySpecialAllocations()
	require.True(t, sa.TotalShares().Empty())

	// Sum of genesis distribution and claimable airdrop
	genesisDistribution := sample.Shares(r)
	claimableAirdrop := sample.Shares(r)
	sa = types.NewSpecialAllocations(genesisDistribution, claimableAirdrop)
	require.True(t, types.IsEqualShares(
		sa.TotalShares(),
		types.IncreaseShares(sa.GenesisDistribution, sa.ClaimableAirdrop),
	))
}
