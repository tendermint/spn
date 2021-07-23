package sample

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	launch "github.com/tendermint/spn/x/launch/types"
	"time"
)

func Chain(chainID string, coordinatorID uint64) *launch.Chain {
	defaultGenesis, err := types.NewAnyWithValue(&launch.DefaultInitialGenesis{})
	if err != nil {
		panic(err)
	}

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
