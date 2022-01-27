package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrSample = sdkerrors.Register(ModuleName, 2, "sample error")
)
