package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
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
	ErrVerifiedClientIDsNotFound    = sdkerrors.Register(ModuleName, 13, "verified client IDs not found")
	ErrInvalidClientCreatorAddress  = sdkerrors.Register(ModuleName, 14, "invalid client creator address")
	ErrCannotCloseChannel           = sdkerrors.Register(ModuleName, 15, "user cannot close channel")
	ErrJSONUnmarshal                = sdkerrors.Register(ModuleName, 16, "failed to unmarshal JSON")
	ErrJSONMarshal                  = sdkerrors.Register(ModuleName, 17, "failed to marshal JSON")
	ErrUnrecognizedPacketType       = sdkerrors.Register(ModuleName, 18, "unrecognized packet type")
)
