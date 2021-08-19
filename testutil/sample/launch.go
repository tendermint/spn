package sample

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/gogo/protobuf/proto"
	launch "github.com/tendermint/spn/x/launch/types"
)

// ChainID returns a sample chain id with the associated chain name
func ChainID(number uint64) (string, string) {
	chainName := AlphaString(5)
	return launch.ChainIDFromChainName(chainName, number), chainName
}

// Chain returns a sample Chain
func Chain(chainID string, coordinatorID uint64) *launch.Chain {
	return &launch.Chain{
		ChainID:         chainID,
		CoordinatorID:   coordinatorID,
		CreatedAt:       time.Now().Unix(),
		SourceURL:       String(10),
		SourceHash:      String(10),
		LaunchTriggered: false,
		InitialGenesis:  launch.NewDefaultInitialGenesis(),
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

// AccountRemoval returns a sample AccountRemoval
func AccountRemoval(address string) *launch.AccountRemoval {
	return &launch.AccountRemoval{
		Address: address,
	}
}

// ValidatorRemoval returns a sample ValidatorRemoval
func ValidatorRemoval(address string) *launch.ValidatorRemoval {
	return &launch.ValidatorRemoval{
		ValAddress: address,
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

// GenesisValidator returns a sample GenesisValidator
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

// RequestWithContent creates a launch request object with chain id and content
func RequestWithContent(chainID string, content *types.Any) *launch.Request {
	return &launch.Request{
		ChainID:   chainID,
		Creator:   AccAddress(),
		CreatedAt: time.Now().Unix(),
		Content:   content,
	}
}

// AllRequestContents creates all contents types for request and
// returns a list of all pack contents converted to `types.Any` object
func AllRequestContents(chainID, genesis, vested, validator string) []*types.Any {
	contents := make([]proto.Message, 0)
	contents = append(contents,
		GenesisAccount(chainID, genesis),
		AccountRemoval(genesis),
		VestedAccount(chainID, vested),
		AccountRemoval(vested),
		GenesisValidator(chainID, validator),
		ValidatorRemoval(validator),
	)

	result := make([]*types.Any, 0)
	for _, content := range contents {
		msg, err := types.NewAnyWithValue(content)
		if err != nil {
			panic(err)
		}
		result = append(result, msg)
	}
	return result
}

// GenesisAccountContent returns a sample GenesisAccount request content packed into an *Any object
func GenesisAccountContent(chainID, address string) *types.Any {
	content, err := types.NewAnyWithValue(GenesisAccount(chainID, address))
	if err != nil {
		panic(err)
	}
	return content
}

// Request returns a sample Request
func Request(chainID string) *launch.Request {
	content := GenesisAccountContent(chainID, AccAddress())
	return RequestWithContent(chainID, content)
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
	var initialGenesis *launch.InitialGenesis
	if modifyInitialGenesis {
		if genesisURL {
			initialGenesis = launch.NewGenesisURL(String(30), GenesisHash())
		} else {
			initialGenesis = launch.NewDefaultInitialGenesis()
		}
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
