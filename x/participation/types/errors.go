package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// DONTCOVER

// x/participation module sentinel errors
var (
	ErrAuctionNotFound           = sdkerrors.Register(ModuleName, 2, "auction not found")
	ErrInvalidBidder             = sdkerrors.Register(ModuleName, 3, "invalid bidder")
	ErrInvalidAllocationAmount   = sdkerrors.Register(ModuleName, 4, "invalid allocation amount")
	ErrCannotWithdrawAllocations = sdkerrors.Register(ModuleName, 5, "unable to withdraw allocations")
	ErrUsedAllocationsNotFound   = sdkerrors.Register(ModuleName, 6, "unable to find used allocations entry")
)
