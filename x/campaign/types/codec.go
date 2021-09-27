package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgInitializeMainnet{}, "campaign/InitializeMainnet", nil)
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgUpdateTotalShares{}, "campaign/UpdateTotalShares", nil)
	cdc.RegisterConcrete(&MsgUpdateTotalSupply{}, "campaign/UpdateTotalSupply", nil)
	cdc.RegisterConcrete(&MsgCreateCampaign{}, "campaign/CreateCampaign", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgInitializeMainnet{},
	)
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateTotalShares{},
		&MsgUpdateTotalSupply{},
		&MsgCreateCampaign{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
