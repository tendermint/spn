package types

import (
	"errors"

	"github.com/tendermint/spn/pkg/chainid"
)

// Validate checks the chain has valid data
func (m Chain) Validate() error {
	if _, _, err := chainid.ParseGenesisChainID(m.GenesisChainID); err != nil {
		return err
	}

	// A chain that is a mainnet is always associated to a campaign
	if m.IsMainnet && !m.HasCampaign {
		return errors.New("chain is a mainnet but not associated to a campaign")
	}

	return nil
}
