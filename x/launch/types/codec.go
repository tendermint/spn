package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgEditChain{}, "launch/EditChain", nil)

	cdc.RegisterConcrete(&MsgRequestAddVestedAccount{}, "launch/RequestAddVestedAccount", nil)

	cdc.RegisterConcrete(&MsgCreateChain{}, "launch/CreateChain", nil)

	cdc.RegisterInterface((*RequestContent)(nil), nil)
	cdc.RegisterConcrete(&GenesisAccount{}, "spn/launch/GenesisAccount", nil)
	cdc.RegisterConcrete(&VestedAccount{}, "spn/launch/VestedAccount", nil)

	cdc.RegisterInterface((*VestingOptions)(nil), nil)
	cdc.RegisterConcrete(&DelayedVesting{}, "spn/launch/DelayedVesting", nil)

	cdc.RegisterInterface((*InitialGenesis)(nil), nil)
	cdc.RegisterConcrete(&DefaultInitialGenesis{}, "launch/DefaultInitialGenesis", nil)
	cdc.RegisterConcrete(&GenesisURL{}, "launch/GenesisURL", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgEditChain{},
		&MsgCreateChain{},
		&MsgRequestAddVestedAccount{},
	)

	registry.RegisterInterface(
		"launch.RequestContent",
		(*RequestContent)(nil),
		&GenesisAccount{},
		&VestedAccount{},
	)

	registry.RegisterInterface(
		"launch.VestingOptions",
		(*VestingOptions)(nil),
		&DelayedVesting{},
	)

	registry.RegisterInterface(
		"launch.InitialGenesis",
		(*InitialGenesis)(nil),
		&DefaultInitialGenesis{},
		&GenesisURL{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
