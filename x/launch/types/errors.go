package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrChainNotFound         = sdkerrors.Register(ModuleName, 1, "chain not found")
	ErrInvalidChainName = sdkerrors.Register(ModuleName, 2, "the chain name is invalid")
	ErrLaunchTriggered = sdkerrors.Register(ModuleName, 10, "the chain launch is already triggered")
	ErrLaunchTimeTooLow = sdkerrors.Register(ModuleName, 11, "the remaining time is below authorized launch time")
	ErrLaunchTimeToohigh = sdkerrors.Register(ModuleName, 12, "the remaining time is above authorized launch time")
)
