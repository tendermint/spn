package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/claim module sentinel errors
var (
	ErrMissionNotFound        = sdkerrors.Register(ModuleName, 2, "mission not found")
	ErrClaimRecordNotFound    = sdkerrors.Register(ModuleName, 3, "claim record not found")
	ErrMissionCompleted       = sdkerrors.Register(ModuleName, 4, "mission already completed")
	ErrAirdropSupplyNotFound  = sdkerrors.Register(ModuleName, 5, "airdrop supply not found")
	ErrInitialClaimNotFound   = sdkerrors.Register(ModuleName, 6, "initial claim information not found")
	ErrInitialClaimNotEnabled = sdkerrors.Register(ModuleName, 7, "initial claim not enabled")
	ErrMissionCompleteFailure = sdkerrors.Register(ModuleName, 8, "mission failed to complete")
)
