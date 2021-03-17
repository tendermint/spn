package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrInvalidChain                = sdkerrors.Register(ModuleName, 1, "invalid chain")
	ErrInvalidVote                 = sdkerrors.Register(ModuleName, 2, "invalid vote")
	ErrInvalidProposal             = sdkerrors.Register(ModuleName, 3, "invalid proposal")
	ErrInvalidProposalChange       = sdkerrors.Register(ModuleName, 4, "invalid change proposal")
	ErrInvalidProposalAddValidator = sdkerrors.Register(ModuleName, 5, "invalid add validator proposal")
	ErrInvalidProposalAddAccount   = sdkerrors.Register(ModuleName, 6, "invalid add account proposal")
)
