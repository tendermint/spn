package sample

import (
	"math/rand"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	project "github.com/tendermint/spn/x/project/types"
)

// Shares returns a sample shares
func Shares(r *rand.Rand) project.Shares {
	return project.NewSharesFromCoins(Coins(r))
}

// SpecialAllocations returns a sample special allocations
func SpecialAllocations(r *rand.Rand) project.SpecialAllocations {
	return project.NewSpecialAllocations(Shares(r), Shares(r))
}

// ShareVestingOptions returns a sample ShareVestingOptions
func ShareVestingOptions(r *rand.Rand) project.ShareVestingOptions {
	// use vesting shares as total shares
	vestingShares := Shares(r)
	return *project.NewShareDelayedVesting(vestingShares, vestingShares, Time(r))
}

// Voucher returns a sample voucher structure
func Voucher(r *rand.Rand, projectID uint64) sdk.Coin {
	denom := project.VoucherDenom(projectID, AlphaString(r, 5))
	return sdk.NewCoin(denom, sdkmath.NewInt(int64(r.Intn(10000)+1)))
}

// Vouchers returns a sample vouchers structure
func Vouchers(r *rand.Rand, projectID uint64) sdk.Coins {
	return sdk.NewCoins(Voucher(r, projectID), Voucher(r, projectID), Voucher(r, projectID))
}

// CustomShareVestingOptions returns a sample ShareVestingOptions with shares
func CustomShareVestingOptions(r *rand.Rand, shares project.Shares) project.ShareVestingOptions {
	return *project.NewShareDelayedVesting(shares, shares, Time(r))
}

// ProjectName returns a sample project name
func ProjectName(r *rand.Rand) string {
	return String(r, 20)
}

// Project returns a sample project
func Project(r *rand.Rand, id uint64) project.Project {
	genesisDistribution := Shares(r)
	claimableAirdrop := Shares(r)
	shares := project.IncreaseShares(genesisDistribution, claimableAirdrop)

	project := project.NewProject(
		id,
		ProjectName(r),
		Uint64(r),
		TotalSupply(r),
		Metadata(r, 20),
		Duration(r).Milliseconds(),
	)

	// set random shares for special allocations
	project.AllocatedShares = shares
	project.SpecialAllocations.GenesisDistribution = genesisDistribution
	project.SpecialAllocations.ClaimableAirdrop = claimableAirdrop

	return project
}

// MainnetAccount returns a sample MainnetAccount
func MainnetAccount(r *rand.Rand, projectID uint64, address string) project.MainnetAccount {
	return project.MainnetAccount{
		ProjectID: projectID,
		Address:   address,
		Shares:    Shares(r),
	}
}

// MsgCreateProject returns a sample MsgCreateProject
func MsgCreateProject(r *rand.Rand, coordAddr string) project.MsgCreateProject {
	return project.MsgCreateProject{
		Coordinator: coordAddr,
		ProjectName: ProjectName(r),
		TotalSupply: TotalSupply(r),
	}
}

// ProjectParams returns a sample of params for the project module
func ProjectParams(r *rand.Rand) project.Params {
	// no point in randomizing these values, using defaults
	minTotalSupply := project.DefaultMinTotalSupply
	maxTotalSupply := project.DefaultMaxTotalSupply
	maxMetadataLength := project.DefaultMaxMetadataLength

	// assign random small amount of staking denom
	projectCreationFee := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, r.Int63n(100)+1))

	return project.NewParams(minTotalSupply, maxTotalSupply, projectCreationFee, maxMetadataLength)
}

// ProjectGenesisState returns a sample genesis state for the project module
func ProjectGenesisState(r *rand.Rand) project.GenesisState {
	project1, project2 := Project(r, 0), Project(r, 1)

	return project.GenesisState{
		Projects: []project.Project{
			project1,
			project2,
		},
		ProjectCounter: 2,
		ProjectChains: []project.ProjectChains{
			{
				ProjectID: 0,
				Chains:    []uint64{0, 1},
			},
		},
		TotalShares: spntypes.TotalShareNumber,
		Params:      ProjectParams(r),
	}
}

// ProjectGenesisStateWithAccounts returns a sample genesis state for the project module that includes accounts
func ProjectGenesisStateWithAccounts(r *rand.Rand) project.GenesisState {
	genState := ProjectGenesisState(r)
	genState.MainnetAccounts = make([]project.MainnetAccount, 0)

	for i, c := range genState.Projects {
		for j := 0; j < 5; j++ {
			mainnetAccount := MainnetAccount(r, c.ProjectID, Address(r))
			genState.MainnetAccounts = append(genState.MainnetAccounts, mainnetAccount)

			// increase project allocated shares accordingly
			c.AllocatedShares = project.IncreaseShares(c.AllocatedShares, mainnetAccount.Shares)
		}
		genState.Projects[i] = c
	}

	return genState
}
