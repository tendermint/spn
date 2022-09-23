package sample

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/pkg/chainid"
	launch "github.com/tendermint/spn/x/launch/types"
)

// Metadata returns sample metadata bytes
func Metadata(r *rand.Rand, len int) []byte {
	return Bytes(r, len)
}

// GenesisChainID returns a sample chain id
func GenesisChainID(r *rand.Rand) string {
	chainName := AlphaString(r, 5)
	number := Uint64(r)
	return chainid.NewGenesisChainID(chainName, number)
}

// Chain returns a sample Chain
func Chain(r *rand.Rand, id, coordinatorID uint64) launch.Chain {
	return launch.Chain{
		LaunchID:        id,
		CoordinatorID:   coordinatorID,
		GenesisChainID:  GenesisChainID(r),
		CreatedAt:       Duration(r).Milliseconds(),
		SourceURL:       String(r, 10),
		SourceHash:      String(r, 10),
		LaunchTime:      ZeroTime(),
		LaunchTriggered: false,
		InitialGenesis:  launch.NewDefaultInitialGenesis(),
		AccountBalance:  Coins(r),
		Metadata:        Metadata(r, 20),
	}
}

// GenesisAccount returns a sample GenesisAccount
func GenesisAccount(r *rand.Rand, launchID uint64, address string) launch.GenesisAccount {
	return launch.GenesisAccount{
		LaunchID: launchID,
		Address:  address,
		Coins:    Coins(r),
	}
}

// VestingOptions returns a sample VestingOptions
func VestingOptions(r *rand.Rand) launch.VestingOptions {
	balance := Coins(r)
	return *launch.NewDelayedVesting(balance, balance, Time(r))
}

// VestingAccount returns a sample VestingAccount
func VestingAccount(r *rand.Rand, launchID uint64, address string) launch.VestingAccount {
	return launch.VestingAccount{
		LaunchID:       launchID,
		Address:        address,
		VestingOptions: VestingOptions(r),
	}
}

// AccountRemoval returns a sample AccountRemoval
func AccountRemoval(address string) *launch.AccountRemoval {
	return &launch.AccountRemoval{
		Address: address,
	}
}

// ParamChange returns a sample ParamChange
func ParamChange(r *rand.Rand, launchID uint64) launch.ParamChange {
	return launch.ParamChange{
		LaunchID: launchID,
		Module:   AlphaString(r, 5),
		Param:    AlphaString(r, 5),
		Value:    Bytes(r, 5),
	}
}

// GenesisValidator returns a sample GenesisValidator
func GenesisValidator(r *rand.Rand, launchID uint64, address string) launch.GenesisValidator {
	return launch.GenesisValidator{
		LaunchID:       launchID,
		Address:        address,
		GenTx:          Bytes(r, 200),
		ConsPubKey:     PubKey(r).Bytes(),
		SelfDelegation: Coin(r),
		Peer:           GenesisValidatorPeer(r),
	}
}

func GenesisValidatorPeer(r *rand.Rand) launch.Peer {
	return launch.Peer{
		Id: String(r, 10),
		Connection: &launch.Peer_TcpAddress{
			TcpAddress: fmt.Sprintf("%s@%s", String(r, 5), String(r, 10)),
		},
	}
}

// RequestWithContent creates a launch request object with launch id and content
func RequestWithContent(r *rand.Rand, launchID uint64, content launch.RequestContent) launch.Request {
	return launch.Request{
		RequestID: 1,
		LaunchID:  launchID,
		Creator:   Address(r),
		CreatedAt: Duration(r).Milliseconds(),
		Content:   content,
	}
}

// RequestWithContentAndCreator creates a launch request object with launch id and content and creator
func RequestWithContentAndCreator(r *rand.Rand, launchID uint64, content launch.RequestContent, creator string) launch.Request {
	return launch.Request{
		RequestID: 1,
		LaunchID:  launchID,
		Creator:   creator,
		CreatedAt: Duration(r).Milliseconds(),
		Content:   content,
	}
}

// RequestContent returns a request content with Genesis Account
func RequestContent(r *rand.Rand, launchID uint64) launch.RequestContent {
	return launch.NewGenesisAccount(launchID, Address(r), Coins(r))
}

// AllRequestContents creates all contents types for request
func AllRequestContents(r *rand.Rand, launchID uint64, genesis, vesting, validator string) []launch.RequestContent {
	return []launch.RequestContent{
		launch.NewGenesisAccount(launchID, genesis, Coins(r)),
		launch.NewAccountRemoval(genesis),
		launch.NewVestingAccount(launchID, vesting, VestingOptions(r)),
		launch.NewAccountRemoval(vesting),
		launch.NewGenesisValidator(launchID, validator, Bytes(r, 300), Bytes(r, 30), Coin(r), GenesisValidatorPeer(r)),
		launch.NewValidatorRemoval(validator),
	}
}

// GenesisAccountContent returns a sample GenesisAccount request content
func GenesisAccountContent(r *rand.Rand, launchID uint64, address string) launch.RequestContent {
	return launch.NewGenesisAccount(launchID, address, Coins(r))
}

// Request returns a sample Request
func Request(r *rand.Rand, launchID uint64, address string) launch.Request {
	content := GenesisAccountContent(r, launchID, address)
	req := RequestWithContent(r, launchID, content)
	req.Status = launch.Request_PENDING // set to pending by default
	return req
}

// MsgCreateChain returns a sample MsgCreateChain
func MsgCreateChain(r *rand.Rand, coordAddress, genesisURL string, hasCampaign bool, campaignID uint64) launch.MsgCreateChain {
	initialGenesis := launch.NewDefaultInitialGenesis()
	var genesisHash string
	if len(genesisURL) > 0 {
		genesisHash = GenesisHash(r)
		ig := launch.NewGenesisURL(
			genesisURL,
			genesisHash,
		)
		initialGenesis = ig
	}

	return *launch.NewMsgCreateChain(
		coordAddress,
		GenesisChainID(r),
		String(r, 10),
		String(r, 10),
		initialGenesis,
		hasCampaign,
		campaignID,
		Coins(r),
		Metadata(r, 20),
	)
}

// MsgEditChain returns a sample MsgEditChain
func MsgEditChain(
	r *rand.Rand,
	coordAddress string,
	launchID uint64,
	setCampaignID bool,
	campaignID uint64,
	modifyMetadata bool,
) launch.MsgEditChain {
	var metadata []byte
	if modifyMetadata {
		metadata = Metadata(r, 20)
	}

	return *launch.NewMsgEditChain(
		coordAddress,
		launchID,
		setCampaignID,
		campaignID,
		metadata,
	)
}

// MsgUpdateLaunchInformation returns a sample MsgUpdateLaunchInformation
func MsgUpdateLaunchInformation(
	r *rand.Rand,
	coordAddress string,
	launchID uint64,
	modifyGenesisChainID,
	modifySource,
	modifyInitialGenesis,
	genesisURL bool,
) launch.MsgUpdateLaunchInformation {
	var genesisChainID, sourceURL, sourceHash string
	var initialGenesis *launch.InitialGenesis

	if modifyGenesisChainID {
		genesisChainID = GenesisChainID(r)
	}
	if modifySource {
		sourceURL, sourceHash = String(r, 30), String(r, 10)
	}
	if modifyInitialGenesis {
		if genesisURL {
			newGenesisURL := launch.NewGenesisURL(String(r, 30), GenesisHash(r))
			initialGenesis = &newGenesisURL
		} else {
			newDefault := launch.NewDefaultInitialGenesis()
			initialGenesis = &newDefault
		}
	}

	return *launch.NewMsgUpdateLaunchInformation(
		coordAddress,
		launchID,
		genesisChainID,
		sourceURL,
		sourceHash,
		initialGenesis,
	)
}

// MsgSendRequestWithAddAccount returns a sample MsgSendRequest with AddAccount request content
func MsgSendRequestWithAddAccount(r *rand.Rand, creator, address string, launchID uint64) launch.MsgSendRequest {
	return *launch.NewMsgSendRequest(
		creator,
		launchID,
		launch.NewGenesisAccount(launchID, address, Coins(r)),
	)
}

// MsgSendRequestWithAddVestingAccount returns a sample MsgSendRequest with AddVestingAccount request content
func MsgSendRequestWithAddVestingAccount(r *rand.Rand, creator, address string, launchID uint64) launch.MsgSendRequest {
	return *launch.NewMsgSendRequest(
		creator,
		launchID,
		launch.NewVestingAccount(launchID, address, VestingOptions(r)),
	)
}

// MsgSendRequestWithRemoveAccount returns a sample MsgSendRequest with RemoveAccount request content
func MsgSendRequestWithRemoveAccount(creator, address string, launchID uint64) launch.MsgSendRequest {
	return *launch.NewMsgSendRequest(
		creator,
		launchID,
		launch.NewAccountRemoval(address),
	)
}

// MsgSendRequestWithRemoveValidator returns a sample MsgSendRequest with RemoveValidator request content
func MsgSendRequestWithRemoveValidator(creator, validatorAddr string, launchID uint64) launch.MsgSendRequest {
	return *launch.NewMsgSendRequest(
		creator,
		launchID,
		launch.NewValidatorRemoval(validatorAddr),
	)
}

// MsgSendRequestWithParamChange returns a sample MsgSendRequest with ParamChange request content
func MsgSendRequestWithParamChange(r *rand.Rand, creator string, launchID uint64) launch.MsgSendRequest {
	pc := ParamChange(r, launchID)

	return *launch.NewMsgSendRequest(
		creator,
		launchID,
		launch.NewParamChange(pc.LaunchID, pc.Module, pc.Param, pc.Value),
	)
}

// MsgSendRequestWithAddValidator returns a sample MsgSendRequest with AddValidator request content
func MsgSendRequestWithAddValidator(r *rand.Rand, creator, address string, launchID uint64) launch.MsgSendRequest {
	return *launch.NewMsgSendRequest(
		creator,
		launchID,
		launch.NewGenesisValidator(
			launchID,
			address,
			Bytes(r, 500),
			Bytes(r, 30),
			Coin(r),
			GenesisValidatorPeer(r),
		),
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
func MsgTriggerLaunch(r *rand.Rand, coordinator string, launchID uint64, currentTime time.Time) launch.MsgTriggerLaunch {
	// Get a random duration between min and max launch time
	launchTimeRange := launch.DefaultMaxLaunchTime.Milliseconds() - launch.DefaultMinLaunchTime.Milliseconds()
	remainingTime := launch.DefaultMinLaunchTime + time.Millisecond*time.Duration(r.Int63n(launchTimeRange))

	return *launch.NewMsgTriggerLaunch(
		coordinator,
		launchID,
		currentTime.Add(remainingTime),
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
func GenesisHash(r *rand.Rand) string {
	hash := sha256.Sum256([]byte(String(r, 50)))
	return hex.EncodeToString(hash[:])
}

// LaunchParams returns a sample of params for the launch module
func LaunchParams(r *rand.Rand) launch.Params {
	maxLaunchTime := launch.DefaultMaxLaunchTime - time.Second*time.Duration(r.Int63n(10))
	minLaunchTime := launch.DefaultMinLaunchTime + time.Second*time.Duration(r.Int63n(10))

	// assign random small amount of staking denom
	chainCreationFee := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, r.Int63n(100)+1))
	requestFee := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, r.Int63n(100)+1))

	return launch.NewParams(minLaunchTime, maxLaunchTime, launch.DefaultRevertDelay, chainCreationFee, requestFee)
}

// LaunchGenesisState returns a sample genesis state for the launch module
func LaunchGenesisState(r *rand.Rand, addresses ...string) launch.GenesisState {
	for len(addresses) < 11 {
		addresses = append(addresses, Address(r))
	}
	return launch.GenesisState{
		Chains: []launch.Chain{
			Chain(r, 0, 0),
			Chain(r, 1, 1),
		},
		ChainCounter: 2,
		GenesisAccounts: []launch.GenesisAccount{
			GenesisAccount(r, 0, addresses[0]),
			GenesisAccount(r, 0, addresses[1]),
			GenesisAccount(r, 1, addresses[2]),
		},
		VestingAccounts: []launch.VestingAccount{
			VestingAccount(r, 0, addresses[3]),
			VestingAccount(r, 0, addresses[4]),
			VestingAccount(r, 1, addresses[5]),
		},
		GenesisValidators: []launch.GenesisValidator{
			GenesisValidator(r, 0, addresses[6]),
			GenesisValidator(r, 0, addresses[7]),
			GenesisValidator(r, 1, addresses[8]),
		},
		Requests: []launch.Request{
			Request(r, 0, addresses[9]),
			Request(r, 1, addresses[10]),
		},
		RequestCounters: []launch.RequestCounter{
			{
				LaunchID: 0,
				Counter:  1,
			},
			{
				LaunchID: 1,
				Counter:  2,
			},
		},
		Params: LaunchParams(r),
	}
}
