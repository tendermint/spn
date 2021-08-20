package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Validatable defines the interface for a type that can be validated
type Validatable interface {
	Validate() error
}

// NewAccountRemoval returns an InitialGenesis containing a DefaultInitialGenesis
func NewAccountRemoval(address string) RequestContent {
	return RequestContent{
		Content: &RequestContent_AccountRemoval{
			AccountRemoval: &AccountRemoval{
				Address: address,
			},
		},
	}
}

var _ Validatable = &AccountRemoval{}

// Validate implements AccountRemoval validation
func (c AccountRemoval) Validate() error {
	_, err := sdk.AccAddressFromBech32(c.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	return nil
}

var _ Validatable = &GenesisValidator{}

// Validate implements GenesisValidator validation
func (c GenesisValidator) Validate() error {
	_, err := sdk.AccAddressFromBech32(c.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	_, _, err = ParseChainID(c.ChainID)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidChainID, c.ChainID)
	}

	if len(c.GenTx) == 0 {
		return sdkerrors.Wrap(ErrInvalidGenTx, "empty gentx")
	}

	if len(c.ConsPubKey) == 0 {
		return sdkerrors.Wrap(ErrInvalidConsPubKey, "empty consensus public key")
	}

	if !c.SelfDelegation.IsValid() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "")
	}

	if c.SelfDelegation.IsZero() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "self delegation is zero")
	}

	if c.Peer == "" {
		return sdkerrors.Wrap(ErrInvalidPeer, "empty peer")
	}
	return nil
}

var _ Validatable = &GenesisAccount{}

// Validate implements GenesisAccount validation
func (c GenesisAccount) Validate() error {
	_, err := sdk.AccAddressFromBech32(c.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	_, _, err = ParseChainID(c.ChainID)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidChainID, c.ChainID)
	}

	if !c.Coins.IsValid() || c.Coins.Empty() {
		return sdkerrors.Wrap(ErrInvalidCoins, c.Address)
	}
	return nil
}

var _ Validatable = &ValidatorRemoval{}

// Validate implements ValidatorRemoval validation
func (c ValidatorRemoval) Validate() error {
	_, err := sdk.AccAddressFromBech32(c.ValAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}


var _ Validatable = &VestedAccount{}

// Validate implements VestedAccount validation
func (c VestedAccount) Validate() error {
	_, err := sdk.AccAddressFromBech32(c.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	_, _, err = ParseChainID(c.ChainID)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidChainID, c.ChainID)
	}

	if !c.StartingBalance.IsValid() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, c.StartingBalance.String())
	}

	if err := c.VestingOptions.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidVestingOption, err.Error())
	}
	return nil
}
