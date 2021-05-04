package testing

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	identitykeeper "github.com/tendermint/spn/x/identity/keeper"
	identitytypes "github.com/tendermint/spn/x/identity/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"

	"encoding/json"
	"math/rand"
	"time"
)

// MockCodec mocks a codec for the app that contains the necessary types for proto encoding
func MockCodec() codec.Marshaler {
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	// Register basic message and cryto
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil),
		&types.MsgProposalAddAccount{},
		&types.MsgProposalAddValidator{},
		&types.MsgChainCreate{},
		&types.MsgApprove{},
		&types.MsgReject{},
	)
	staking.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
	authtypes.RegisterInterfaces(interfaceRegistry)

	cdc := codec.NewProtoCodec(interfaceRegistry)

	return cdc
}

// MockGenesisContext mocks the context and the keepers of the launch module for test purposes
func MockGenesisContext() (sdk.Context, *keeper.Keeper) {
	cdc := MockCodec()

	// Store keys
	keys := sdk.NewKVStoreKeys(types.StoreKey, identitytypes.StoreKey)

	// Create a identity keeper
	identityKeeper := identitykeeper.NewKeeper(cdc, keys[identitytypes.StoreKey], keys[identitytypes.MemStoreKey])

	// Create a chat keeper
	genesisKeeper := keeper.NewKeeper(cdc, keys[types.StoreKey], keys[types.MemStoreKey], identityKeeper)

	// Create multiStore in memory
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)

	// Mount stores
	cms.MountStoreWithDB(keys[types.StoreKey], sdk.StoreTypeIAVL, db)
	if err := cms.LoadLatestVersion(); err != nil {
		panic(err.Error())
	}

	// Create context
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return ctx, genesisKeeper
}

// MockGenesisQueryClient mocks a query client for the launch module
func MockGenesisQueryClient(ctx sdk.Context, k *keeper.Keeper) types.QueryClient {
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, codectypes.NewInterfaceRegistry())
	types.RegisterQueryServer(queryHelper, k)
	return types.NewQueryClient(queryHelper)
}

// MockGenesis mocks a genesis structure
func MockGenesis() []byte {
	var genesisObject tmtypes.GenesisDoc

	consensusParam := tmtypes.DefaultConsensusParams()
	genesisObject.ChainID = MockRandomAlphaString(10)
	genesisObject.ConsensusParams = consensusParam

	// Create a mock app state
	testMbm := module.NewBasicManager(genutil.AppModuleBasic{})
	cdc := MockCodec()
	appState, err := json.MarshalIndent(testMbm.DefaultGenesis(cdc), "", " ")
	if err != nil {
		panic("Cannot unmarshal marshal app state")
	}
	genesisObject.AppState = appState

	// JSON encode the genesis
	genesis, err := tmjson.Marshal(genesisObject)
	if err != nil {
		panic("Cannot marshal genesis")
	}

	genesisBytes, err := json.Marshal(genesis)
	if err != nil {
		panic(fmt.Sprintf("cannot marshal genesis: %v", err.Error()))
	}

	return genesisBytes
}

// MockChain mocks a chain information structure
func MockChain() *types.Chain {
	chain, _ := types.NewChain(
		MockRandomAlphaString(5)+"-"+MockRandomAlphaString(5),
		MockRandomString(20),
		MockRandomString(20),
		MockRandomString(20),
		time.Now(),
		"",
		"",
	)

	chain.Peers = []string(nil)

	return chain
}

// MockProposal mocks a proposal
func MockProposal() *types.Proposal {
	proposal, _ := types.NewProposalAddAccount(
		MockProposalInformation(),
		MockProposalAddAccountPayload(),
	)
	return proposal
}

// MockProposalList mocks a list of proposal ID
func MockProposalList() *types.ProposalList {
	return &types.ProposalList{
		ProposalIDs: []int32{
			rand.Int31n(10000),
			rand.Int31n(10000),
			rand.Int31n(10000),
			rand.Int31n(10000),
			rand.Int31n(10000),
			rand.Int31n(10000),
		},
	}
}

// MockProposalChangePayload mocks a valid payload
func MockProposalChangePayload() *types.ProposalChangePayload {
	return types.NewProposalChangePayload(
		[]string{MockRandomString(5), MockRandomString(5), MockRandomString(5)},
		MockRandomString(10),
	)
}

// MockProposalAddAccountPayload mocks a valid payload
func MockProposalAddAccountPayload() *types.ProposalAddAccountPayload {
	stake := sdk.NewCoin("stake", sdk.NewInt(100000))
	token := sdk.NewCoin("token", sdk.NewInt(100000))

	return types.NewProposalAddAccountPayload(
		MockAccAddress(),
		sdk.NewCoins(stake, token),
	)
}

// MockProposalAddValidatorPayload mocks a valid payload
func MockProposalAddValidatorPayload() *types.ProposalAddValidatorPayload {
	genTx := MockGenTx()
	_, validatorAddress := MockValAddress()
	selfDelegation := sdk.NewCoin("stake", sdk.NewInt(10000))

	return types.NewProposalAddValidatorPayload(
		genTx,
		validatorAddress,
		selfDelegation,
		MockRandomString(20),
	)
}

// MockProposalInformation mocks information for a proposal
func MockProposalInformation() *types.ProposalInformation {
	return types.NewProposalInformation(
		MockRandomString(5)+"-"+MockRandomString(5),
		int32(rand.Intn(10)),
		MockRandomString(10),
		time.Now(),
	)
}

// MockGenTx mocks a gentx transaction
// A gentx is not interpreted on-chain and can be any data
func MockGenTx() []byte {
	return []byte(MockRandomString(250))
}
