package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrChainNotFound         = sdkerrors.Register(ModuleName, 1, "chain not found")
	ErrInvalidGenesisChainID = sdkerrors.Register(ModuleName, 2, "the genesis chain id is invalid")
	ErrInvalidInitialGenesis = sdkerrors.Register(ModuleName, 3, "the initial genesis is invalid")
	ErrCodecNotPacked        = sdkerrors.Register(ModuleName, 4, "codec value couldn't be packed")
	ErrTriggeredLaunch       = sdkerrors.Register(ModuleName, 5, "launch is triggered for the chain")
	ErrNoAddressPermission   = sdkerrors.Register(ModuleName, 6, "you must be the coordinator or address owner to perform this action")
	ErrInvalidCoins          = sdkerrors.Register(ModuleName, 7, "the coin list is invalid")
	ErrInvalidVestingOption  = sdkerrors.Register(ModuleName, 8, "invalid vesting option")
	ErrInvalidTimestamp      = sdkerrors.Register(ModuleName, 9, "timestamp must be greater than zero")
	ErrLaunchTimeTooLow      = sdkerrors.Register(ModuleName, 10, "the remaining time is below authorized launch time")
	ErrLaunchTimeTooHigh     = sdkerrors.Register(ModuleName, 11, "the remaining time is above authorized launch time")
	ErrNotTriggeredLaunch    = sdkerrors.Register(ModuleName, 12, "the chain launch has not been triggered")
	ErrRevertDelayNotReached = sdkerrors.Register(ModuleName, 13, "the revert delay has not been reached")
	ErrRequestNotFound       = sdkerrors.Register(ModuleName, 14, "request not found")
	ErrInvalidConsPubKey     = sdkerrors.Register(ModuleName, 15, "the consensus public key is invalid")
	ErrInvalidGenTx          = sdkerrors.Register(ModuleName, 16, "the gentx is invalid")
	ErrInvalidSelfDelegation = sdkerrors.Register(ModuleName, 17, "the self delegation is invalid")
	ErrInvalidPeer           = sdkerrors.Register(ModuleName, 18, "the peer is invalid")
	ErrAccountAlreadyExist   = sdkerrors.Register(ModuleName, 19, "account already exists")
	ErrAccountNotFound       = sdkerrors.Register(ModuleName, 20, "account not found")
	ErrValidatorAlreadyExist = sdkerrors.Register(ModuleName, 21, "validator already exists")
	ErrValidatorNotFound     = sdkerrors.Register(ModuleName, 22, "validator not found")
	ErrChainInactive         = sdkerrors.Register(ModuleName, 23, "the chain is inactive")
)
