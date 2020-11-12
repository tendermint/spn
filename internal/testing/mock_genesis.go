package testing

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
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

	"math/rand"
	"time"
	"encoding/json"
)

// MockGenesisContext mocks the context and the keepers of the genesis module for test purposes
func MockGenesisContext() (sdk.Context, *keeper.Keeper) {
	// Codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(interfaceRegistry)

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

	// AppSate can be any json, let reuse consensus default param as a sample
	appState, err := json.Marshal(*consensusParam)
	if err != nil {
		panic("Cannot unmarshal consensusParam")
	}
	genesisObject.AppState = appState

	// JSON encode the genesis
	genesis, err := tmjson.Marshal(genesisObject)
	if err != nil {
		panic("Cannot marshal genesis")
	}

	return genesis
}

// MockChain mocks a chain information structure
func MockChain() *types.Chain {
	chain, _ := types.NewChain(
		MockRandomAlphaString(5)+"-"+MockRandomAlphaString(5),
		MockRandomString(20),
		MockRandomString(20),
		MockRandomString(20),
		time.Now(),
		MockGenesis(),
	)

	chain.Peers = []string(nil)

	return chain
}

// MockProposal mocks a proposal
func MockProposal() *types.Proposal {
	proposal, _ := types.NewProposalChange(
		MockProposalInformation(),
		MockProposalChangePayload(),
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
	return types.NewProposalAddAccountPayload(
		MockAccAddress(),
		MockCoins(),
	)
}

// MockProposalAddValidatorPayload mocks a valid payload
func MockProposalAddValidatorPayload() *types.ProposalAddValidatorPayload {
	genTx := MockGenTx()
	return types.NewProposalAddValidatorPayload(genTx, MockRandomString(20))
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

// MockProposalVote mocks a vote for a genesis proposal
func MockProposalVote(voter string) *types.Vote {
	voteValue := types.Vote_REJECT

	if r := rand.Intn(10); r > 5 {
		voteValue = types.Vote_APPROVE
	}

	vote, _ := types.NewVote(
		int32(rand.Intn(10)),
		voter,
		time.Now(),
		voteValue,
	)
	return vote
}

// MockGenTx mocks a gentx transaction
func MockGenTx() tx.Tx {
	privKey, opAddress := MockValAddress()

	// Create validator message
	message := staking.NewMsgCreateValidator(
		opAddress,
		MockPubKey(),
		MockCoin(),
		MockDescription(),
		MockCommissionRates(),
		sdk.NewInt(1),
	)

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

	// TODO: Add a correct signature in the mock
	//// Sign with all info
	//signerData := authsign.SignerData{
	//	ChainID:       MockRandomAlphaString(5),
	//	AccountNumber: 0,
	//	Sequence:      0,
	//}
	//signBytes, err := txConfig.SignModeHandler().GetSignBytes(signMode, signerData, &genTx)
	//if err != nil {
	//	panic("Can't get sign bytes")
	//}
	//signature, err := privKey.Sign(signBytes)
	//if err != nil {
	//	panic("Can't sign transaction")
	//}
	//sig.Data.(*signing.SingleSignatureData).Signature = signature
	//modeInfo, rawSigs[0] = authtx.SignatureDataToModeInfoAndSig(sig.Data)
	//any, err = authtx.PubKeyToAny(sig.PubKey)
	//if err != nil {
	//	panic("Can't encode public key")
	//}
	//signerInfos[0] = &tx.SignerInfo{
	//	PublicKey: any,
	//	ModeInfo:  modeInfo,
	//	Sequence:  sig.Sequence,
	//}
	//genTx.AuthInfo.SignerInfos = signerInfos
	//genTx.Signatures = rawSigs

	return genTx
}
