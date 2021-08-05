package types

import "errors"

// InitialGenesis defines the interface for initial genesis types
type InitialGenesis interface{
	Validate() error
}

var _ InitialGenesis = &DefaultInitialGenesis{}

// Validate implements InitialGenesis
func (DefaultInitialGenesis) Validate() error {return nil}

var _ InitialGenesis = &GenesisURL{}

const HashLength = 64

// Validate implements InitialGenesis
func (g GenesisURL) Validate() error {
	if g.Url == "" {
		return errors.New("no url provided")
	}
	if len(g.Hash) != HashLength {
		return errors.New("hash must be sha256")
	}
	return nil
}