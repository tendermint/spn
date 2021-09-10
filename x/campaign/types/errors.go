package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/campaign module sentinel errors
var (
	ErrCampaignNotFound   = sdkerrors.Register(ModuleName, 3, "campaign not found")
	ErrMainnetInitialized = sdkerrors.Register(ModuleName, 4, "mainnet initialized")
	ErrInvalidShares = sdkerrors.Register(ModuleName, 5, "invalid shares")
	ErrNoDynamicShares = sdkerrors.Register(ModuleName, 6, "no dynamic shares")
)
