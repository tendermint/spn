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

// AddressableUser represents a user that possesses an Cosmos SDK account address
type AddressableUser interface {
	Address() sdk.AccAddress
	ToProtobuf() (User, error)
}

// DecodeAddressableUser decodes the protobuf user into an addressable user
func (user User) DecodeAddressableUser() (AddressableUser, error) {
	var addressableUser AddressableUser

	// CHeck the user format
	switch user.Format {
	case UserFormatAccountAddress:
		// Decoding into a account address user
		var accountAddressUser AccountAddressUser
		err := accountAddressUser.Unmarshal(user.Data.GetValue())
		if err != nil {
			return accountAddressUser, sdkerrors.Wrap(ErrInvalidPoll, "user cannot be decoded")
		}
		return accountAddressUser, nil
	default:
		// The user format is not recognize
		return addressableUser, sdkerrors.Wrap(ErrInvalidPoll, "invalid user format")
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

// Address returns the account address of the user
func (aaUser AccountAddressUser) Address() sdk.AccAddress {
	return aaUser.AccountAddress
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
