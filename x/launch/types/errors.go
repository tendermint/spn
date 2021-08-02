package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrChainNotFound    = sdkerrors.Register(ModuleName, 1, "chain not found")
	ErrInvalidChainName = sdkerrors.Register(ModuleName, 2, "the chain name is invalid")
	// this line is used by starport scaffolding # ibc/errors
)
