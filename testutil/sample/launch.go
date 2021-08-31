package sample

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	launch "github.com/tendermint/spn/x/launch/types"
)

// GenesisChainID returns a sample chain id
func GenesisChainID() string {
	chainName := AlphaString(5)
	number := Uint64()
	return launch.NewGenesisChainID(chainName, number)
}

// Chain returns a sample Chain
func Chain(id uint64, coordinatorID uint64) launch.Chain {
	return launch.Chain{
		Id:              id,
		CoordinatorID:   coordinatorID,
		GenesisChainID:  GenesisChainID(),
		CreatedAt:       time.Now().Unix(),
		SourceURL:       String(10),
		SourceHash:      String(10),
		LaunchTriggered: false,
		InitialGenesis:  launch.NewDefaultInitialGenesis(),
	}
}

// GenesisAccount returns a sample GenesisAccount
func GenesisAccount(chainID uint64, address string) launch.GenesisAccount {
	return launch.GenesisAccount{
		ChainID: chainID,
		Address: address,
		Coins:   Coins(),
	}
}

// VestingOptions returns a sample VestingOptions
func VestingOptions() launch.VestingOptions {
	return *launch.NewDelayedVesting(Coins(), time.Now().Unix())
}

// VestingAccount returns a sample VestingAccount
func VestingAccount(chainID uint64, address string) launch.VestingAccount {
	return launch.VestingAccount{
		ChainID:         chainID,
		Address:         address,
		StartingBalance: Coins(),
		VestingOptions:  *launch.NewDelayedVesting(Coins(), time.Now().Unix()),
	}
}

// AccountRemoval returns a sample AccountRemoval
func AccountRemoval(address string) *launch.AccountRemoval {
	return &launch.AccountRemoval{
		Address: address,
	}
}

// GenesisValidator returns a sample GenesisValidator
func GenesisValidator(chainID uint64, address string) launch.GenesisValidator {
	return launch.GenesisValidator{
		ChainID:        chainID,
		Address:        address,
		GenTx:          Bytes(200),
		ConsPubKey:     Bytes(10),
		SelfDelegation: Coin(),
		Peer:           String(10),
	}
}

// ValidatorRemoval returns a sample ValidatorRemoval
func ValidatorRemoval(address string) launch.ValidatorRemoval {
	return launch.ValidatorRemoval{
		ValAddress: address,
	}
}

// RequestWithContent creates a launch request object with chain id and content
func RequestWithContent(chainID uint64, content launch.RequestContent) launch.Request {
	return launch.Request{
		ChainID:   chainID,
		Creator:   AccAddress(),
		CreatedAt: time.Now().Unix(),
		Content:   content,
	}
}

// AllRequestContents creates all contents types for request
func AllRequestContents(chainID uint64, genesis, vesting, validator string) []launch.RequestContent {
	return []launch.RequestContent{
		launch.NewGenesisAccount(chainID, genesis, Coins()),
		launch.NewAccountRemoval(genesis),
		launch.NewVestingAccount(chainID, vesting, Coins(), VestingOptions()),
		launch.NewAccountRemoval(vesting),
		launch.NewGenesisValidator(chainID, validator, Bytes(300), Bytes(30), Coin(), String(30)),
		launch.NewValidatorRemoval(validator),
	}
}

// GenesisAccountContent returns a sample GenesisAccount request content
func GenesisAccountContent(chainID uint64, address string) launch.RequestContent {
	return launch.NewGenesisAccount(chainID, address, Coins())
}

// Request returns a sample Request
func Request(chainID uint64) launch.Request {
	content := GenesisAccountContent(chainID, AccAddress())
	return RequestWithContent(chainID, content)
}

// MsgCreateChain returns a sample MsgCreateChain
func MsgCreateChain(coordAddress, genesisURL string) launch.MsgCreateChain {
	var genesisHash string
	if len(genesisURL) > 0 {
		genesisHash = GenesisHash()
	}

	return *launch.NewMsgCreateChain(
		coordAddress,
		GenesisChainID(),
		String(10),
		String(10),
		genesisURL,
		genesisHash,
	)
}

// MsgEditChain returns a sample MsgEditChain
func MsgEditChain(
	coordAddress string,
	chainID uint64,
	modifyGenesisChainID,
	modifySource,
	modifyInitialGenesis,
	genesisURL bool,
) launch.MsgEditChain {
	var genesisChainID, sourceURL, sourceHash string

	if modifyGenesisChainID {
		genesisChainID = GenesisChainID()
	}
	if modifySource {
		sourceURL, sourceHash = String(30), String(10)
	}
	var initialGenesis *launch.InitialGenesis
	if modifyInitialGenesis {
		if genesisURL {
			newGenesisURL := launch.NewGenesisURL(String(30), GenesisHash())
			initialGenesis = &newGenesisURL
		} else {
			newDefault := launch.NewDefaultInitialGenesis()
			initialGenesis = &newDefault
		}
	}

	return *launch.NewMsgEditChain(
		coordAddress,
		chainID,
		genesisChainID,
		sourceURL,
		sourceHash,
		initialGenesis,
	)
}

// MsgRequestAddValidator returns a sample MsgRequestAddValidator
func MsgRequestAddValidator(address string, chainID uint64) launch.MsgRequestAddValidator {
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

// LaunchParams returns a sample of params for the launch module
func LaunchParams() launch.Params {
	maxLaunchTime := rand.Intn(int(launch.MaxParametrableLaunchTime))
	minLaunchTime := rand.Intn(maxLaunchTime)

	return launch.Params{
		MinLaunchTime: uint64(minLaunchTime),
		MaxLaunchTime: uint64(maxLaunchTime),
	}
}

// LaunchGenesisState returns a sample genesis state for the launch module
func LaunchGenesisState() launch.GenesisState {
	chainID1 := uint64(0)
	chainID2 := uint64(1)

	return launch.GenesisState{
		ChainList: []launch.Chain{
			Chain(chainID1, Uint64()),
			Chain(chainID2, Uint64()),
		},
		ChainCount: 10,
		GenesisAccountList: []launch.GenesisAccount{
			GenesisAccount(chainID1, AccAddress()),
			GenesisAccount(chainID1, AccAddress()),
			GenesisAccount(chainID2, AccAddress()),
		},
		VestingAccountList: []launch.VestingAccount{
			VestingAccount(chainID1, AccAddress()),
			VestingAccount(chainID1, AccAddress()),
			VestingAccount(chainID2, AccAddress()),
		},
		GenesisValidatorList: []launch.GenesisValidator{
			GenesisValidator(chainID1, AccAddress()),
			GenesisValidator(chainID1, AccAddress()),
			GenesisValidator(chainID2, AccAddress()),
		},
		RequestList: []launch.Request{
			Request(chainID1),
			Request(chainID2),
		},
		RequestCountList: []launch.RequestCount{
			{
				ChainID: chainID1,
				Count:   1,
			},
			{
				ChainID: chainID2,
				Count:   1,
			},
		},
		Params: LaunchParams(),
	}
}
