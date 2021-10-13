package sample

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"

	"github.com/tendermint/spn/pkg/chainid"
	launch "github.com/tendermint/spn/x/launch/types"
	profile "github.com/tendermint/spn/x/profile/types"
)

// GenesisChainID returns a sample chain id
func GenesisChainID() string {
	chainName := AlphaString(5)
	number := Uint64()
	return chainid.NewGenesisChainID(chainName, number)
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
		VestingOptions:  VestingOptions(),
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
		Creator:   Address(),
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
	content := GenesisAccountContent(chainID, Address())
	return RequestWithContent(chainID, content)
}

// MsgCreateChain returns a sample MsgCreateChain
func MsgCreateChain(coordAddress, genesisURL string, hasCampaign bool, campaignID uint64) launch.MsgCreateChain {
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
		hasCampaign,
		campaignID,
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

// MsgRequestAddAccount returns a sample MsgRequestAddAccount
func MsgRequestAddAccount(address string, chainID uint64) launch.MsgRequestAddAccount {
	return *launch.NewMsgRequestAddAccount(
		address,
		chainID,
		Coins(),
	)
}

// MsgRequestAddVestingAccount returns a sample MsgRequestAddVestingAccount
func MsgRequestAddVestingAccount(address string, chainID uint64) launch.MsgRequestAddVestingAccount {
	return *launch.NewMsgRequestAddVestingAccount(
		address,
		chainID,
		Coins(),
		VestingOptions(),
	)
}

// MsgRequestRemoveAccount returns a sample MsgRequestRemoveAccount
func MsgRequestRemoveAccount(creator, address string, chainID uint64) launch.MsgRequestRemoveAccount {
	return *launch.NewMsgRequestRemoveAccount(
		chainID,
		creator,
		address,
	)
}

// MsgRequestRemoveValidator returns a sample MsgRequestRemoveValidator
func MsgRequestRemoveValidator(creator, validatorAddr string, chainID uint64) launch.MsgRequestRemoveValidator {
	return *launch.NewMsgRequestRemoveValidator(
		chainID,
		creator,
		validatorAddr,
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

// MsgRevertLaunch returns a sample MsgRevertLaunch
func MsgRevertLaunch(coordinator string, chainID uint64) launch.MsgRevertLaunch {
	return *launch.NewMsgRevertLaunch(
		coordinator,
		chainID,
	)
}

// MsgTriggerLaunch returns a sample MsgTriggerLaunch
func MsgTriggerLaunch(coordinator string, chainID uint64) launch.MsgTriggerLaunch {
	launchTimeRange := int(launch.DefaultMaxLaunchTime - launch.DefaultMinLaunchTime)
	launchTime := uint64(rand.Intn(launchTimeRange)) + launch.DefaultMinLaunchTime
	return *launch.NewMsgTriggerLaunch(
		coordinator,
		chainID,
		launchTime,
	)
}

// MsgSettleRequest returns a sample MsgSettleRequest
func MsgSettleRequest(coordinator string, chainID, requestID uint64, approve bool) launch.MsgSettleRequest {
	return *launch.NewMsgSettleRequest(
		coordinator,
		chainID,
		requestID,
		approve,
	)
}

// GenesisHash returns a sample sha256 hash of custom genesis for GenesisURL
func GenesisHash() string {
	hash := sha256.Sum256([]byte(String(50)))
	return hex.EncodeToString(hash[:])
}

// LaunchParams returns a sample of params for the launch module
func LaunchParams() launch.Params {
	maxLaunchTime := launch.DefaultMaxLaunchTime - uint64(rand.Intn(10))
	minLaunchTime := uint64(rand.Intn(10)) + launch.DefaultMinLaunchTime
	return launch.Params{
		MinLaunchTime: minLaunchTime,
		MaxLaunchTime: maxLaunchTime,
	}
}

// LaunchGenesisState returns a sample genesis state for the launch module
func LaunchGenesisState(coordinators ...profile.Coordinator) launch.GenesisState {
	for len(coordinators) < 11 {
		coordinators = append(coordinators, Coordinator(Address()))
	}

	chainsLength := 3
	chains := make([]launch.Chain, chainsLength)
	for i := 0; i < chainsLength; i++ {
		chains[i] = Chain(uint64(i), coordinators[i].CoordinatorId)
	}

	return launch.GenesisState{
		ChainList:  chains,
		ChainCount: uint64(len(chains)),
		GenesisAccountList: []launch.GenesisAccount{
			GenesisAccount(chains[0].Id, coordinators[2].Address),
			GenesisAccount(chains[0].Id, coordinators[3].Address),
			GenesisAccount(chains[1].Id, coordinators[4].Address),
		},
		VestingAccountList: []launch.VestingAccount{
			VestingAccount(chains[0].Id, coordinators[5].Address),
			VestingAccount(chains[0].Id, coordinators[6].Address),
			VestingAccount(chains[1].Id, coordinators[7].Address),
		},
		GenesisValidatorList: []launch.GenesisValidator{
			GenesisValidator(chains[0].Id, coordinators[8].Address),
			GenesisValidator(chains[0].Id, coordinators[9].Address),
			GenesisValidator(chains[1].Id, coordinators[10].Address),
		},
		RequestList: []launch.Request{
			Request(chains[0].Id),
			Request(chains[1].Id),
		},
		RequestCountList: []launch.RequestCount{
			{
				ChainID: chains[0].Id,
				Count:   1,
			},
			{
				ChainID: chains[1].Id,
				Count:   2,
			},
		},
		Params: LaunchParams(),
	}
}
