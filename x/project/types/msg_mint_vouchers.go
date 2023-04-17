package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	profile "github.com/tendermint/spn/x/profile/types"
)

const TypeMsgMintVouchers = "mint_vouchers"

var _ sdk.Msg = &MsgMintVouchers{}

func NewMsgMintVouchers(coordinator string, projectID uint64, shares Shares) *MsgMintVouchers {
	return &MsgMintVouchers{
		Coordinator: coordinator,
		ProjectID:   projectID,
		Shares:      shares,
	}
}

func (msg *MsgMintVouchers) Route() string {
	return RouterKey
}

func (msg *MsgMintVouchers) Type() string {
	return TypeMsgMintVouchers
}

func (msg *MsgMintVouchers) GetSigners() []sdk.AccAddress {
	coordinator, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{coordinator}
}

func (msg *MsgMintVouchers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintVouchers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Coordinator)
	if err != nil {
		return sdkerrors.Wrap(profile.ErrInvalidCoordAddress, err.Error())
	}

	if !sdk.Coins(msg.Shares).IsValid() {
		return sdkerrors.Wrap(ErrInvalidShares, sdk.Coins(msg.Shares).String())
	}

	if sdk.Coins(msg.Shares).Empty() {
		return sdkerrors.Wrap(ErrInvalidShares, "shares is empty")
	}

	return nil
}
