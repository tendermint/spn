package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/monitoringp module sentinel errors
var (
	ErrInvalidVersion                = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrInvalidClientState            = sdkerrors.Register(ModuleName, 4, "invalid client state")
	ErrInvalidConsensusState         = sdkerrors.Register(ModuleName, 5, "invalid consensus state")
	ErrClientCreationFailure         = sdkerrors.Register(ModuleName, 6, "failed to create an IBC client")
	ErrInvalidHandshake              = sdkerrors.Register(ModuleName, 7, "invalid handshake")
	ErrNoConsumerClient              = sdkerrors.Register(ModuleName, 8, "consumer IBC client doesn't exist")
	ErrConsumerConnectionEstablished = sdkerrors.Register(ModuleName, 9, "consumer connection already established")
	ErrInvalidClient                 = sdkerrors.Register(ModuleName, 10, "invalid IBC client")
	ErrJSONUnmarshal                 = sdkerrors.Register(ModuleName, 11, "failed to unmarshal JSON")
	ErrJSONMarshal                   = sdkerrors.Register(ModuleName, 12, "failed to marshal JSON")
	ErrNotImplemented                = sdkerrors.Register(ModuleName, 13, "not implemented")
	ErrUnrecognizedAckType           = sdkerrors.Register(ModuleName, 14, "unrecognized acknowledgement type")
	ErrUnrecognizedPacketType        = sdkerrors.Register(ModuleName, 15, "unrecognized packet type")
	ErrCannotCloseChannel            = sdkerrors.Register(ModuleName, 16, "user cannot close channel")
)
