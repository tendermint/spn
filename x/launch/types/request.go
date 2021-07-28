package types

import (
	codec "github.com/cosmos/cosmos-sdk/codec/types"
)

// AnyFromRequest the proto any type for a Request
func AnyFromRequest() *codec.Any {
	defaultGenesis, err := codec.NewAnyWithValue(&DefaultInitialGenesis{})
	if err != nil {
		panic("DefaultInitialGenesis can't be used as initial genesis")
	}
	return defaultGenesis
}
