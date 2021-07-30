package types

import (
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// RequestContent defines the interface for a request content
type RequestContent interface{}

// UnpackRequestRemoveValidator returns the RequestRemoveValidator
// structure from the codec unpack
func (r Request) UnpackRequestRemoveValidator(cdc codec.AnyUnpacker) (*ContentRemoveValidator, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var removeValidator *ContentRemoveValidator
	return removeValidator, cdc.UnpackAny(r.Content, &removeValidator)
}
