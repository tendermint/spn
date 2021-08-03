package types

import (
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// RequestContent defines the interface for a request content
type RequestContent interface{}

// UnpackValidatorRemoval returns the ValidatorRemoval structure from the codec unpack
func (r Request) UnpackValidatorRemoval(cdc codec.AnyUnpacker) (*ValidatorRemoval, error) {
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
		return nil, fmt.Errorf("not a validatorRemoval request")
	}
	return removeValidator, nil
}
