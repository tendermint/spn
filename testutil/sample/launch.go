package sample

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	launch "github.com/tendermint/spn/x/launch/types"
)

// Chain returns a sample Chain
func Chain(chainID string, coordinatorID uint64) *launch.Chain {
	defaultGenesis, err := types.NewAnyWithValue((*launch.DefaultInitialGenesis)(nil))
	if err != nil {
		panic(err)
	}

	// Byte array is nullified in the store if empty
	defaultGenesis.Value = []byte(nil)

	return &launch.Chain{
		ChainID:         chainID,
		CoordinatorID:   coordinatorID,
		CreatedAt:       time.Now().Unix(),
		SourceURL:       String(10),
		SourceHash:      String(10),
		LaunchTriggered: false,
		InitialGenesis:  defaultGenesis,
	}
}

// MsgCreateChain returns a sample MsgCreateChain
func MsgCreateChain(coordAddress, chainName, genesisURL string) launch.MsgCreateChain {
	var genesisHash string
	if len(genesisURL) > 0 {
		genesisHash = String(10)
	}

	return *launch.NewMsgCreateChain(
		coordAddress,
		chainName,
		String(10),
		String(10),
		genesisURL,
		genesisHash,
	)
}
