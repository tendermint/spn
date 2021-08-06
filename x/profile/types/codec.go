package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgUpdateValidatorDescription{}, "profile/UpdateValidatorDescription", nil)
	cdc.RegisterConcrete(&MsgDeleteValidator{}, "profile/DeleteValidator", nil)
	cdc.RegisterConcrete(&MsgCreateCoordinator{}, "profile/CreateCoordinator", nil)
	cdc.RegisterConcrete(&MsgDeleteCoordinator{}, "profile/DeleteCoordinator", nil)
	cdc.RegisterConcrete(&MsgUpdateCoordinatorAddress{}, "profile/UpdateCoordinatorAddress", nil)
	cdc.RegisterConcrete(&MsgUpdateCoordinatorDescription{}, "profile/UpdateCoordinatorDescription", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateValidatorDescription{},
		&MsgDeleteValidator{},
		&MsgCreateCoordinator{},
		&MsgDeleteCoordinator{},
		&MsgUpdateCoordinatorAddress{},
		&MsgUpdateCoordinatorDescription{},
	)
	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
