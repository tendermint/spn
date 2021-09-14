package types

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewChain creates a new chain information object
func NewChain(
	chainID string,
	creator string,
	sourceURL string,
	sourceHash string,
	createdAt time.Time,
	genesisURL string,
	genesisHash string,
) (*Chain, error) {
	var chain Chain

	// Check chainID validity
	if !checkChainID(chainID) {
		return nil, sdkerrors.Wrap(ErrInvalidChain, fmt.Sprintf("invalid chain ID %v", chainID))
	}
	chain.ChainID = chainID
	chain.Creator = creator
	chain.Peers = []string{}
	chain.SourceURL = sourceURL
	chain.SourceHash = sourceHash
	chain.CreatedAt = createdAt.Unix()

	// Check if initial genesis is the default genesis or genesis from a URL
	if genesisURL == "" {
		// Default genesis
		chain.InitialGenesis = NewInitialGenesisDefault()
	} else {
		// Genesis from url
		genesisURL, err := NewGenesisURL(genesisURL, genesisHash)
		if err != nil {
			return nil, sdkerrors.Wrap(ErrInvalidChain, err.Error())
		}

		chain.InitialGenesis = NewInitialGenesisURL(genesisURL)
	}

	return &chain, nil
}

// AppendPeer appends a new peer in the peer list of the chain
func (c *Chain) AppendPeer(peer string) {
	c.Peers = append(c.Peers, peer)
}

// checkChainID performs stateless verification of the chainID
// The chainID must contain only alphanumeric character or hyphen
func checkChainID(chainID string) bool {
	// Check the chainID is not empty
	if len(chainID) == 0 {
		return false
	}

	// Iterate characters
	for _, c := range chainID {
		if !isChainAuthorizedChar(c) {
			return false
		}
	}
	return true
}

// isChainAuthorizedChar checks to ensure that all letters in the ChainID are valid
func isChainAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '-'
}

// MarshalChain encodes chains for the store
func MarshalChain(cdc codec.BinaryCodec, chain Chain) []byte {
	return cdc.MustMarshal(&chain)
}

// UnmarshalChain decodes chains from the store
func UnmarshalChain(cdc codec.BinaryCodec, value []byte) Chain {
	var chain Chain
	cdc.MustUnmarshal(value, &chain)
	return chain
}
