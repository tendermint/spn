package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/campaign module sentinel errors
var (
	ErrInvalidShares = sdkerrors.Register(ModuleName, 5, "invalid shares")
)
