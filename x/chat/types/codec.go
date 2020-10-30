package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgCreateChannel{}, "chat/MsgCreateChannel", nil)
	cdc.RegisterConcrete(MsgSendMessage{}, "chat/MsgSendMessage", nil)
	cdc.RegisterConcrete(MsgVotePoll{}, "chat/MsgVotePoll", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateChannel{},
		&MsgSendMessage{},
		&MsgVotePoll{},
	)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)
