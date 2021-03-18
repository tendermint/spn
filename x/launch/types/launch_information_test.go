package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/launch/types"
	"testing"
)

func TestApplyProposal(t *testing.T) {
	var li types.LaunchInformation

	// Can apply an add account proposal
	addAccoundPayload := spnmocks.MockProposalAddAccountPayload()
	proposal, _ := types.NewProposalAddAccount(
		spnmocks.MockProposalInformation(),
		addAccoundPayload,
	)

	err := li.ApplyProposal(*proposal)
	require.NoError(t, err)
	require.Equal(t, 1, len(li.Accounts))
	require.Equal(t, 0, len(li.GenTxs))
	require.Equal(t, 0, len(li.Peers))
	require.Equal(t, addAccoundPayload, li.Accounts[0])

	// Can apply an add validator proposal
	addValidatorPayload := spnmocks.MockProposalAddValidatorPayload()
	proposal, _ = types.NewProposalAddValidator(
		spnmocks.MockProposalInformation(),
		addValidatorPayload,
	)

	err = li.ApplyProposal(*proposal)
	require.NoError(t, err)
	require.Equal(t, 1, len(li.Accounts))
	require.Equal(t, 1, len(li.GenTxs))
	require.Equal(t, 1, len(li.Peers))
	require.Equal(t, addValidatorPayload.Peer, li.Peers[0])
	require.Equal(t, addValidatorPayload.GenTx, li.GenTxs[0])

	// Cannot apply an invalid proposal
	proposal.Payload = nil
	err = li.ApplyProposal(*proposal)
	require.Error(t, err)
}