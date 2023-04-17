package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/project/types"
)

func TestNewSpecialAllocation(t *testing.T) {
	genesisDistribution := sample.Shares(r)
	claimableAirdrop := sample.Shares(r)

	t.Run("should allow creation of new special allocations", func(t *testing.T) {
		sa := types.NewSpecialAllocations(genesisDistribution, claimableAirdrop)
		require.True(t, types.IsEqualShares(sa.GenesisDistribution, genesisDistribution))
		require.True(t, types.IsEqualShares(sa.ClaimableAirdrop, claimableAirdrop))
	})
}

func TestEmptySpecialAllocations(t *testing.T) {
	t.Run("should allow creation of new empty special allocations", func(t *testing.T) {
		sa := types.EmptySpecialAllocations()
		require.True(t, sa.GenesisDistribution.Empty())
		require.True(t, sa.ClaimableAirdrop.Empty())
	})
}

func TestSpecialAllocations_Validate(t *testing.T) {
	t.Run("should prevent check of invalid shares", func(t *testing.T) {
		invalidShares := types.Shares(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(100)),
			sdk.NewCoin("s/bar", sdkmath.NewInt(200)),
		))
		require.Error(t, types.CheckShares(invalidShares))
	})

	invalidShares := types.Shares{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}

	for _, tc := range []struct {
		desc  string
		sa    types.SpecialAllocations
		valid bool
	}{
		{
			desc:  "should allow validation of empty",
			sa:    types.EmptySpecialAllocations(),
			valid: true,
		},
		{
			desc:  "should allow validation of valid special allocations",
			sa:    types.NewSpecialAllocations(sample.Shares(r), sample.Shares(r)),
			valid: true,
		},
		{
			desc:  "should prevent validation of invalid genesis distribution",
			sa:    types.NewSpecialAllocations(invalidShares, sample.Shares(r)),
			valid: false,
		},
		{
			desc:  "should prevent validation of invalid claimable airdrop",
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
	t.Run("should verify special allocations are empty", func(t *testing.T) {
		sa := types.EmptySpecialAllocations()
		require.True(t, sa.TotalShares().Empty())
	})

	t.Run("should increase shares", func(t *testing.T) {
		genesisDistribution := sample.Shares(r)
		claimableAirdrop := sample.Shares(r)
		sa := types.NewSpecialAllocations(genesisDistribution, claimableAirdrop)
		require.True(t, types.IsEqualShares(
			sa.TotalShares(),
			types.IncreaseShares(sa.GenesisDistribution, sa.ClaimableAirdrop),
		))
	})
}
