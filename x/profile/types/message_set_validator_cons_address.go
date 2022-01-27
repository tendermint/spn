package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/pkg/types"
)

const TypeMsgSetValidatorConsAddress = "set_validator_cons_address"

var _ sdk.Msg = &MsgSetValidatorConsAddress{}

func NewMsgSetValidatorConsAddress(signature string, validatorKey []byte) *MsgSetValidatorConsAddress {
	return &MsgSetValidatorConsAddress{
		ValidatorKey: validatorKey,
		Signature:    signature,
	}
}

func (msg *MsgSetValidatorConsAddress) Route() string {
	return RouterKey
}

func (msg *MsgSetValidatorConsAddress) Type() string {
	return TypeMsgSetValidatorConsAddress
}

func (msg *MsgSetValidatorConsAddress) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgSetValidatorConsAddress) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetValidatorConsAddress) ValidateBasic() error {
	if _, err := types.LoadValidatorKey(msg.ValidatorKey); err != nil {
		return sdkerrors.Wrapf(ErrInvalidValidatorKey, "invalid validator key (%s)", err)
	}
	return nil
}
