package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/launch module sentinel errors
var (
	ErrChainNotFound         = sdkerrors.Register(ModuleName, 1, "chain not found")
	ErrInvalidChainName      = sdkerrors.Register(ModuleName, 2, "the chain name is invalid")
	ErrInvalidChainID        = sdkerrors.Register(ModuleName, 3, "the chain id is invalid")
	ErrInvalidInitialGenesis = sdkerrors.Register(ModuleName, 4, "the initial genesis is invalid")
	ErrCodecNotPacked        = sdkerrors.Register(ModuleName, 5, "codec value couldn't be packed")
	ErrTriggeredLaunch       = sdkerrors.Register(ModuleName, 6, "launch is triggered for the chain")
	ErrNoAddressPermission   = sdkerrors.Register(ModuleName, 7, "you must be the coordinator or address owner to perform this action")
	ErrInvalidConsPubKey = sdkerrors.Register(ModuleName, 20, "the consensus public key is invalid")
	ErrInvalidGenTx = sdkerrors.Register(ModuleName, 21, "the gentx is invalid")
	ErrInvalidSelfDelegation = sdkerrors.Register(ModuleName, 22, "the self delegation is invalid")
	ErrInvalidPeer = sdkerrors.Register(ModuleName, 23, "the peer is invalid")
)
