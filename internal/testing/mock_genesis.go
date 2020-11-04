package testing

import (
	"github.com/tendermint/spn/x/genesis/types"
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
