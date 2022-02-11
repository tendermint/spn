package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidRewardPoolCoins = sdkerrors.Register(ModuleName, 2, "invalid coins for reward pool")
	ErrInvalidCoordinatorID   = sdkerrors.Register(ModuleName, 3, "invalid coordinator id for reward pool")
	ErrRewardPoolNotFound     = sdkerrors.Register(ModuleName, 4, "reward pool not found")
)
