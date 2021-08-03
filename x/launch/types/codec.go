package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateChain{}, "launch/CreateChain", nil)

	cdc.RegisterConcrete(&MsgRequestRemoveValidator{}, "launch/RequestRemoveValidator", nil)

	cdc.RegisterInterface((*RequestContent)(nil), nil)
	cdc.RegisterConcrete(&GenesisAccount{}, "spn/launch/GenesisAccount", nil)
	cdc.RegisterConcrete(&ValidatorRemoval{}, "spn/launch/ValidatorRemoval", nil)

	cdc.RegisterInterface((*VestingOptions)(nil), nil)
	cdc.RegisterConcrete(&DelayedVesting{}, "spn/launch/DelayedVesting", nil)

	cdc.RegisterInterface((*InitialGenesis)(nil), nil)
	cdc.RegisterConcrete(&DefaultInitialGenesis{}, "launch/DefaultInitialGenesis", nil)
	cdc.RegisterConcrete(&GenesisURL{}, "launch/GenesisURL", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChain{},
		&MsgRequestRemoveValidator{},
	)

	registry.RegisterInterface(
		"launch.RequestContent",
		(*RequestContent)(nil),
		&ValidatorRemoval{},
		&GenesisAccount{},
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
