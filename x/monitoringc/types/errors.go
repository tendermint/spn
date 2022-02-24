package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/monitoringc module sentinel errors
var (
	ErrInvalidVersion               = sdkerrors.Register(ModuleName, 2, "invalid version")
	ErrInvalidClientState           = sdkerrors.Register(ModuleName, 3, "invalid client state")
	ErrInvalidConsensusState        = sdkerrors.Register(ModuleName, 4, "invalid consensus state")
	ErrInvalidValidatorSet          = sdkerrors.Register(ModuleName, 5, "invalid validator set")
	ErrInvalidValidatorSetHash      = sdkerrors.Register(ModuleName, 6, "invalid validator set hash")
	ErrClientCreationFailure        = sdkerrors.Register(ModuleName, 7, "failed to create IBC client")
	ErrInvalidHandshake             = sdkerrors.Register(ModuleName, 8, "invalid handshake")
	ErrClientNotVerified            = sdkerrors.Register(ModuleName, 9, "ibc client not verified")
	ErrConnectionAlreadyEstablished = sdkerrors.Register(ModuleName, 10, "ibc connection already established")
	ErrInvalidUnbondingPeriod       = sdkerrors.Register(ModuleName, 11, "invalid unbonding period")
	ErrInvalidRevisionHeight        = sdkerrors.Register(ModuleName, 12, "invalid revision height")
)
