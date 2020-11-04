package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
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
	require.Equal(t, 0, len(ps.Votes), "NewProposalState should create a proposal with no vote")

	// Can append a vote in the proposal state
	fooVote := spnmocks.MockProposalVote("foo")
	ps.AppendVote(fooVote)
	ps.AppendVote(spnmocks.MockProposalVote("bar"))
	require.Equal(t, 2, len(ps.Votes), "AppendVote should append votes")
	require.Equal(t, fooVote, ps.Votes["foo"], "AppendVote should append votes")

	// Prevent voting twice for the same identity
	err := ps.AppendVote(fooVote)
	require.Error(t, err, "AppendVote shoud prevent voting twice")

	// Can change the state of the proposal
	err = ps.SetStatus(types.ProposalState_APPROVED)
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
	payload.ChangePath = spnmocks.MockRandomString(5) + "_" + spnmocks.MockRandomString(5)
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
	err := types.ValidateProposalPayloadAddValidator(payload)
	require.NoError(t, err, fmt.Sprintf("ValidateProposalPayloadAddValidator should return no error: %v", err))

	// Can create a proposal for a genesis change
	_, err = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.NoError(t, err, "NewProposalAddValidator should create a new proposal")

	// Invalid operator address
	payload.OperatorAddress = sdk.ValAddress([]byte(""))
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// Invalid consensus key
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.ConsensusPubKey = []byte("")
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// Invalid self delegation
	payload = spnmocks.MockProposalAddValidatorPayload()
	invalidCoin := sdk.NewCoin("atom", sdk.NewInt(10))
	invalidCoin.Denom = ""
	payload.SelfDelegation = &invalidCoin
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// Empty gentx
	payload = spnmocks.MockProposalAddValidatorPayload()
	payload.GenTx = []byte("")
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// Invalid commissions
	payload = spnmocks.MockProposalAddValidatorPayload()
	commissions := staking.NewCommissionRates(
		sdk.ZeroDec(),
		sdk.NewDec(10),
		sdk.ZeroDec(),
	)
	payload.Commissions = &commissions
	err = types.ValidateProposalPayloadAddValidator(payload)
	require.Error(t, err, "ValidateProposalPayloadAddValidator should return error on invalid payload")

	// Can't create a proposal with an invalid payload
	_, err = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		payload,
	)
	require.Error(t, err, "NewProposalAddValidator should prevent invalid payload")
}
