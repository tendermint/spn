package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/identity module sentinel errors
var (
	ErrInvalidUsername = sdkerrors.Register(ModuleName, 0, "the username is invalid")
	ErrInvalidAddress  = sdkerrors.Register(ModuleName, 1, "the address is invalid")
)
