package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	valtypes "github.com/tendermint/spn/pkg/types"
)

const TypeMsgSetValidatorConsAddress = "set_validator_cons_address"

var _ sdk.Msg = &MsgSetValidatorConsAddress{}

func NewMsgSetValidatorConsAddress(validatorAddress, signature, keyType string, validatorConsPubKey []byte) *MsgSetValidatorConsAddress {
	return &MsgSetValidatorConsAddress{
		ValidatorAddress:    validatorAddress,
		ValidatorConsPubKey: validatorConsPubKey,
		ValidatorKeyType:    keyType,
		Signature:           signature,
	}
}

func (msg *MsgSetValidatorConsAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetValidatorConsAddress) Type() string {
	return TypeMsgSetValidatorConsAddress
}

func (msg *MsgSetValidatorConsAddress) GetSigners() []sdk.AccAddress {
	valAddress, err := sdk.AccAddressFromBech32(msg.ValidatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{valAddress}
}

func (msg *MsgSetValidatorConsAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetValidatorConsAddress) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.ValidatorAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	if _, err := valtypes.NewValidatorConsPubKey(msg.ValidatorConsPubKey, msg.ValidatorKeyType); err != nil {
		return sdkerrors.Wrap(ErrInvalidValidatorKey, string(msg.ValidatorConsPubKey))
	}
	return nil
}
