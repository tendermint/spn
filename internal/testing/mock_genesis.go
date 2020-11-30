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
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
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

	"github.com/tendermint/spn/x/genesis/keeper"
	"github.com/tendermint/spn/x/genesis/types"

	"encoding/json"
	"math/rand"
	"time"
)

// MockCodec mocks a codec for the app that contains the necessary types for proto enconding
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

// MockGenesisContext mocks the context and the keepers of the genesis module for test purposes
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
	cms.LoadLatestVersion()

	// Create context
	ctx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())

	return ctx, genesisKeeper
}

// MockGenesisQueryClient mocks a query client for the genesis module
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
			int32(rand.Intn(10000)),
			int32(rand.Intn(10000)),
			int32(rand.Intn(10000)),
			int32(rand.Intn(10000)),
			int32(rand.Intn(10000)),
			int32(rand.Intn(10000)),
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
func MockGenTx() []byte {
	selfDelegation := sdk.NewCoin("stake", sdk.NewInt(10000))
	return MockGenTxWithDelegation(selfDelegation)
}

// MockGenTxWithDelegation mocks a gentx transaction with a custom self-delegation
func MockGenTxWithDelegation(selfDelegation sdk.Coin) []byte {
	privKey, opAddress := MockValAddress()

	// Create validator message
	message, err := staking.NewMsgCreateValidator(
		opAddress,
		MockPubKey(),
		selfDelegation,
		MockDescription(),
		MockCommissionRates(),
		sdk.NewInt(1),
	)

	if err != nil {
		panic(err)
	}

	// txConfig
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	marshaler := codec.NewProtoCodec(interfaceRegistry)
	txConfig := authtx.NewTxConfig(marshaler, authtx.DefaultSignModes)

	// Tx message
	var genTx tx.Tx
	any, err := codectypes.NewAnyWithValue(message)
	if err != nil {
		panic("Can't encode MsgCreateValidator")
	}
	if genTx.Body == nil {
		genTx.Body = &tx.TxBody{}
	}
	genTx.Body.Messages = []*codectypes.Any{any}

	// Signature
	signMode := txConfig.SignModeHandler().DefaultMode()
	sig := signing.SignatureV2{
		PubKey: privKey.PubKey(),
		Data: &signing.SingleSignatureData{
			SignMode: signMode,
		},
		Sequence: 0,
	}
	signerInfos := make([]*tx.SignerInfo, 1)
	rawSigs := make([][]byte, 1)
	var modeInfo *tx.ModeInfo
	modeInfo, rawSigs[0] = authtx.SignatureDataToModeInfoAndSig(sig.Data)
	any, err = authtx.PubKeyToAny(sig.PubKey)
	if err != nil {
		panic("Can't encode PubKey")
	}
	signerInfos[0] = &tx.SignerInfo{
		PublicKey: any,
		ModeInfo:  modeInfo,
		Sequence:  sig.Sequence,
	}
	if genTx.AuthInfo == nil {
		genTx.AuthInfo = &tx.AuthInfo{}
	}
	genTx.AuthInfo.SignerInfos = signerInfos
	genTx.Signatures = rawSigs

	// Other info
	genTx.Body.Memo = MockRandomAlphaString(20)
	if genTx.AuthInfo.Fee == nil {
		genTx.AuthInfo.Fee = &tx.Fee{}
	}
	genTx.AuthInfo.Fee.Amount = MockCoins()
	if genTx.AuthInfo.Fee == nil {
		genTx.AuthInfo.Fee = &tx.Fee{}
	}
	genTx.AuthInfo.Fee.GasLimit = 1000000

	// Marshal gentx
	cdc := MockCodec()
	jsonCodec, ok := cdc.(codec.JSONMarshaler)
	if !ok {
		panic("Can't get json codec")
	}
	gentxBytes, err := jsonCodec.MarshalJSON(&genTx)
	if err != nil {
		panic(fmt.Sprintf("cannot marshal gentx: %v", err.Error()))
	}

	return gentxBytes
}
