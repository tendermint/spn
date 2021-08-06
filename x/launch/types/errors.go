package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrChainNotFound         = sdkerrors.Register(ModuleName, 1, "chain not found")
	ErrInvalidChainName      = sdkerrors.Register(ModuleName, 2, "the chain name is invalid")
	ErrInvalidChainID        = sdkerrors.Register(ModuleName, 3, "the chain id is invalid")
	ErrInvalidInitialGenesis = sdkerrors.Register(ModuleName, 4, "the initial genesis is invalid")
	ErrCodecNotPacked        = sdkerrors.Register(ModuleName, 5, "codec value couldn't be packed")
	ErrTriggeredLaunch       = sdkerrors.Register(ModuleName, 6, "launch is triggered for the chain")
	ErrNoAddressPermission   = sdkerrors.Register(ModuleName, 7, "you must be the coordinator or address owner to perform this action")
	ErrNotTriggeredLaunch    = sdkerrors.Register(ModuleName, 13, "the chain launch has not been triggered")
	ErrRevertDelayNotReached = sdkerrors.Register(ModuleName, 14, "the revert delay has not been reached")
	// this line is used by starport scaffolding # ibc/errors
)
