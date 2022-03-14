package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

// DONTCOVER

// x/participation module sentinel errors
var (
	ErrInvalidAllocationAmount = sdkerrors.Register(ModuleName, 4, "invalid allocation amount")
)
