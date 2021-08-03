package types

import (
	"fmt"
	
	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// RequestContent defines the interface for a request content
type RequestContent interface{}

// UnpackGenesisAccount returns the GenesisAccount structure from the codec unpack
func (r Request) UnpackGenesisAccount(cdc codec.AnyUnpacker) (*GenesisAccount, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}
	removeValidator, ok := content.(*GenesisAccount)
	if !ok {
		return nil, ErrFailedCodecCast
	}
	return removeValidator, nil
}
