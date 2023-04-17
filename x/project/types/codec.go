package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateProject{}, "project/CreateProject", nil)
	cdc.RegisterConcrete(&MsgEditProject{}, "project/EditProject", nil)
	cdc.RegisterConcrete(&MsgUpdateTotalSupply{}, "project/UpdateTotalSupply", nil)
	cdc.RegisterConcrete(&MsgUpdateSpecialAllocations{}, "project/UpdateSpecialAllocations", nil)
	cdc.RegisterConcrete(&MsgInitializeMainnet{}, "project/InitializeMainnet", nil)
	cdc.RegisterConcrete(&MsgMintVouchers{}, "project/MintVouchers", nil)
	cdc.RegisterConcrete(&MsgBurnVouchers{}, "project/BurnVouchers", nil)
	cdc.RegisterConcrete(&MsgRedeemVouchers{}, "project/RedeemVouchers", nil)
	cdc.RegisterConcrete(&MsgUnredeemVouchers{}, "project/UnredeemVouchers", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateProject{},
		&MsgEditProject{},
		&MsgUpdateTotalSupply{},
		&MsgUpdateSpecialAllocations{},
		&MsgInitializeMainnet{},
		&MsgMintVouchers{},
		&MsgBurnVouchers{},
		&MsgRedeemVouchers{},
		&MsgUnredeemVouchers{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
