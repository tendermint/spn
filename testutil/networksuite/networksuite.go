// Package networksuite provides base test suite for tests that need a local network instance
package networksuite

import (
	"math/rand"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/testutil/sample"
	campaign "github.com/tendermint/spn/x/campaign/types"
	claim "github.com/tendermint/spn/x/claim/types"
	launch "github.com/tendermint/spn/x/launch/types"
	monitoringc "github.com/tendermint/spn/x/monitoringc/types"
	participation "github.com/tendermint/spn/x/participation/types"
	profile "github.com/tendermint/spn/x/profile/types"
	reward "github.com/tendermint/spn/x/reward/types"
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

	updateConfigGenesisState := func(moduleName string, moduleState proto.Message) {
		buf, err := cfg.Codec.MarshalJSON(moduleState)
		require.NoError(nts.T(), err)
		cfg.GenesisState[moduleName] = buf
	}

	// initialize launch
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[launch.ModuleName], &nts.LaunchState))
	nts.LaunchState = populateLaunch(r, nts.LaunchState)
	updateConfigGenesisState(launch.ModuleName, &nts.LaunchState)

	// initialize campaign
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[campaign.ModuleName], &nts.CampaignState))
	nts.CampaignState = populateCampaign(r, nts.CampaignState)
	updateConfigGenesisState(campaign.ModuleName, &nts.CampaignState)

	// initialize claim
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[claim.ModuleName], &nts.ClaimState))
	nts.ClaimState = populateClaim(r, nts.ClaimState)
	updateConfigGenesisState(claim.ModuleName, &nts.ClaimState)

	// initialize monitoring consumer
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[monitoringc.ModuleName], &nts.MonitoringcState))
	nts.MonitoringcState = populateMonitoringc(nts.MonitoringcState)
	updateConfigGenesisState(monitoringc.ModuleName, &nts.MonitoringcState)

	// initialize participation
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[participation.ModuleName], &nts.ParticipationState))
	nts.ParticipationState = populateParticipation(r, nts.ParticipationState)
	updateConfigGenesisState(participation.ModuleName, &nts.ParticipationState)

	// initialize profile
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[profile.ModuleName], &nts.ProfileState))
	nts.ProfileState = populateProfile(r, nts.ProfileState)
	updateConfigGenesisState(profile.ModuleName, &nts.ProfileState)

	// initialize reward
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[reward.ModuleName], &nts.RewardState))
	nts.RewardState = populateReward(nts.RewardState)
	updateConfigGenesisState(reward.ModuleName, &nts.RewardState)

	nts.Network = network.New(nts.T(), cfg)
}

func populateLaunch(r *rand.Rand, launchState launch.GenesisState) launch.GenesisState {
	// add chains
	for i := 0; i < 5; i++ {
		chain := sample.Chain(r, uint64(i), uint64(i))
		launchState.ChainList = append(
			launchState.ChainList,
			chain,
		)
	}

	// add genesis accounts
	for i := 0; i < 5; i++ {
		launchState.GenesisAccountList = append(
			launchState.GenesisAccountList,
			sample.GenesisAccount(r, 0, sample.Address(r)),
		)
	}

	// add vesting accounts
	for i := 0; i < 5; i++ {
		launchState.VestingAccountList = append(
			launchState.VestingAccountList,
			sample.VestingAccount(r, 0, sample.Address(r)),
		)
	}

	// add genesis validators
	for i := 0; i < 5; i++ {
		launchState.GenesisValidatorList = append(
			launchState.GenesisValidatorList,
			sample.GenesisValidator(r, uint64(0), sample.Address(r)),
		)
	}

	// add request
	for i := 0; i < 5; i++ {
		request := sample.Request(r, 0, sample.Address(r))
		request.RequestID = uint64(i)
		launchState.RequestList = append(
			launchState.RequestList,
			request,
		)
	}

	return launchState
}

func populateCampaign(r *rand.Rand, campaignState campaign.GenesisState) campaign.GenesisState {
	// add campaigns
	for i := 0; i < 5; i++ {
		campaignState.CampaignList = append(campaignState.CampaignList, sample.Campaign(r, uint64(i)))
	}

	// add campaign chains
	for i := 0; i < 5; i++ {
		campaignState.CampaignChainsList = append(campaignState.CampaignChainsList, campaign.CampaignChains{
			CampaignID: uint64(i),
			Chains:     []uint64{uint64(i)},
		})
	}

	// add mainnet accounts
	campaignID := uint64(5)
	for i := 0; i < 5; i++ {
		campaignState.MainnetAccountList = append(
			campaignState.MainnetAccountList,
			sample.MainnetAccount(r, campaignID, sample.Address(r)),
		)
	}

	return campaignState
}

func populateClaim(r *rand.Rand, claimState claim.GenesisState) claim.GenesisState {
	claimState.AirdropSupply = sample.Coin(r)

	// add claim records
	for i := 0; i < 5; i++ {
		claimRecord := claim.ClaimRecord{
			Address:   sample.Address(r),
			Claimable: sdk.NewInt(r.Int63()),
		}
		nullify.Fill(&claimRecord)
		claimState.ClaimRecords = append(claimState.ClaimRecords, claimRecord)
	}

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
		monitoringcState.LaunchIDFromChannelIDList = append(
			monitoringcState.LaunchIDFromChannelIDList,
			launchIDFromChannelID,
		)
	}

	// add monitoring history
	for i := 0; i < 5; i++ {
		monitoringHistory := monitoringc.MonitoringHistory{
			LaunchID: uint64(i),
		}
		nullify.Fill(&monitoringHistory)
		monitoringcState.MonitoringHistoryList = append(monitoringcState.MonitoringHistoryList, monitoringHistory)
	}

	// add provider client ID
	for i := 0; i < 5; i++ {
		providerClientID := monitoringc.ProviderClientID{
			LaunchID: uint64(i),
		}
		nullify.Fill(&providerClientID)
		monitoringcState.ProviderClientIDList = append(monitoringcState.ProviderClientIDList, providerClientID)
	}

	// add verified client IDs
	for i := 0; i < 5; i++ {
		verifiedClientID := monitoringc.VerifiedClientID{
			LaunchID: uint64(i),
		}
		nullify.Fill(&verifiedClientID)
		monitoringcState.VerifiedClientIDList = append(monitoringcState.VerifiedClientIDList, verifiedClientID)
	}

	return monitoringcState
}

func populateParticipation(r *rand.Rand, participationState participation.GenesisState) participation.GenesisState {
	// add used allocations
	for i := 0; i < 5; i++ {
		usedAllocations := participation.UsedAllocations{
			Address: sample.Address(r),
		}
		nullify.Fill(&usedAllocations)
		participationState.UsedAllocationsList = append(participationState.UsedAllocationsList, usedAllocations)
	}

	// add auction used allocations
	address := sample.Address(r)
	for i := 0; i < 5; i++ {
		auctionUsedAllocations := participation.AuctionUsedAllocations{
			Address:   address,
			AuctionID: uint64(i),
		}
		nullify.Fill(&auctionUsedAllocations)
		participationState.AuctionUsedAllocationsList = append(participationState.AuctionUsedAllocationsList, auctionUsedAllocations)
	}

	return participationState
}

func populateProfile(r *rand.Rand, profileState profile.GenesisState) profile.GenesisState {
	// add coordinators
	for i := 0; i < 5; i++ {
		profileState.CoordinatorList = append(
			profileState.CoordinatorList,
			profile.Coordinator{CoordinatorID: uint64(i)},
		)
	}

	// add coordinator by address
	for i := 0; i < 5; i++ {
		profileState.CoordinatorByAddressList = append(
			profileState.CoordinatorByAddressList,
			profile.CoordinatorByAddress{Address: sample.Address(r)},
		)
	}

	// add validator
	for i := 0; i < 5; i++ {
		profileState.ValidatorList = append(profileState.ValidatorList, profile.Validator{
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
		rewardState.RewardPoolList = append(rewardState.RewardPoolList, rewardPool)
	}

	return rewardState
}
