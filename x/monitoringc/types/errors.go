package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/monitoringc module sentinel errors
var (
	ErrInvalidPacketTimeout    = sdkerrors.Register(ModuleName, 2, "invalid packet timeout")
	ErrInvalidVersion          = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrInvalidClientState      = sdkerrors.Register(ModuleName, 4, "invalid client state")
	ErrInvalidConsensusState   = sdkerrors.Register(ModuleName, 5, "invalid consensus state")
	ErrInvalidValidatorSet     = sdkerrors.Register(ModuleName, 6, "invalid validator set")
	ErrInvalidValidatorSetHash = sdkerrors.Register(ModuleName, 7, "invalid validator set hash")
	ErrClientCreationFailure   = sdkerrors.Register(ModuleName, 8, "failed to create IBC client")
	ErrInvalidHandshake        = sdkerrors.Register(ModuleName, 9, "invalid handshake")
)
