package sample

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	launch "github.com/tendermint/spn/x/launch/types"
)

func GenesisAccount(chainID, address string) *launch.GenesisAccount {
	return &launch.GenesisAccount{
		ChainID: chainID,
		Address: address,
		Coins:   Coins(),
	}
}

func VestedAccount(chainID, address string) *launch.VestedAccount {
	delayedVesting, err := types.NewAnyWithValue(&launch.DelayedVesting{
		Vesting: Coins(),
		EndTime: time.Now().Unix(),
	})
	if err != nil {
		panic(err)
	}
	return &launch.VestedAccount{
		ChainID:         chainID,
		Address:         address,
		StartingBalance: Coins(),
		VestingOptions:  delayedVesting,
	}
}

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
