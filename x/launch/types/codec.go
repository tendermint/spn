package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"

	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterInterface((*InitialGenesis)(nil), nil)
	cdc.RegisterConcrete(&DefaultInitialGenesis{}, "spn/launch/DefaultInitialGenesis", nil)
	cdc.RegisterConcrete(&GenesisURL{}, "spn/launch/GenesisURL", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)

	registry.RegisterInterface(
		"tendermint.spn.launch.InitialGenesis",
		(*InitialGenesis)(nil),
		&DefaultInitialGenesis{},
		&GenesisURL{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
