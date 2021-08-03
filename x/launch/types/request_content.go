package types

import (
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// RequestContent defines the interface for a request content
type RequestContent interface{}

// UnpackRequestRemoveValidator returns the RequestRemoveValidator
// structure from the codec unpack
func (r Request) UnpackRequestRemoveValidator(cdc codec.AnyUnpacker) (*ValidatorRemoval, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}
	removeValidator, ok := content.(*ValidatorRemoval)
	if !ok {
		return nil, ErrFailedCodecCast
	}
	return removeValidator, nil
}
