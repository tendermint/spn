package types

import (
	types "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	// List the different formats for users
	UserFormatAccountAddress = iota
)

// ChatUser represents a user that can be used with the Chat module
type ChatUser interface {
	Addresses() []sdk.AccAddress
	Username() string
	Identifier() string
	ToProtobuf() (User, error)
}

// DecodeChatUser decodes the protobuf user into an addressable user
func (user User) DecodeChatUser() (ChatUser, error) {
	// CHeck the user format
	switch user.Format {
	case UserFormatAccountAddress:
		// Decoding into a account address user
		var accountAddressUser AccountAddressUser
		err := accountAddressUser.Unmarshal(user.Data.GetValue())
		if err != nil {
			return nil, sdkerrors.Wrap(ErrInvalidPoll, "user cannot be decoded")
		}
		return accountAddressUser, nil
	default:
		// The user format is not recognize
		return nil, sdkerrors.Wrap(ErrInvalidPoll, "invalid user format")
	}
}

// NewAccountAddressUser create a new user that is represented only by its account address
func NewAccountAddressUser(
	address sdk.AccAddress,
) (AccountAddressUser, error) {
	accountAddressUser := new(AccountAddressUser)

	if address.Empty() {
		return *accountAddressUser, sdkerrors.Wrap(ErrInvalidPoll, "empty address")
	}
	accountAddressUser.AccountAddress = address

	return *accountAddressUser, nil
}

// Addresses returns the account address of the user
func (aaUser AccountAddressUser) Addresses() []sdk.AccAddress {
	return []sdk.AccAddress{aaUser.AccountAddress}
}

// Username returns a username that can be displayed in the chat
func (aaUser AccountAddressUser) Username() string {
	return aaUser.AccountAddress.String()
}

// Identifier returns a string that uniquely idenitfy a user
// This ensure a user votes only once for a poll
func (aaUser AccountAddressUser) Identifier() string {
	return aaUser.AccountAddress.String()
}

// ToProtobuf returns protobuf encoded user
func (aaUser AccountAddressUser) ToProtobuf() (User, error) {
	user := new(User)

	encodedUser, err := types.NewAnyWithValue(&aaUser)
	if err != nil {
		return *user, sdkerrors.Wrap(ErrInvalidPoll, "user cannot be encoded")
	}
	user.Data = encodedUser
	user.Format = UserFormatAccountAddress

	return *user, nil
}
