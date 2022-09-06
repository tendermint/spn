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
	cdc.RegisterConcrete(&MsgSendRequest{}, "launch/SendRequest", nil)
	cdc.RegisterConcrete(&MsgSettleRequest{}, "launch/SettleRequest", nil)
	cdc.RegisterConcrete(&MsgTriggerLaunch{}, "launch/TriggerLaunch", nil)
	cdc.RegisterConcrete(&MsgRevertLaunch{}, "launch/RevertLaunch", nil)
	cdc.RegisterConcrete(&MsgUpdateLaunchInformation{}, "launch/UpdateLaunchInformation", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChain{},
		&MsgEditChain{},
		&MsgUpdateLaunchInformation{},
		&MsgSendRequest{},
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
