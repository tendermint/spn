package testing

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	tx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/tendermint/spn/x/genesis/types"

	"math/rand"
	"time"
)

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
	return types.NewProposalAddValidatorPayload(genTx)
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