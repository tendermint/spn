package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/campaign module sentinel errors
var (
	ErrInvalidCampaignName = sdkerrors.Register(ModuleName, 1, "invalid campaign name")
	ErrInvalidTotalSupply  = sdkerrors.Register(ModuleName, 2, "invalid total supply")
	ErrCampaignNotFound   = sdkerrors.Register(ModuleName, 3, "campaign not found")
	ErrMainnetInitialized = sdkerrors.Register(ModuleName, 4, "mainnet initialized")
)
