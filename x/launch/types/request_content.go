package types

import codec "github.com/cosmos/cosmos-sdk/codec/types"

// RequestContent defines the interface for a request content
type RequestContent interface{}

// UnpackRequestRemoveValidator returns the RequestRemoveValidator
// structure from the codec unpack
func (r Request) UnpackRequestRemoveValidator(cdc codec.AnyUnpacker) (*ContentRemoveValidator, error) {
	var removeValidator *ContentRemoveValidator
	return removeValidator, cdc.UnpackAny(r.Content, &removeValidator)
}
