package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/profile module sentinel errors
var (
	ErrCoordAlreadyExist         = sdkerrors.Register(ModuleName, 2, "coordinator address already exist")
	ErrCoordAddressNotFound      = sdkerrors.Register(ModuleName, 3, "coordinator address not found")
	ErrCoordInvalid              = sdkerrors.Register(ModuleName, 4, "invalid coordinator")
	ErrEmptyDescription          = sdkerrors.Register(ModuleName, 5, "you must provide at least one description parameter")
	ErrInvalidValidatorSignature = sdkerrors.Register(ModuleName, 7, "invalid validator signature")
	ErrInvalidValidatorKey       = sdkerrors.Register(ModuleName, 8, "invalid validator key")
	ErrInvalidValidatorNonce     = sdkerrors.Register(ModuleName, 9, "invalid validator nonce")
	ErrInvalidValidatorChainID   = sdkerrors.Register(ModuleName, 10, "invalid validator chain id")
	ErrCoordInactive             = sdkerrors.Register(ModuleName, 11, "inactive coordinator")
)
