// Package networksuite provides base test suite for tests that need a local network instance
package networksuite

import (
	campaign "github.com/tendermint/spn/x/campaign/types"
	"math/rand"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/sample"
	launch "github.com/tendermint/spn/x/launch/types"
)

// NetworkTestSuite is a test suite for query tests that initializes a network instance
type NetworkTestSuite struct {
	suite.Suite
	Network       *network.Network
	LaunchState   launch.GenesisState
	CampaignState campaign.GenesisState
}

// SetupSuite setups the local network with a genesis state
func (nts *NetworkTestSuite) SetupSuite() {
	r := sample.Rand()
	cfg := network.DefaultConfig()

	// initialize launch
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[launch.ModuleName], &nts.LaunchState))
	nts.LaunchState = populateLaunch(r, nts.LaunchState)
	buf, err := cfg.Codec.MarshalJSON(&nts.LaunchState)
	require.NoError(nts.T(), err)
	cfg.GenesisState[launch.ModuleName] = buf

	// initialize campaign
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[campaign.ModuleName], &nts.CampaignState))
	nts.CampaignState = populateCampaign(r, nts.CampaignState)
	buf, err = cfg.Codec.MarshalJSON(&nts.CampaignState)
	require.NoError(nts.T(), err)
	cfg.GenesisState[campaign.ModuleName] = buf

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
