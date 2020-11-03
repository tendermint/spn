package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/genesis module sentinel errors
var (
	ErrInvalidChain = sdkerrors.Register(ModuleName, 1, "the chain is invalid")
	ErrInvalidVote  = sdkerrors.Register(ModuleName, 2, "the vote is invalid")
)
