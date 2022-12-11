package types_test

import (
	"testing"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	project "github.com/tendermint/spn/x/project/types"
)

var (
	invalidProjectName  = "not_valid"
	invalidProjectCoins = sdk.Coins{sdk.Coin{Denom: "invalid denom", Amount: sdkmath.ZeroInt()}}
)

func TestNewProject(t *testing.T) {
	projectID := sample.Uint64(r)
	projectName := sample.ProjectName(r)
	coordinator := sample.Uint64(r)
	totalSupply := sample.TotalSupply(r)
	metadata := sample.Metadata(r, 20)
	createdAt := sample.Duration(r).Milliseconds()

	t.Run("should allow creation of project", func(t *testing.T) {
		c := project.NewProject(
			projectID,
			projectName,
			coordinator,
			totalSupply,
			metadata,
			createdAt,
		)
		require.EqualValues(t, projectID, c.ProjectID)
		require.EqualValues(t, projectName, c.ProjectName)
		require.EqualValues(t, coordinator, c.CoordinatorID)
		require.EqualValues(t, createdAt, c.CreatedAt)
		require.False(t, c.MainnetInitialized)
		require.True(t, totalSupply.IsEqual(c.TotalSupply))
		require.EqualValues(t, project.EmptyShares(), c.AllocatedShares)
	})
}

func TestProject_Validate(t *testing.T) {
	var (
		invalidAllocatedShares           project.Project
		totalSharesReached               project.Project
		projectInvalidSpecialAllocations project.Project
	)

	t.Run("should verify that invalid coins is invalid", func(t *testing.T) {
		require.False(t, invalidProjectCoins.IsValid())
	})

	t.Run("should allow creation of valid allocations with totalshares reached", func(t *testing.T) {
		invalidAllocatedShares = sample.Project(r, 0)
		invalidAllocatedShares.AllocatedShares = project.NewSharesFromCoins(invalidProjectCoins)
		totalSharesReached = sample.Project(r, 0)
		totalSharesReached.AllocatedShares = project.NewSharesFromCoins(sdk.NewCoins(
			sdk.NewCoin("foo", sdkmath.NewInt(spntypes.TotalShareNumber+1)),
		))
		reached, err := project.IsTotalSharesReached(totalSharesReached.AllocatedShares, spntypes.TotalShareNumber)
		require.NoError(t, err)
		require.True(t, reached)
	})

	t.Run("should allow creation of project with invalid special allocations", func(t *testing.T) {
		invalidSpecialAllocations := project.NewSpecialAllocations(
			sample.Shares(r),
			project.Shares(sdk.NewCoins(
				sdk.NewCoin("foo", sdkmath.NewInt(100)),
				sdk.NewCoin("s/bar", sdkmath.NewInt(200)),
			)),
		)
		require.Error(t, invalidSpecialAllocations.Validate())
		projectInvalidSpecialAllocations = sample.Project(r, 0)
		projectInvalidSpecialAllocations.SpecialAllocations = invalidSpecialAllocations
	})

	for _, tc := range []struct {
		desc    string
		project project.Project
		valid   bool
	}{
		{
			desc:    "should allow validation of valid project",
			project: sample.Project(r, 0),
			valid:   true,
		},
		{
			desc: "invalid project name",
			project: project.NewProject(
				0,
				invalidProjectName,
				sample.Uint64(r),
				sample.TotalSupply(r),
				sample.Metadata(r, 20),
				sample.Duration(r).Milliseconds(),
			),
			valid: false,
		},
		{
			desc: "should prevent validation of project with invalid total supply",
			project: project.NewProject(
				0,
				sample.ProjectName(r),
				sample.Uint64(r),
				invalidProjectCoins,
				sample.Metadata(r, 20),
				sample.Duration(r).Milliseconds(),
			),
			valid: false,
		},
		{
			desc:    "should prevent validation of project with invalid allocated shares",
			project: invalidAllocatedShares,
			valid:   false,
		},
		{
			desc:    "should prevent validation of project with allocated shares greater than total shares",
			project: totalSharesReached,
			valid:   false,
		},
		{
			desc:    "should prevent validation of project with invalid special allocations",
			project: projectInvalidSpecialAllocations,
			valid:   false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, tc.project.Validate(spntypes.TotalShareNumber) == nil)
		})
	}
}

func TestCheckProjectName(t *testing.T) {
	for _, tc := range []struct {
		desc  string
		name  string
		valid bool
	}{
		{
			desc:  "should allow check of project with valid name",
			name:  "ThisIs-a-ValidProjectName123",
			valid: true,
		},
		{
			desc:  "should prevent check of project with special character outside hyphen",
			name:  invalidProjectName,
			valid: false,
		},
		{
			desc:  "should prevent check of project with empty name",
			name:  "",
			valid: false,
		},
		{
			desc:  "should prevent check of project with name exceeding max length",
			name:  sample.String(r, project.ProjectNameMaxLength+1),
			valid: false,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			require.EqualValues(t, tc.valid, project.CheckProjectName(tc.name) == nil)
		})
	}
}
