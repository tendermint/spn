package types

import (
	"errors"
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// RequestContent defines the interface for a request content
type RequestContent interface{}

// UnpackAccountRemoval returns the AccountRemoval structure from the codec unpack
func (r Request) UnpackAccountRemoval(cdc codec.AnyUnpacker) (*AccountRemoval, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}
	result, ok := content.(*AccountRemoval)
	if !ok {
		return nil, errors.New("not a accountRemoval request")
	}
	return result, nil
}

// UnpackGenesisValidator returns the GenesisValidator structure from the codec unpack
func (r Request) UnpackGenesisValidator(cdc codec.AnyUnpacker) (*GenesisValidator, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}
	result, ok := content.(*GenesisValidator)
	if !ok {
		return nil, errors.New("not a genesisValidator request")
	}
	return result, nil
}
