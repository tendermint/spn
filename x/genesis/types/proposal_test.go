package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"

	// staking "github.com/cosmos/cosmos-sdk/x/staking/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"

	"github.com/tendermint/spn/x/genesis/types"

	"fmt"
	"testing"
)

func TestNewProposalState(t *testing.T) {
	// Can create a proposal state
	ps := types.NewProposalState()
	require.Equal(t, types.ProposalState_PENDING, ps.Status, "NewProposalState should create a pending proposal")

	// Can change the state of the proposal
	err := ps.SetStatus(types.ProposalState_APPROVED)
	require.NoError(t, err, "SetStatus should set status of the proposal")
	require.Equal(t, types.ProposalState_APPROVED, ps.Status, "SetStatus should set status of the proposal")
}

func TestNewProposalChange(t *testing.T) {
	// Test valid payload
	payload := spnmocks.MockProposalChangePayload()
	err := types.ValidateProposalPayloadChange(payload)
	require.NoError(t, err, "ValidateProposalPayloadChange should return no error")

	// Can create a proposal for a genesis change
	_, err = types.NewProposalChange(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.NoError(t, err, "NewProposalChange should create a new proposal")

	// Prevent invalid change path
	payload.ChangePath = []string{spnmocks.MockRandomString(5), "_"}
	err = types.ValidateProposalPayloadChange(payload)
	require.Error(t, err, "ValidateProposalPayloadChange should return error on invalid payload")

	// Can't create a proposal with an invalid payload
	_, err = types.NewProposalChange(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.Error(t, err, "NewProposalChange should prevent invalid payload")
}

func TestNewProposalAddAccount(t *testing.T) {
	// Test valid payload
	payload := spnmocks.MockProposalAddAccountPayload()
	err := types.ValidateProposalPayloadAddAccount(payload)
	require.NoError(t, err, fmt.Sprintf("ValidateProposalPayloadAddAccount should return no error: %v", err))

	// Can create a proposal for a genesis change
	_, err = types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.NoError(t, err, "NewProposalAddAccount should create a new proposal")

	// Prevent add account with invalid address
	payload.Address = sdk.AccAddress([]byte(""))
	err = types.ValidateProposalPayloadAddAccount(payload)
	require.Error(t, err, "ValidateProposalPayloadAddAccount should return error on invalid payload")

	// Prevent invalid coins allocation
	payload = spnmocks.MockProposalAddAccountPayload()
	payload.Coins = []sdk.Coin{}
	err = types.ValidateProposalPayloadAddAccount(payload)
	require.Error(t, err, "ValidateProposalPayloadAddAccount should return error on invalid payload")

	// Test with non sorted denomination
	payload.Coins = []sdk.Coin{sdk.NewCoin("bbb", sdk.NewInt(10)), sdk.NewCoin("aaa", sdk.NewInt(10))}
	err = types.ValidateProposalPayloadAddAccount(payload)
	require.Error(t, err, "ValidateProposalPayloadAddAccount should return error on invalid payload")

	// Can't create a proposal with an invalid payload
	_, err = types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.Error(t, err, "NewProposalAddAccount should prevent invalid payload")
}

func TestNewProposalAddValidator(t *testing.T) {
	// Test valid payload
	payload := spnmocks.MockProposalAddValidatorPayload()
	peer := "aaa@0.0.0.0:443"
	payload.Peer = peer
	err := types.ValidateProposalPayloadAddValidator(payload)
	require.NoError(t, err, fmt.Sprintf("ValidateProposalPayloadAddValidator should return no error: %v", err))

	//Can get the peer
	retrievedPeer := payload.GetPeer()
	require.NoError(t, err, "GetPeer should get the validator peer")
	require.Equal(t, peer, retrievedPeer, "GetPeer should get the validator peer")

	// Can create a proposal for a genesis change
	_, err = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.NoError(t, err, "NewProposalAddValidator should create a new proposal")

	// Invalid tx
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.GenTx.Body = nil
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// No message
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.GenTx.Body.Messages = []*codectypes.Any{}
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// Invalid message
	payload = spnmocks.MockProposalAddValidatorPayload()
	message := staking.NewMsgCreateValidator(
		sdk.ValAddress(""),
		spnmocks.MockPubKey(),
		spnmocks.MockCoin(),
		spnmocks.MockDescription(),
		spnmocks.MockCommissionRates(),
		sdk.NewInt(1),
	)
	payload.GenTx.Body.Messages[0], _ = codectypes.NewAnyWithValue(message)
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// No peer
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.Peer = "" // Peer is inside the memo
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// Can't create a proposal with an invalid payload
	_, err = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.Error(t, err, "NewProposalAddValidator should prevent invalid payload")
}
