package types

import (
	"fmt"
	
	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// RequestContent defines the interface for a request content
type RequestContent interface{}


// UnpackRequestRemoveAccount returns the RequestRemoveAccount
// structure from the codec unpack
func (r Request) UnpackRequestRemoveAccount(cdc codec.AnyUnpacker) (*ContentRemoveAccount, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var removeValidator *ContentRemoveAccount
	return removeValidator, cdc.UnpackAny(r.Content, &removeValidator)
}