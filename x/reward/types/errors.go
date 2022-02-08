package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrRewardPoolNotFound = sdkerrors.Register(ModuleName, 4, "reward pool not found")
)
