// Package networksuite provides base test suite for tests that need a local network instance
package networksuite

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/sample"
	launch "github.com/tendermint/spn/x/launch/types"
)

// NetworkTestSuite is a test suite for query tests that initializes a network instance
type NetworkTestSuite struct {
	suite.Suite
	Network     *network.Network
	LaunchState launch.GenesisState
}

// SetupSuite setups the local network with a genesis state
func (nts *NetworkTestSuite) SetupSuite() {
	r := sample.Rand()
	cfg := network.DefaultConfig()

	// initialize launch
	launchState := launch.GenesisState{}
	require.NoError(nts.T(), cfg.Codec.UnmarshalJSON(cfg.GenesisState[launch.ModuleName], &launchState))

	for i := 0; i < 5; i++ {
		chain := sample.Chain(r, uint64(i), uint64(i))
		launchState.ChainList = append(
			launchState.ChainList,
			chain,
		)
	}
	buf, err := cfg.Codec.MarshalJSON(&launchState)
	require.NoError(nts.T(), err)
	cfg.GenesisState[launch.ModuleName] = buf

	nts.Network = network.New(nts.T(), cfg)
	nts.LaunchState = launchState
}
