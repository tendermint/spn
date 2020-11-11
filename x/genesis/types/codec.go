package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgChainCreate{}, "genesis/MsgChainCreate", nil)
	cdc.RegisterConcrete(MsgProposalAddAccount{}, "genesis/MsgProposalAddAccount", nil)
	cdc.RegisterConcrete(MsgProposalAddValidator{}, "genesis/MsgProposalAddValidator", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgProposalAddAccount{},
		&MsgProposalAddValidator{},
		&MsgChainCreate{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
