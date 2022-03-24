package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/campaign module sentinel errors
var (
	ErrInvalidTotalSupply    = sdkerrors.Register(ModuleName, 2, "invalid total supply")
	ErrCampaignNotFound      = sdkerrors.Register(ModuleName, 3, "campaign not found")
	ErrMainnetInitialized    = sdkerrors.Register(ModuleName, 4, "mainnet initialized")
	ErrInvalidShares         = sdkerrors.Register(ModuleName, 5, "invalid shares")
	ErrTotalSharesLimit      = sdkerrors.Register(ModuleName, 6, "allocated shares greater than total shares")
	ErrAccountNotFound       = sdkerrors.Register(ModuleName, 7, "account not found")
	ErrSharesDecrease        = sdkerrors.Register(ModuleName, 8, "shares can't be decreased")
	ErrVouchersMinting       = sdkerrors.Register(ModuleName, 9, "vouchers can't be minted")
	ErrInvalidVouchers       = sdkerrors.Register(ModuleName, 10, "invalid vouchers")
	ErrNoMatchVouchers       = sdkerrors.Register(ModuleName, 11, "vouchers don't match to campaign")
	ErrInsufficientVouchers  = sdkerrors.Register(ModuleName, 12, "account with insufficient vouchers")
	ErrInvalidCampaignName   = sdkerrors.Register(ModuleName, 13, "invalid campaign name")
	ErrInvalidSupplyRange    = sdkerrors.Register(ModuleName, 14, "invalid total supply range")
	ErrInvalidMetadataLength = sdkerrors.Register(ModuleName, 15, "metadata field too long")
)
