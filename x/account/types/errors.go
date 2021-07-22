package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/account module sentinel errors
var (
	ErrCoordAddressNotFound = sdkerrors.Register(ModuleName, 2, "coordinator address not found")
	// this line is used by starport scaffolding # ibc/errors
)
