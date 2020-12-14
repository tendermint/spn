package types

import "errors"

// ApplyProposal performs the modification to the launch information from a approved proposal
func (li *LaunchInformation) ApplyProposal(proposal Proposal) error {
	// Dispatch the proposal
	switch payload := proposal.Payload.(type) {
	case *Proposal_AddAccountPayload:
		li.Accounts = append(li.Accounts, payload.AddAccountPayload)
	case *Proposal_AddValidatorPayload:
		li.GenTxs = append(li.GenTxs, payload.AddValidatorPayload.GenTx)
		li.Peers = append(li.Peers, payload.AddValidatorPayload.Peer)
	default:
		return errors.New("invalid proposal")
	}

	return nil
}
