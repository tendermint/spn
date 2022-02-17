package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	valtypes "github.com/tendermint/spn/pkg/types"
)

const TypeMsgSetValidatorConsAddress = "set_validator_cons_address"

var _ sdk.Msg = &MsgSetValidatorConsAddress{}

func NewMsgSetValidatorConsAddress(
	validatorAddress,
	signature,
	keyType,
	chainID string,
	nonce uint64,
	validatorConsPubKey []byte,
) *MsgSetValidatorConsAddress {
	return &MsgSetValidatorConsAddress{
		ValidatorAddress:    validatorAddress,
		ValidatorConsPubKey: validatorConsPubKey,
		ValidatorKeyType:    keyType,
		Signature:           signature,
		Nonce:               nonce,
		ChainID:             chainID,
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
	valPubKey, err := valtypes.NewValidatorConsPubKey(msg.ValidatorConsPubKey, msg.ValidatorKeyType)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidValidatorKey, msg.ValidatorKeyType)
	}
	if !valPubKey.VerifySignature(msg.Nonce, msg.ChainID, msg.Signature) {
		return sdkerrors.Wrap(ErrInvalidValidatorSignature, msg.Signature)
	}
	return nil
}
