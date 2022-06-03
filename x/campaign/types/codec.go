package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCampaign{}, "campaign/CreateCampaign", nil)
	cdc.RegisterConcrete(&MsgEditCampaign{}, "campaign/EditCampaign", nil)
	cdc.RegisterConcrete(&MsgUpdateTotalSupply{}, "campaign/UpdateTotalSupply", nil)
	cdc.RegisterConcrete(&MsgUpdateSpecialAllocations{}, "campaign/UpdateSpecialAllocations", nil)
	cdc.RegisterConcrete(&MsgInitializeMainnet{}, "campaign/InitializeMainnet", nil)
	cdc.RegisterConcrete(&MsgAddShares{}, "campaign/AddShares", nil)
	cdc.RegisterConcrete(&MsgAddVestingOptions{}, "campaign/AddVestingOptions", nil)
	cdc.RegisterConcrete(&MsgMintVouchers{}, "campaign/MintVouchers", nil)
	cdc.RegisterConcrete(&MsgBurnVouchers{}, "campaign/BurnVouchers", nil)
	cdc.RegisterConcrete(&MsgRedeemVouchers{}, "campaign/RedeemVouchers", nil)
	cdc.RegisterConcrete(&MsgUnredeemVouchers{}, "campaign/UnredeemVouchers", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCampaign{},
		&MsgEditCampaign{},
		&MsgUpdateTotalSupply{},
		&MsgUpdateSpecialAllocations{},
		&MsgInitializeMainnet{},
		&MsgAddShares{},
		&MsgAddVestingOptions{},
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
