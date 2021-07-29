package sample

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	launch "github.com/tendermint/spn/x/launch/types"
)

// ChainID returns a sample chain id with the associated chain name
func ChainID(number uint64) (string, string) {
	chainName := AlphaString(5)
	return launch.ChainIDFromChainName(chainName, number), chainName
}

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

// GenesisAccount returns a sample GenesisAccount
func GenesisAccount(chainID, address string) *launch.GenesisAccount {
	return &launch.GenesisAccount{
		ChainID: chainID,
		Address: address,
		Coins:   Coins(),
	}
}

// VestedAccount returns a sample VestedAccount
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

func Request(chainID string) *launch.Request {
	content, err := types.NewAnyWithValue(GenesisAccount(chainID, AccAddress()))
	if err != nil {
		panic(err)
	}

	return &launch.Request{
		ChainID:   chainID,
		Creator:   AccAddress(),
		CreatedAt: time.Now().Unix(),
		Content:   content,
	}
}
