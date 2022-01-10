package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/monitoringc module sentinel errors
var (
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 1, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 2, "invalid version")
	ErrInvalidClientState       = sdkerrors.Register(ModuleName, 3, "invalid client state")

)
