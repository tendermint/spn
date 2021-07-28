package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/profile module sentinel errors
var (
	ErrCoordAlreadyExist    = sdkerrors.Register(ModuleName, 1, "coordinator address already exist")
	ErrCoordAddressNotFound = sdkerrors.Register(ModuleName, 2, "coordinator address not found")
	ErrValidatorNotFound    = sdkerrors.Register(ModuleName, 4, "validator address not found")

	// this line is used by starport scaffolding # ibc/errors
)
