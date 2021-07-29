package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrInvalidChainName = sdkerrors.Register(ModuleName, 1, "the chain name is invalid")
	ErrChainIdNotFound  = sdkerrors.Register(ModuleName, 2, "chain id not found")
	// this line is used by starport scaffolding # ibc/errors
)
