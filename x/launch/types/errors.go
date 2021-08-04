package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrChainNotFound     = sdkerrors.Register(ModuleName, 1, "chain not found")
	ErrInvalidChainName = sdkerrors.Register(ModuleName, 2, "the chain name is invalid")
	ErrLaunchNotTriggered   = sdkerrors.Register(ModuleName, 13, "the chain launch has not been triggered")
	ErrRevertDelayNotReached   = sdkerrors.Register(ModuleName, 14, "the revert delay has not been reached")

)
