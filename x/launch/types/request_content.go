package types

import (
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// RequestContent defines the interface for a request content
type RequestContent interface{}

// UnpackVestedAccount returns the VestedAccount structure from the codec unpack
func (r Request) UnpackVestedAccount(cdc codec.AnyUnpacker) (*VestedAccount, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}
	result, ok := content.(*VestedAccount)
	if !ok {
		return nil, fmt.Errorf("not a vestedAccount request")
	}
	return result, nil
}
