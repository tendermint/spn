package sample

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	campaign "github.com/tendermint/spn/x/campaign/types"
)

// Shares returns a sample shares
func Shares(r *rand.Rand) campaign.Shares {
	return campaign.NewSharesFromCoins(Coins(r))
}

// SpecialAllocations returns a sample special allocations
func SpecialAllocations(r *rand.Rand) campaign.SpecialAllocations {
	return campaign.NewSpecialAllocations(Shares(r), Shares(r))
}

// ShareVestingOptions returns a sample ShareVestingOptions
func ShareVestingOptions(r *rand.Rand) campaign.ShareVestingOptions {
	// use vesting shares as total shares
	vestingShares := Shares(r)
	return *campaign.NewShareDelayedVesting(vestingShares, vestingShares, Duration(r).Microseconds())
}

// Voucher returns a sample voucher structure
func Voucher(r *rand.Rand, campaignID uint64) sdk.Coin {
	denom := campaign.VoucherDenom(campaignID, AlphaString(r, 5))
	return sdk.NewCoin(denom, sdk.NewInt(int64(r.Intn(10000)+1)))
}

// Vouchers returns a sample vouchers structure
func Vouchers(r *rand.Rand, campaignID uint64) sdk.Coins {
	return sdk.NewCoins(Voucher(r, campaignID), Voucher(r, campaignID), Voucher(r, campaignID))
}

// CustomShareVestingOptions returns a sample ShareVestingOptions with shares
func CustomShareVestingOptions(r *rand.Rand, shares campaign.Shares) campaign.ShareVestingOptions {
	return *campaign.NewShareDelayedVesting(shares, shares, Duration(r).Microseconds())
}

// MainnetVestingAccount returns a sample MainnetVestingAccount
func MainnetVestingAccount(r *rand.Rand, campaignID uint64, address string) campaign.MainnetVestingAccount {
	return MainnetVestingAccountWithShares(r, campaignID, address, Shares(r))
}

// MainnetVestingAccountWithShares returns a sample MainnetVestingAccount with custom shares
func MainnetVestingAccountWithShares(
	r *rand.Rand,
	campaignID uint64,
	address string,
	shares campaign.Shares,
) campaign.MainnetVestingAccount {
	return campaign.MainnetVestingAccount{
		CampaignID:     campaignID,
		Address:        address,
		VestingOptions: CustomShareVestingOptions(r, shares),
	}
}

// CampaignName returns a sample campaign name
func CampaignName(r *rand.Rand) string {
	return String(r, 20)
}

// Campaign returns a sample campaign
func Campaign(r *rand.Rand, id uint64) campaign.Campaign {
	genesisDistribution := Shares(r)
	claimableAirdrop := Shares(r)
	shares := campaign.IncreaseShares(genesisDistribution, claimableAirdrop)

	campaign := campaign.NewCampaign(
		id,
		CampaignName(r),
		Uint64(r),
		TotalSupply(r),
		Metadata(r, 20),
		Duration(r).Milliseconds(),
	)

	// set random shares for special allocations
	campaign.AllocatedShares = shares
	campaign.SpecialAllocations.GenesisDistribution = genesisDistribution
	campaign.SpecialAllocations.ClaimableAirdrop = claimableAirdrop

	return campaign
}

// MainnetAccount returns a sample MainnetAccount
func MainnetAccount(r *rand.Rand, campaignID uint64, address string) campaign.MainnetAccount {
	return campaign.MainnetAccount{
		CampaignID: campaignID,
		Address:    address,
		Shares:     Shares(r),
	}
}

// MsgCreateCampaign returns a sample MsgCreateCampaign
func MsgCreateCampaign(r *rand.Rand, coordAddr string) campaign.MsgCreateCampaign {
	return campaign.MsgCreateCampaign{
		Coordinator:  coordAddr,
		CampaignName: CampaignName(r),
		TotalSupply:  TotalSupply(r),
	}
}

// CampaignParams returns a sample of params for the campaign module
func CampaignParams(r *rand.Rand) campaign.Params {
	// no point in randomizing these values, using defaults
	minTotalSupply := campaign.DefaultMinTotalSupply
	maxTotalSupply := campaign.DefaultMaxTotalSupply

	// assign random small amount of staking denom
	campaignCreationFee := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, r.Int63n(100)+1))

	return campaign.NewParams(minTotalSupply, maxTotalSupply, campaignCreationFee)
}

// CampaignGenesisState returns a sample genesis state for the campaign module
func CampaignGenesisState(r *rand.Rand) campaign.GenesisState {
	campaign1, campaign2 := Campaign(r, 0), Campaign(r, 1)

	return campaign.GenesisState{
		CampaignList: []campaign.Campaign{
			campaign1,
			campaign2,
		},
		CampaignCounter: 2,
		CampaignChainsList: []campaign.CampaignChains{
			{
				CampaignID: 0,
				Chains:     []uint64{0, 1},
			},
		},
		TotalShares: spntypes.TotalShareNumber,
		Params:      CampaignParams(r),
	}
}

// CampaignGenesisStateWithAccounts returns a sample genesis state for the campaign module that includes accounts
func CampaignGenesisStateWithAccounts(r *rand.Rand) campaign.GenesisState {
	genState := CampaignGenesisState(r)
	genState.MainnetAccountList = make([]campaign.MainnetAccount, 0)
	genState.MainnetVestingAccountList = make([]campaign.MainnetVestingAccount, 0)

	for i, c := range genState.CampaignList {
		for j := 0; j < 5; j++ {
			mainnetAccount := MainnetAccount(r, c.CampaignID, Address(r))
			mainnetVestingAccount := MainnetVestingAccount(r, c.CampaignID, Address(r))
			genState.MainnetAccountList = append(genState.MainnetAccountList, mainnetAccount)
			genState.MainnetVestingAccountList = append(genState.MainnetVestingAccountList, mainnetVestingAccount)

			// increase campaign allocated shares accordingly
			c.AllocatedShares = campaign.IncreaseShares(c.AllocatedShares, mainnetAccount.Shares)
			shares, _ := mainnetVestingAccount.GetTotalShares()
			c.AllocatedShares = campaign.IncreaseShares(c.AllocatedShares, shares)
		}
		genState.CampaignList[i] = c
	}

	return genState
}
