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
)
