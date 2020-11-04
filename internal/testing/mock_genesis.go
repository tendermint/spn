package testing

import (
	"github.com/tendermint/spn/x/genesis/types"
	"math/rand"
	"time"
)

// MockProposalChangePayload mocks a valid payload
func MockProposalChangePayload() *types.ProposalChangePayload {
	return types.NewProposalChangePayload(
		MockRandomString(5)+"."+MockRandomString(5)+"."+MockRandomString(5),
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
	return types.NewProposalAddValidatorPayload(
		MockValAddress(),
		[]byte(MockRandomString(20)), // TODO: Generate a correct consensus pub key
		MockDescription(),
		MockCommissionRates(),
		MockCoin(),
		[]byte(MockRandomString(1000)),
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

// MockProposalVote mocks a vote for a genesis proposal
func MockProposalVote() *types.Vote {
	voteValue := types.Vote_REJECT

	if r := rand.Intn(10); r > 5 {
		voteValue = types.Vote_APPROVE
	}

	vote, _ := types.NewVote(
		int32(rand.Intn(10)),
		MockRandomString(10),
		time.Now(),
		voteValue,
	)
	return vote
}
