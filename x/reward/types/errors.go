package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

var (
	ErrInvalidRewardPoolCoins = sdkerrors.Register(ModuleName, 2, "invalid coins for reward pool")
	ErrInvalidCoordinatorID   = sdkerrors.Register(ModuleName, 3, "invalid coordinator id for reward pool")
	ErrRewardPoolNotFound     = sdkerrors.Register(ModuleName, 4, "reward pool not found")
	ErrRewardPoolClosed       = sdkerrors.Register(ModuleName, 5, "reward pool is closed")
	ErrInvalidSignatureCounts = sdkerrors.Register(ModuleName, 6, "invalid signature counts")
	ErrInvalidLastBlockHeight = sdkerrors.Register(ModuleName, 7, "invalid last block height")
	ErrInvalidRewardHeight    = sdkerrors.Register(ModuleName, 8, "invalid reward height")
	ErrInvalidProviderAddress = sdkerrors.Register(ModuleName, 9, "invalid provider address")
	ErrInsufficientFunds      = sdkerrors.Register(ModuleName, 10, "insufficient funds")
)
