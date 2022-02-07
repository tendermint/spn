package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidRewardPoolCoins = sdkerrors.Register(ModuleName, 2, "invalid coins for reward pool")
	ErrInvalidCoordinatorID   = sdkerrors.Register(ModuleName, 3, "invalid coordinator id for reward pool")
	ErrModuleWithoutBalance   = sdkerrors.Register(ModuleName, 4, "module with insufficient funds")
	ErrAddressWithoutBalance  = sdkerrors.Register(ModuleName, 5, "address with insufficient funds")
)
