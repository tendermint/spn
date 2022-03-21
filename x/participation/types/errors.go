package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// DONTCOVER

// x/participation module sentinel errors
var (
	ErrAuctionNotFound             = sdkerrors.Register(ModuleName, 2, "auction not found")
	ErrInvalidBidder               = sdkerrors.Register(ModuleName, 3, "invalid bidder")
	ErrInvalidAllocationAmount     = sdkerrors.Register(ModuleName, 4, "invalid allocation amount")
	ErrTierNotFound                = sdkerrors.Register(ModuleName, 5, "tier not found")
	ErrInsufficientAllocations     = sdkerrors.Register(ModuleName, 6, "insufficient allocations")
	ErrAlreadyParticipating        = sdkerrors.Register(ModuleName, 7, "address is already participating")
	ErrAllocationsLocked           = sdkerrors.Register(ModuleName, 8, "unable to withdraw allocations")
	ErrUsedAllocationsNotFound     = sdkerrors.Register(ModuleName, 9, "unable to find used allocations entry")
	ErrAllocationsAlreadyWithdrawn = sdkerrors.Register(ModuleName, 10, "used allocations already withdrawn")
)
