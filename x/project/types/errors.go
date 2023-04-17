package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/project module sentinel errors
var (
	ErrInvalidTotalSupply        = sdkerrors.Register(ModuleName, 2, "invalid total supply")
	ErrProjectNotFound           = sdkerrors.Register(ModuleName, 3, "project not found")
	ErrMainnetInitialized        = sdkerrors.Register(ModuleName, 4, "mainnet initialized")
	ErrInvalidShares             = sdkerrors.Register(ModuleName, 5, "invalid shares")
	ErrTotalSharesLimit          = sdkerrors.Register(ModuleName, 6, "allocated shares greater than total shares")
	ErrAccountNotFound           = sdkerrors.Register(ModuleName, 7, "account not found")
	ErrSharesDecrease            = sdkerrors.Register(ModuleName, 8, "shares can't be decreased")
	ErrVouchersMinting           = sdkerrors.Register(ModuleName, 9, "vouchers can't be minted")
	ErrInvalidVouchers           = sdkerrors.Register(ModuleName, 10, "invalid vouchers")
	ErrNoMatchVouchers           = sdkerrors.Register(ModuleName, 11, "vouchers don't match to project")
	ErrInsufficientVouchers      = sdkerrors.Register(ModuleName, 12, "account with insufficient vouchers")
	ErrInvalidProjectName        = sdkerrors.Register(ModuleName, 13, "invalid project name")
	ErrInvalidSupplyRange        = sdkerrors.Register(ModuleName, 14, "invalid total supply range")
	ErrInvalidMetadataLength     = sdkerrors.Register(ModuleName, 15, "metadata field too long")
	ErrMainnetLaunchTriggered    = sdkerrors.Register(ModuleName, 16, "mainnet launch already triggered")
	ErrInvalidSpecialAllocations = sdkerrors.Register(ModuleName, 17, "invalid special allocations")
	ErrInvalidMainnetInfo        = sdkerrors.Register(ModuleName, 18, "invalid mainnet info")
	ErrCannotUpdateProject       = sdkerrors.Register(ModuleName, 19, "cannot update project")
	ErrInvalidVoucherAddress     = sdkerrors.Register(ModuleName, 20, "invalid address for voucher operation")
	ErrFundCommunityPool         = sdkerrors.Register(ModuleName, 21, "unable to fund community pool")
)
