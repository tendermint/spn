package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateChain{}, "launch/CreateChain", nil)
	cdc.RegisterConcrete(&MsgEditChain{}, "launch/EditChain", nil)
	cdc.RegisterConcrete(&MsgRequestAddAccount{}, "launch/RequestAddAccount", nil)
	cdc.RegisterConcrete(&MsgRequestAddVestingAccount{}, "launch/RequestAddVestingdAccount", nil)
	cdc.RegisterConcrete(&MsgRequestRemoveAccount{}, "launch/RequestRemoveAccount", nil)
	cdc.RegisterConcrete(&MsgRequestAddValidator{}, "launch/RequestAddValidator", nil)
	cdc.RegisterConcrete(&MsgRequestRemoveValidator{}, "launch/RequestRemoveValidator", nil)
	cdc.RegisterConcrete(&MsgSettleRequest{}, "launch/SettleRequest", nil)
	cdc.RegisterConcrete(&MsgTriggerLaunch{}, "launch/TriggerLaunch", nil)
	cdc.RegisterConcrete(&MsgRevertLaunch{}, "launch/RevertLaunch", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChain{},
		&MsgEditChain{},
		&MsgRequestAddAccount{},
		&MsgRequestAddVestingAccount{},
		&MsgRequestRemoveAccount{},
		&MsgRequestAddValidator{},
		&MsgRequestRemoveValidator{},
		&MsgSettleRequest{},
		&MsgTriggerLaunch{},
		&MsgRevertLaunch{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
