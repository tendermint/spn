package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/genesis module sentinel errors
var (
	ErrInvalidChain                = sdkerrors.Register(ModuleName, 1, "invalid chain")
	ErrInvalidVote                 = sdkerrors.Register(ModuleName, 2, "invalid vote")
	ErrInvalidProposalChange       = sdkerrors.Register(ModuleName, 3, "invalid change proposal")
	ErrInvalidProposalAddValidator = sdkerrors.Register(ModuleName, 4, "invalid add validator porposal")
	ErrInvalidProposalAddAccount   = sdkerrors.Register(ModuleName, 5, "invalid add account proposal")
)
