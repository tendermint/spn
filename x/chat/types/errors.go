package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/chat module sentinel errors
var (
	ErrInvalidChannel = sdkerrors.Register(ModuleName, 0, "channel is invalid")
	ErrInvalidMessage = sdkerrors.Register(ModuleName, 1, "message is invalid")
	ErrInvalidPoll    = sdkerrors.Register(ModuleName, 2, "poll is invalid")
	ErrInvalidVote    = sdkerrors.Register(ModuleName, 3, "vote is invalid")
	ErrInvalidUser    = sdkerrors.Register(ModuleName, 4, "user is invalid")
)
