package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/profile module sentinel errors
var (
	ErrCoordAlreadyExist    = sdkerrors.Register(ModuleName, 1, "coordinator address already exist")
	ErrCoordAddressNotFound = sdkerrors.Register(ModuleName, 2, "coordinator address not found")
	ErrCoordInvalid         = sdkerrors.Register(ModuleName, 3, "invalid coordinator")
	ErrEmptyDescription     = sdkerrors.Register(ModuleName, 4, "you must provide at least one description parameter")
	ErrValidatorNotFound    = sdkerrors.Register(ModuleName, 5, "validator address not found")
)
