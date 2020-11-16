package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"
)

type GenesisFile []byte

// NewGenesisFile returns a new genesis file from the bytes of the genesis
func NewGenesisFile(genesisBytes []byte) GenesisFile {
	return genesisBytes
}

// GetGenesisDoc returns the genesis in the tendermint/types.GenesisDoc format
func (g GenesisFile) GetGenesisDoc() (genesisDoc tmtypes.GenesisDoc, err error) {
	err = tmjson.Unmarshal(g, &genesisDoc)
	return genesisDoc, err
}

// SetChainID set the chain ID for the genesis file
func (g *GenesisFile) SetChainID(chainID string) error {
	// Unmarshal bytes
	var genesisObject tmtypes.GenesisDoc
	if err := tmjson.Unmarshal(*g, &genesisObject); err != nil {
		return err
	}

	genesisObject.ChainID = chainID

	// Marshal
	bz, err := tmjson.Marshal(genesisObject)
	if err != nil {
		return err
	}
	*g = bz

	return nil
}

// ValidateAndComplete checks that all necessary fields are present and fills in defaults for optional fields left empty
func (g *GenesisFile) ValidateAndComplete() error {
	// Unmarshal bytes
	var genesisObject tmtypes.GenesisDoc
	if err := tmjson.Unmarshal(*g, &genesisObject); err != nil {
		return sdkerrors.Wrap(ErrInvalidChain, err.Error())
	}

	// Validate and complete
	if err := genesisObject.ValidateAndComplete(); err != nil {
		return sdkerrors.Wrap(ErrInvalidChain, err.Error())
	}

	// Marshal
	bz, err := tmjson.Marshal(genesisObject)
	if err != nil {
		return err
	}
	*g = bz

	return nil
}