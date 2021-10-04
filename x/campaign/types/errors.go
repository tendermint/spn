package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/campaign module sentinel errors
var (
	ErrInvalidCampaignName  = sdkerrors.Register(ModuleName, 1, "invalid campaign name")
	ErrInvalidTotalSupply   = sdkerrors.Register(ModuleName, 2, "invalid total supply")
	ErrCampaignNotFound     = sdkerrors.Register(ModuleName, 3, "campaign not found")
	ErrMainnetInitialized   = sdkerrors.Register(ModuleName, 4, "mainnet initialized")
	ErrInvalidShares        = sdkerrors.Register(ModuleName, 5, "invalid shares")
	ErrNoDynamicShares      = sdkerrors.Register(ModuleName, 6, "no dynamic shares")
	ErrInvalidAccountShares = sdkerrors.Register(ModuleName, 7, "invalid account shares")
	ErrTotalSharesLimit     = sdkerrors.Register(ModuleName, 8, "allocated shares greater than total shares")
	ErrVouchersMinting      = sdkerrors.Register(ModuleName, 10, "vouchers can't be minted")
)
