package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
)
