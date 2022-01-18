package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/monitoringp module sentinel errors
var (
	ErrInvalidPacketTimeout  = sdkerrors.Register(ModuleName, 2, "invalid packet timeout")
	ErrInvalidVersion        = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrInvalidClientState    = sdkerrors.Register(ModuleName, 4, "invalid client state")
	ErrInvalidConsensusState    = sdkerrors.Register(ModuleName, 5, "invalid consensus state")
	ErrClientCreationFailure = sdkerrors.Register(ModuleName, 6, "failed to create IBC client")
)
