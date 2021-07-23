package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrCoordinatorNotExist = sdkerrors.Register(ModuleName, 1, "the coordinator doesn't exist")
)
