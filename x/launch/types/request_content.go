package types

import (
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
	removeValidator, ok := content.(*AccountRemoval)
	if !ok {
		return nil, ErrFailedCodecCast
	}
	return removeValidator, nil
}
