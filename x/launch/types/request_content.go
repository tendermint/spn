package types

import (
	"errors"
	"fmt"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RequestContent defines the interface for a request content
type RequestContent interface {
	Validate() error
}

var _ RequestContent = &AccountRemoval{}

// Validate implements AccountRemoval validation
func (c AccountRemoval) Validate() error {
	_, err := sdk.AccAddressFromBech32(c.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	return nil
}

// UnpackAccountRemoval returns the AccountRemoval structure from the codec unpack
func (r Request) UnpackAccountRemoval(cdc codec.AnyUnpacker) (*AccountRemoval, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}

	result, ok := content.(*AccountRemoval)
	if !ok {
		return nil, errors.New("not a accountRemoval request")
	}
	return result, nil
}

var _ RequestContent = &GenesisValidator{}

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

// UnpackGenesisValidator returns the GenesisValidator structure from the codec unpack
func (r Request) UnpackGenesisValidator(cdc codec.AnyUnpacker) (*GenesisValidator, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}

	result, ok := content.(*GenesisValidator)
	if !ok {
		return nil, errors.New("not a genesisValidator request")
	}
	return result, nil
}

var _ RequestContent = &GenesisAccount{}

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

// UnpackGenesisAccount returns the GenesisAccount structure from the codec unpack
func (r Request) UnpackGenesisAccount(cdc codec.AnyUnpacker) (*GenesisAccount, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}

	result, ok := content.(*GenesisAccount)
	if !ok {
		return nil, errors.New("not a genesisAccount request")
	}
	return result, nil
}

var _ RequestContent = &ValidatorRemoval{}

// Validate implements ValidatorRemoval validation
func (c ValidatorRemoval) Validate() error {
	_, err := sdk.AccAddressFromBech32(c.ValAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}

// UnpackValidatorRemoval returns the ValidatorRemoval structure from the codec unpack
func (r Request) UnpackValidatorRemoval(cdc codec.AnyUnpacker) (*ValidatorRemoval, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}

	removeValidator, ok := content.(*ValidatorRemoval)
	if !ok {
		return nil, errors.New("not a validatorRemoval request")
	}
	return removeValidator, nil
}

var _ RequestContent = &VestedAccount{}

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

	if c.VestingOptions == nil {
		return sdkerrors.Wrap(ErrInvalidAccountOption, c.Address)
	}

	cdc := codec.NewInterfaceRegistry()
	RegisterInterfaces(cdc)

	var option VestingOptions
	if err := cdc.UnpackAny(c.VestingOptions, &option); err != nil {
		return sdkerrors.Wrap(ErrInvalidAccountOption, err.Error())
	}

	switch option.(type) {
	case *DelayedVesting:
	default:
		return sdkerrors.Wrap(ErrInvalidAccountOption, "unknown vested account option type")
	}
	return option.Validate()
}

// UnpackVestedAccount returns the VestedAccount structure from the codec unpack
func (r Request) UnpackVestedAccount(cdc codec.AnyUnpacker) (*VestedAccount, error) {
	if r.Content == nil {
		return nil, fmt.Errorf("empty request content for request %d", r.RequestID)
	}
	var content RequestContent
	err := cdc.UnpackAny(r.Content, &content)
	if err != nil {
		return nil, err
	}
	result, ok := content.(*VestedAccount)
	if !ok {
		return nil, errors.New("not a vestedAccount request")
	}
	return result, nil
}
