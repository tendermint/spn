package sample

import (
	"crypto/sha256"
	"encoding/hex"
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

func GenesisValidator(chainID, address string) *launch.GenesisValidator {
	return &launch.GenesisValidator{
		ChainID:        chainID,
		Address:        address,
		GenTx:          Bytes(200),
		ConsPubKey:     Bytes(10),
		SelfDelegation: Coin(),
		Peer:           String(10),
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

// MsgCreateChain returns a sample MsgCreateChain
func MsgCreateChain(coordAddress, chainName, genesisURL string) launch.MsgCreateChain {
	var genesisHash string
	if len(genesisURL) > 0 {
		genesisHash = GenesisHash()
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

// MsgEditChain returns a sample MsgEditChain
func MsgEditChain(
	coordAddress,
	chainID string,
	modifySource,
	modifyInitialGenesis,
	genesisURL bool,
) launch.MsgEditChain {
	var sourceURL, sourceHash string
	if modifySource {
		sourceURL, sourceHash = String(30), String(10)
	}
	var initialGenesis *types.Any
	if modifyInitialGenesis {
		if genesisURL {
			initialGenesis, _ = types.NewAnyWithValue(&launch.GenesisURL{
				Url:  String(30),
				Hash: GenesisHash(),
			})
		} else {
			initialGenesis, _ = types.NewAnyWithValue(&launch.DefaultInitialGenesis{})
			initialGenesis.Value = nil
		}
		initialGenesis.ClearCachedValue()
	}

	return *launch.NewMsgEditChain(
		coordAddress,
		chainID,
		sourceURL,
		sourceHash,
		initialGenesis,
	)
}

// MsgRequestAddValidator returns a sample MsgRequestAddValidator
func MsgRequestAddValidator(address, chainID string) launch.MsgRequestAddValidator {
	return *launch.NewMsgRequestAddValidator(
		address,
		chainID,
		Bytes(500),
		Bytes(30),
		Coin(),
		String(30),
	)
}

// GenesisHash returns a sample sha256 hash of custom genesis for GenesisURL
func GenesisHash() string {
	hash := sha256.Sum256([]byte(String(50)))
	return hex.EncodeToString(hash[:])
}
