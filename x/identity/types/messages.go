package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	UsernameMaxLength  = 30
	TypeMsgSetUsername = "set_username"
)

// Verify interface at compile time
var (
	_ sdk.Msg = &MsgSetUsername{}
)

// NewMsgSetUsername set the username for an address
func NewMsgSetUsername(
	address sdk.AccAddress,
	username string,
) (*MsgSetUsername, error) {
	if !CheckUsername(username) {
		return nil, ErrInvalidUsername
	}

	return &MsgSetUsername{
		Address:  address,
		Username: username,
	}, nil
}

// Route implements the sdk.Msg interface.
func (msg MsgSetUsername) Route() string { return RouterKey }

// Type implements the sdk.Msg interface.
func (msg MsgSetUsername) Type() string { return TypeMsgSetUsername }

// GetSigners implements the sdk.Msg interface. It returns the address(es) that
// must sign over msg.GetSignBytes().
func (msg MsgSetUsername) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Address}
}

// GetSignBytes returns the message bytes to sign over.
func (msg MsgSetUsername) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface.
func (msg MsgSetUsername) ValidateBasic() error {
	if msg.Address.Empty() {
		return ErrInvalidAddress
	}
	if !CheckUsername(msg.Username) {
		return ErrInvalidUsername
	}

	return nil
}

// CheckUsername checks if the username is a alphanumeric string with hyphens or underscore
func CheckUsername(username string) bool {
	if len(username) == 0 {
		return false
	}
	if len(username) > UsernameMaxLength {
		return false
	}
	for _, c := range username {
		if !isAuthorizedChar(c) {
			return false
		}
	}
	return true
}

// Alphanumeric or hyphen character
func isAuthorizedChar(c rune) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '-' || c == '_'
}
