package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"

	"github.com/tendermint/spn/x/genesis/types"

	"testing"
)

func TestNewProposalState(t *testing.T) {
	// Can create a proposal state
	ps := types.NewProposalState()
	require.Equal(t, types.ProposalStatus_PENDING, ps.Status)

	// Can change the state of the proposal
	err := ps.SetStatus(types.ProposalStatus_APPROVED)
	require.NoError(t, err)
	require.Equal(t, types.ProposalStatus_APPROVED, ps.Status)
}

func TestGetType(t *testing.T) {
	// Can get the type of an add account proposal
	addAccountPayload := spnmocks.MockProposalAddAccountPayload()
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		addAccountPayload,
	)
	pType, err := proposal.GetType()
	require.NoError(t, err)
	require.Equal(t, types.ProposalType_ADD_ACCOUNT, pType)

	// Can get the type of an add validator proposal
	addValidatorPayload := spnmocks.MockProposalAddValidatorPayload()
	proposal, _ = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		addValidatorPayload,
	)
	pType, err = proposal.GetType()
	require.NoError(t, err)
	require.Equal(t, types.ProposalType_ADD_VALIDATOR, pType)

	// Error if unrecognized type
	addValidatorPayload = spnmocks.MockProposalAddValidatorPayload()
	proposal, _ = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		addValidatorPayload,
	)
	proposal.Payload = nil
	_, err = proposal.GetType()
	require.Error(t, err)
}

func TestNewProposalChange(t *testing.T) {
	// Test valid payload
	payload := spnmocks.MockProposalChangePayload()
	err := types.ValidateProposalPayloadChange(payload)
	require.NoError(t, err)

	// Can create a proposal for a genesis change
	_, err = types.NewProposalChange(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.NoError(t, err)

	// Prevent invalid change path
	payload.ChangePath = []string{spnmocks.MockRandomString(5), "_"}
	err = types.ValidateProposalPayloadChange(payload)
	require.Error(t, err)

	// Can't create a proposal with an invalid payload
	_, err = types.NewProposalChange(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.Error(t, err)
}

func TestNewProposalAddAccount(t *testing.T) {
	// Test valid payload
	payload := spnmocks.MockProposalAddAccountPayload()
	err := types.ValidateProposalPayloadAddAccount(payload)
	require.NoError(t, err)

	// Can create a proposal for a genesis change
	_, err = types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.NoError(t, err)

	// Prevent add account with invalid address
	payload.Address = sdk.AccAddress([]byte(""))
	err = types.ValidateProposalPayloadAddAccount(payload)
	require.Error(t, err)

	// Prevent add account with invalid address
	payload.Address = sdk.AccAddress([]byte("InvalidAddress"))
	err = types.ValidateProposalPayloadAddAccount(payload)
	require.Error(t, err)


	// Prevent invalid coins allocation
	payload = spnmocks.MockProposalAddAccountPayload()
	payload.Coins = []sdk.Coin{}
	err = types.ValidateProposalPayloadAddAccount(payload)
	require.Error(t, err)

	// Test with non sorted denomination
	payload.Coins = []sdk.Coin{sdk.NewCoin("bbb", sdk.NewInt(10)), sdk.NewCoin("aaa", sdk.NewInt(10))}
	err = types.ValidateProposalPayloadAddAccount(payload)
	require.Error(t, err)

	// Can't create a proposal with an invalid payload
	_, err = types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.Error(t, err)
}

func TestNewProposalAddValidator(t *testing.T) {
	// Test valid payload
	payload := spnmocks.MockProposalAddValidatorPayload()
	peer := "aaa@0.0.0.0:443"
	payload.Peer = peer
	err := types.ValidateProposalPayloadAddValidator(payload)
	require.NoError(t, err)

	//Can get the peer
	retrievedPeer := payload.GetPeer()
	require.NoError(t, err)
	require.Equal(t, peer, retrievedPeer)

	// Can create a proposal for a genesis change
	_, err = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.NoError(t, err)

	// Empty gentx
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.GenTx = []byte{}
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err)

	// Invalid self-delegation
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.SelfDelegation.Denom = ""
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err)

	// Empty address
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.ValidatorAddress = []byte{}
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err)

	// Invalid address
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.ValidatorAddress = []byte("InvalidAddress")
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err)

	// No peer
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.Peer = ""
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err)

	// Can't create a proposal with an invalid payload
	_, err = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.Error(t, err)
}
