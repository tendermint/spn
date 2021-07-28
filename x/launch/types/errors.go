package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrSample          = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrChainIdNotFound = sdkerrors.Register(ModuleName, 1, "chain id not found")
	// this line is used by starport scaffolding # ibc/errors
)
