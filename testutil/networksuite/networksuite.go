// Package networksuite provides base test suite for tests that need a local network instance
package networksuite

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	claim "github.com/ignite/modules/x/claim/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	campaign "github.com/tendermint/spn/x/campaign/types"
	launch "github.com/tendermint/spn/x/launch/types"
	monitoringc "github.com/tendermint/spn/x/monitoringc/types"
	participation "github.com/tendermint/spn/x/participation/types"
	profile "github.com/tendermint/spn/x/profile/types"
	reward "github.com/tendermint/spn/x/reward/types"
	"math/rand"
	"strconv"
)

// NetworkTestSuite is a test suite for query tests that initializes a network instance
type NetworkTestSuite struct {
	suite.Suite
	Network            *network.Network
	LaunchState        launch.GenesisState
	CampaignState      campaign.GenesisState
	ClaimState         claim.GenesisState
	MonitoringcState   monitoringc.GenesisState
	ParticipationState participation.GenesisState
	ProfileState       profile.GenesisState
	RewardState        reward.GenesisState
}

// SetupSuite setups the local network with a genesis state
func (nts *NetworkTestSuite) SetupSuite() {
	r := sample.Rand()
	cfg := network.DefaultConfig()

	updateGenesisConfigState := func(moduleName string, moduleState proto.Message) {
		buf, err := cfg.Codec.MarshalJSON(moduleState)
		require.NoError(nts.T(), err)
		cfg.GenesisState[moduleName] = buf
	}

	// initialize launch
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[launch.ModuleName], &nts.LaunchState))
	nts.LaunchState = populateLaunch(r, nts.LaunchState)
	updateGenesisConfigState(launch.ModuleName, &nts.LaunchState)

	// initialize campaign
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[campaign.ModuleName], &nts.CampaignState))
	nts.CampaignState = populateCampaign(r, nts.CampaignState)
	updateGenesisConfigState(campaign.ModuleName, &nts.CampaignState)

	// initialize claim
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[claim.ModuleName], &nts.ClaimState))
	nts.ClaimState = populateClaim(r, nts.ClaimState)
	updateGenesisConfigState(claim.ModuleName, &nts.ClaimState)

	// initialize monitoring consumer
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[monitoringc.ModuleName], &nts.MonitoringcState))
	nts.MonitoringcState = populateMonitoringc(nts.MonitoringcState)
	updateGenesisConfigState(monitoringc.ModuleName, &nts.MonitoringcState)

	// initialize participation
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[participation.ModuleName], &nts.ParticipationState))
	nts.ParticipationState = populateParticipation(r, nts.ParticipationState)
	updateGenesisConfigState(participation.ModuleName, &nts.ParticipationState)

	// initialize profile
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[profile.ModuleName], &nts.ProfileState))
	nts.ProfileState = populateProfile(r, nts.ProfileState)
	updateGenesisConfigState(profile.ModuleName, &nts.ProfileState)

	// initialize reward
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[reward.ModuleName], &nts.RewardState))
	nts.RewardState = populateReward(nts.RewardState)
	updateGenesisConfigState(reward.ModuleName, &nts.RewardState)

	nts.Network = network.New(nts.T(), cfg)
}

func populateLaunch(r *rand.Rand, launchState launch.GenesisState) launch.GenesisState {
	// add chains
	for i := 0; i < 5; i++ {
		chain := sample.Chain(r, uint64(i), uint64(i))
		launchState.Chains = append(
			launchState.Chains,
			chain,
		)
	}

	// add genesis accounts
	for i := 0; i < 5; i++ {
		launchState.GenesisAccounts = append(
			launchState.GenesisAccounts,
			sample.GenesisAccount(r, 0, sample.Address(r)),
		)
	}

	// add vesting accounts
	for i := 0; i < 5; i++ {
		launchState.VestingAccounts = append(
			launchState.VestingAccounts,
			sample.VestingAccount(r, 0, sample.Address(r)),
		)
	}

	// add genesis validators
	for i := 0; i < 5; i++ {
		launchState.GenesisValidators = append(
			launchState.GenesisValidators,
			sample.GenesisValidator(r, uint64(0), sample.Address(r)),
		)
	}

	// add param chagne
	for i := 0; i < 5; i++ {
		launchState.ParamChanges = append(
			launchState.ParamChanges,
			sample.ParamChange(r, uint64(0)),
		)
	}

	// add request
	for i := 0; i < 5; i++ {
		request := sample.Request(r, 0, sample.Address(r))
		request.RequestID = uint64(i)
		launchState.Requests = append(
			launchState.Requests,
			request,
		)
	}

	return launchState
}

func populateCampaign(r *rand.Rand, campaignState campaign.GenesisState) campaign.GenesisState {
	// add campaigns
	for i := 0; i < 5; i++ {
		camp := campaign.Campaign{
			CampaignID: uint64(i),
		}
		nullify.Fill(&camp)
		campaignState.Campaigns = append(campaignState.Campaigns, camp)
	}

	// add campaign chains
	for i := 0; i < 5; i++ {
		campaignState.CampaignChains = append(campaignState.CampaignChains, campaign.CampaignChains{
			CampaignID: uint64(i),
			Chains:     []uint64{uint64(i)},
		})
	}

	// add mainnet accounts
	campaignID := uint64(5)
	for i := 0; i < 5; i++ {
		campaignState.MainnetAccounts = append(
			campaignState.MainnetAccounts,
			sample.MainnetAccount(r, campaignID, sample.Address(r)),
		)
	}

	return campaignState
}

func populateClaim(r *rand.Rand, claimState claim.GenesisState) claim.GenesisState {
	claimState.AirdropSupply = sample.Coin(r)
	totalSupply := sdkmath.ZeroInt()
	for i := 0; i < 5; i++ {
		// fill claim records
		accSupply := sdkmath.NewIntFromUint64(r.Uint64() % 1000)
		claimRecord := claim.ClaimRecord{
			Claimable: accSupply,
			Address:   sample.Address(r),
		}
		totalSupply = totalSupply.Add(accSupply)
		nullify.Fill(&claimRecord)
		claimState.ClaimRecords = append(claimState.ClaimRecords, claimRecord)
	}
	claimState.AirdropSupply.Amount = totalSupply

	// add missions
	for i := 0; i < 5; i++ {
		mission := claim.Mission{
			MissionID: uint64(i),
			Weight:    sdk.NewDec(r.Int63()),
		}
		nullify.Fill(&mission)
		claimState.Missions = append(claimState.Missions, mission)
	}

	return claimState
}

func populateMonitoringc(monitoringcState monitoringc.GenesisState) monitoringc.GenesisState {
	// add launch ID from channel ID
	for i := 0; i < 5; i++ {
		launchIDFromChannelID := monitoringc.LaunchIDFromChannelID{
			ChannelID: strconv.Itoa(i),
		}
		nullify.Fill(&launchIDFromChannelID)
		monitoringcState.LaunchIDsFromChannelID = append(
			monitoringcState.LaunchIDsFromChannelID,
			launchIDFromChannelID,
		)
	}

	// add monitoring history
	for i := 0; i < 5; i++ {
		monitoringHistory := monitoringc.MonitoringHistory{
			LaunchID: uint64(i),
		}
		nullify.Fill(&monitoringHistory)
		monitoringcState.MonitoringHistories = append(monitoringcState.MonitoringHistories, monitoringHistory)
	}

	// add provider client ID
	for i := 0; i < 5; i++ {
		providerClientID := monitoringc.ProviderClientID{
			LaunchID: uint64(i),
		}
		nullify.Fill(&providerClientID)
		monitoringcState.ProviderClientIDs = append(monitoringcState.ProviderClientIDs, providerClientID)
	}

	// add verified client IDs
	for i := 0; i < 5; i++ {
		verifiedClientID := monitoringc.VerifiedClientID{
			LaunchID: uint64(i),
		}
		nullify.Fill(&verifiedClientID)
		monitoringcState.VerifiedClientIDs = append(monitoringcState.VerifiedClientIDs, verifiedClientID)
	}

	return monitoringcState
}

func populateParticipation(r *rand.Rand, participationState participation.GenesisState) participation.GenesisState {
	// add used allocations
	for i := 0; i < 5; i++ {
		usedAllocations := participation.UsedAllocations{
			Address:        sample.Address(r),
			NumAllocations: sample.Int(r),
		}
		nullify.Fill(&usedAllocations)
		participationState.UsedAllocationsList = append(participationState.UsedAllocationsList, usedAllocations)
	}

	// add auction used allocations
	address := sample.Address(r)
	for i := 0; i < 5; i++ {
		auctionUsedAllocations := participation.AuctionUsedAllocations{
			Address:        address,
			AuctionID:      uint64(i),
			NumAllocations: sample.Int(r),
		}
		nullify.Fill(&auctionUsedAllocations)
		participationState.AuctionUsedAllocationsList = append(participationState.AuctionUsedAllocationsList, auctionUsedAllocations)
	}

	return participationState
}

func populateProfile(r *rand.Rand, profileState profile.GenesisState) profile.GenesisState {
	// add coordinators
	for i := 0; i < 5; i++ {
		profileState.Coordinators = append(
			profileState.Coordinators,
			profile.Coordinator{CoordinatorID: uint64(i)},
		)
	}

	// add coordinator by address
	for i := 0; i < 5; i++ {
		profileState.CoordinatorsByAddress = append(
			profileState.CoordinatorsByAddress,
			profile.CoordinatorByAddress{Address: sample.Address(r)},
		)
	}

	// add validator
	for i := 0; i < 5; i++ {
		profileState.Validators = append(profileState.Validators, profile.Validator{
			Address: sample.Address(r),
		})
	}

	return profileState
}

func populateReward(rewardState reward.GenesisState) reward.GenesisState {
	// add reward pool
	for i := 0; i < 5; i++ {
		rewardPool := reward.RewardPool{
			LaunchID: uint64(i),
		}
		nullify.Fill(&rewardPool)
		rewardState.RewardPools = append(rewardState.RewardPools, rewardPool)
	}

	return rewardState
}
