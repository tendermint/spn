package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRequestAddValidator = "request_add_validator"

var _ sdk.Msg = &MsgRequestAddValidator{}

func NewMsgRequestAddValidator(
	creator string,
	launchID uint64,
	valAddress string,
	genTx,
	consPubKey []byte,
	selfDelegation sdk.Coin,
	peer *Peer,
) *MsgRequestAddValidator {
	return &MsgRequestAddValidator{
		Creator:        creator,
		ValAddress:     valAddress,
		LaunchID:       launchID,
		GenTx:          genTx,
		ConsPubKey:     consPubKey,
		SelfDelegation: selfDelegation,
		Peer:           peer,
	}
}

func (msg *MsgRequestAddValidator) Route() string {
	return RouterKey
}

func (msg *MsgRequestAddValidator) Type() string {
	return TypeMsgRequestAddValidator
}

func (msg *MsgRequestAddValidator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRequestAddValidator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRequestAddValidator) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.ValAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	if len(msg.GenTx) == 0 {
		return sdkerrors.Wrap(ErrInvalidGenTx, "empty gentx")
	}

	if len(msg.ConsPubKey) == 0 {
		return sdkerrors.Wrap(ErrInvalidConsPubKey, "empty consensus public key")
	}

	if !msg.SelfDelegation.IsValid() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "")
	}

	if msg.SelfDelegation.IsZero() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "self delegation is zero")
	}

	if msg.Peer == nil {
		return sdkerrors.Wrap(ErrInvalidPeer, "null peer connection")
	}
	switch conn := msg.Peer.Connection.(type) {
	case *Peer_TcpAddress:
		if conn.TcpAddress == "" {
			return sdkerrors.Wrap(ErrInvalidPeer, "empty peer")
		}
	case *Peer_HttpTunnel:
		if conn.HttpTunnel.Name == "" {
			return sdkerrors.Wrap(ErrInvalidPeer, "empty http tunnel peer name")
		}
		if conn.HttpTunnel.Address == "" {
			return sdkerrors.Wrap(ErrInvalidPeer, "empty http tunnel peer address")
		}
	default:
		return sdkerrors.Wrap(ErrInvalidPeer, "invalid peer connection")
	}

	return nil
}
