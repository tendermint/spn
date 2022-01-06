package sample

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/tendermint/spn/pkg/chainid"
	launch "github.com/tendermint/spn/x/launch/types"
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
		LaunchID:        id,
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
func GenesisAccount(launchID uint64, address string) launch.GenesisAccount {
	return launch.GenesisAccount{
		LaunchID: launchID,
		Address:  address,
		Coins:    Coins(),
	}
}

// VestingOptions returns a sample VestingOptions
func VestingOptions() launch.VestingOptions {
	balance := Coins()
	return *launch.NewDelayedVesting(balance, balance, time.Now().Unix())
}

// VestingAccount returns a sample VestingAccount
func VestingAccount(launchID uint64, address string) launch.VestingAccount {
	return launch.VestingAccount{
		LaunchID:       launchID,
		Address:        address,
		VestingOptions: VestingOptions(),
	}
}

// AccountRemoval returns a sample AccountRemoval
func AccountRemoval(address string) *launch.AccountRemoval {
	return &launch.AccountRemoval{
		Address: address,
	}
}

// GenesisValidator returns a sample GenesisValidator
func GenesisValidator(launchID uint64, address string) launch.GenesisValidator {
	return launch.GenesisValidator{
		LaunchID:       launchID,
		Address:        address,
		GenTx:          Bytes(200),
		ConsPubKey:     Bytes(10),
		SelfDelegation: Coin(),
		Peer:           GenesisValidatorPeer(),
	}
}
func GenesisValidatorPeer() *launch.Peer {
	return &launch.Peer{
		Id: String(10),
		Connection: &launch.Peer_TcpAddress{
			TcpAddress: fmt.Sprintf("%s@%s", String(5), String(10)),
		},
	}
}

// ValidatorRemoval returns a sample ValidatorRemoval
func ValidatorRemoval(address string) launch.ValidatorRemoval {
	return launch.ValidatorRemoval{
		ValAddress: address,
	}
}

// RequestWithContent creates a launch request object with launch id and content
func RequestWithContent(launchID uint64, content launch.RequestContent) launch.Request {
	return launch.Request{
		RequestID: 1,
		LaunchID:  launchID,
		Creator:   Address(),
		CreatedAt: time.Now().Unix(),
		Content:   content,
	}
}

// RequestWithContentAndCreator creates a launch request object with launch id and content and creator
func RequestWithContentAndCreator(launchID uint64, content launch.RequestContent, creator string) launch.Request {
	return launch.Request{
		RequestID: 1,
		LaunchID:  launchID,
		Creator:   creator,
		CreatedAt: time.Now().Unix(),
		Content:   content,
	}
}

// AllRequestContents creates all contents types for request
func AllRequestContents(launchID uint64, genesis, vesting, validator string) []launch.RequestContent {
	return []launch.RequestContent{
		launch.NewGenesisAccount(launchID, genesis, Coins()),
		launch.NewAccountRemoval(genesis),
		launch.NewVestingAccount(launchID, vesting, VestingOptions()),
		launch.NewAccountRemoval(vesting),
		launch.NewGenesisValidator(launchID, validator, Bytes(300), Bytes(30), Coin(), GenesisValidatorPeer()),
		launch.NewValidatorRemoval(validator),
	}
}

// GenesisAccountContent returns a sample GenesisAccount request content
func GenesisAccountContent(launchID uint64, address string) launch.RequestContent {
	return launch.NewGenesisAccount(launchID, address, Coins())
}

// Request returns a sample Request
func Request(launchID uint64, address string) launch.Request {
	content := GenesisAccountContent(launchID, address)
	return RequestWithContent(launchID, content)
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
	launchID uint64,
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
		launchID,
		genesisChainID,
		sourceURL,
		sourceHash,
		initialGenesis,
	)
}

// MsgRequestAddAccount returns a sample MsgRequestAddAccount
func MsgRequestAddAccount(creator, address string, launchID uint64) launch.MsgRequestAddAccount {
	return *launch.NewMsgRequestAddAccount(
		creator,
		launchID,
		address,
		Coins(),
	)
}

// MsgRequestAddVestingAccount returns a sample MsgRequestAddVestingAccount
func MsgRequestAddVestingAccount(creator, address string, launchID uint64) launch.MsgRequestAddVestingAccount {
	return *launch.NewMsgRequestAddVestingAccount(
		creator,
		launchID,
		address,
		VestingOptions(),
	)
}

// MsgRequestRemoveAccount returns a sample MsgRequestRemoveAccount
func MsgRequestRemoveAccount(creator, address string, launchID uint64) launch.MsgRequestRemoveAccount {
	return *launch.NewMsgRequestRemoveAccount(
		creator,
		launchID,
		address,
	)
}

// MsgRequestRemoveValidator returns a sample MsgRequestRemoveValidator
func MsgRequestRemoveValidator(creator, validatorAddr string, launchID uint64) launch.MsgRequestRemoveValidator {
	return *launch.NewMsgRequestRemoveValidator(
		creator,
		launchID,
		validatorAddr,
	)
}

// MsgRequestAddValidator returns a sample MsgRequestAddValidator
func MsgRequestAddValidator(creator, address string, launchID uint64) launch.MsgRequestAddValidator {
	return *launch.NewMsgRequestAddValidator(
		creator,
		launchID,
		address,
		Bytes(500),
		Bytes(30),
		Coin(),
		GenesisValidatorPeer(),
	)
}

// MsgRevertLaunch returns a sample MsgRevertLaunch
func MsgRevertLaunch(coordinator string, launchID uint64) launch.MsgRevertLaunch {
	return *launch.NewMsgRevertLaunch(
		coordinator,
		launchID,
	)
}

// MsgTriggerLaunch returns a sample MsgTriggerLaunch
func MsgTriggerLaunch(coordinator string, launchID uint64) launch.MsgTriggerLaunch {
	launchTimeRange := int(launch.DefaultMaxLaunchTime - launch.DefaultMinLaunchTime)
	launchTime := uint64(rand.Intn(launchTimeRange)) + launch.DefaultMinLaunchTime
	return *launch.NewMsgTriggerLaunch(
		coordinator,
		launchID,
		launchTime,
	)
}

// MsgSettleRequest returns a sample MsgSettleRequest
func MsgSettleRequest(coordinator string, launchID, requestID uint64, approve bool) launch.MsgSettleRequest {
	return *launch.NewMsgSettleRequest(
		coordinator,
		launchID,
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
func LaunchGenesisState(addresses ...string) launch.GenesisState {
	for len(addresses) < 11 {
		addresses = append(addresses, Address())
	}
	return launch.GenesisState{
		ChainList: []launch.Chain{
			Chain(0, 0),
			Chain(1, 1),
		},
		ChainCounter: 2,
		GenesisAccountList: []launch.GenesisAccount{
			GenesisAccount(0, addresses[0]),
			GenesisAccount(0, addresses[1]),
			GenesisAccount(1, addresses[2]),
		},
		VestingAccountList: []launch.VestingAccount{
			VestingAccount(0, addresses[3]),
			VestingAccount(0, addresses[4]),
			VestingAccount(1, addresses[5]),
		},
		GenesisValidatorList: []launch.GenesisValidator{
			GenesisValidator(0, addresses[6]),
			GenesisValidator(0, addresses[7]),
			GenesisValidator(1, addresses[8]),
		},
		RequestList: []launch.Request{
			Request(0, addresses[9]),
			Request(1, addresses[10]),
		},
		RequestCounterList: []launch.RequestCounter{
			{
				LaunchID: 0,
				Counter:  1,
			},
			{
				LaunchID: 1,
				Counter:  2,
			},
		},
		Params: LaunchParams(),
	}
}
