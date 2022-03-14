package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// DONTCOVER

// x/participation module sentinel errors
var (
	ErrAuctionNotFound = sdkerrors.Register(ModuleName, 2, "auction not found")
	ErrInvalidBidder   = sdkerrors.Register(ModuleName, 3, "invalid bidder")
)
