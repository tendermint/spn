package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgRequestAddValidator{}, "launch/RequestAddValidator", nil)
	cdc.RegisterConcrete(&MsgCreateChain{}, "launch/CreateChain", nil)
	cdc.RegisterConcrete(&MsgEditChain{}, "launch/EditChain", nil)
	cdc.RegisterConcrete(&MsgRevertLaunch{}, "launch/RevertLaunch", nil)
	cdc.RegisterConcrete(&MsgTriggerLaunch{}, "launch/TriggerLaunch", nil)
	cdc.RegisterConcrete(&MsgSettleRequest{}, "launch/SettleRequest", nil)
	cdc.RegisterConcrete(&MsgRequestAddAccount{}, "launch/RequestAddAccount", nil)
	cdc.RegisterConcrete(&MsgRequestRemoveAccount{}, "launch/RequestRemoveAccount", nil)
	cdc.RegisterConcrete(&MsgRequestRemoveValidator{}, "launch/RequestRemoveValidator", nil)
	cdc.RegisterConcrete(&MsgRequestAddVestingAccount{}, "launch/RequestAddVestedAccount", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRequestAddValidator{},
		&MsgEditChain{},
		&MsgCreateChain{},
		&MsgEditChain{},
		&MsgRevertLaunch{},
		&MsgTriggerLaunch{},
		&MsgSettleRequest{},
		&MsgRequestAddAccount{},
		&MsgRequestRemoveAccount{},
		&MsgRequestRemoveValidator{},
		&MsgRequestAddVestingAccount{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
