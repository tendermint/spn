package types

import (
	"errors"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrortypes "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m RequestContent) Validate() error {
	switch requestContent := m.Content.(type) {
	case *RequestContent_GenesisAccount:
		return requestContent.GenesisAccount.Validate()
	case *RequestContent_VestingAccount:
		return requestContent.VestingAccount.Validate()
	case *RequestContent_GenesisValidator:
		return requestContent.GenesisValidator.Validate()
	case *RequestContent_AccountRemoval:
		return requestContent.AccountRemoval.Validate()
	case *RequestContent_ValidatorRemoval:
		return requestContent.ValidatorRemoval.Validate()
	default:
		return errors.New("unrecognized request content")
	}
}

func (m RequestContent) IsValidForMainnet() error {
	switch _ := m.Content.(type) {
	case *RequestContent_GenesisAccount:
		return errors.New("GenesisAccount request can't be used for mainnet")
	case *RequestContent_VestingAccount:
		return errors.New("VestingAccount request can't be used for mainnet")
	case *RequestContent_AccountRemoval:
		return errors.New("AccountRemoval request can't be used for mainnet")
	}
	return nil
}

// NewGenesisAccount returns a RequestContent containing an GenesisAccount
func NewGenesisAccount(launchID uint64, address string, coins sdk.Coins) RequestContent {
	return RequestContent{
		Content: &RequestContent_GenesisAccount{
			GenesisAccount: &GenesisAccount{
				LaunchID: launchID,
				Address:  address,
				Coins:    coins,
			},
		},
	}
}

// Validate implements GenesisAccount validation
func (m GenesisAccount) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	if !m.Coins.IsValid() || m.Coins.Empty() {
		return sdkerrors.Wrap(ErrInvalidCoins, m.Address)
	}
	return nil
}

// NewVestingAccount returns a RequestContent containing a VestingAccount
func NewVestingAccount(launchID uint64, address string, vestingOptions VestingOptions) RequestContent {
	return RequestContent{
		Content: &RequestContent_VestingAccount{
			VestingAccount: &VestingAccount{
				LaunchID:       launchID,
				Address:        address,
				VestingOptions: vestingOptions,
			},
		},
	}
}

// Validate implements VestingAccount validation
func (m VestingAccount) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid validator address (%s)", err)
	}

	if err := m.VestingOptions.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidVestingOption, err.Error())
	}
	return nil
}

// NewGenesisValidator returns a RequestContent containing a GenesisValidator
func NewGenesisValidator(
	launchID uint64,
	address string,
	genTx,
	consPubKey []byte,
	selfDelegation sdk.Coin,
	peer Peer,
) RequestContent {
	return RequestContent{
		Content: &RequestContent_GenesisValidator{
			GenesisValidator: &GenesisValidator{
				LaunchID:       launchID,
				Address:        address,
				GenTx:          genTx,
				ConsPubKey:     consPubKey,
				SelfDelegation: selfDelegation,
				Peer:           peer,
			},
		},
	}
}

// Validate implements GenesisValidator validation
func (m GenesisValidator) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid account address (%s)", err)
	}

	if len(m.GenTx) == 0 {
		return sdkerrors.Wrap(ErrInvalidGenTx, "empty gentx")
	}

	if len(m.ConsPubKey) == 0 {
		return sdkerrors.Wrap(ErrInvalidConsPubKey, "empty consensus public key")
	}

	if !m.SelfDelegation.IsValid() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "")
	}

	if m.SelfDelegation.IsZero() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "self delegation is zero")
	}

	if err := m.Peer.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidPeer, err.Error())
	}
	return nil
}

// NewAccountRemoval returns a RequestContent containing an AccountRemoval
func NewAccountRemoval(address string) RequestContent {
	return RequestContent{
		Content: &RequestContent_AccountRemoval{
			AccountRemoval: &AccountRemoval{
				Address: address,
			},
		},
	}
}

// Validate implements AccountRemoval validation
func (m AccountRemoval) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	return nil
}

// NewValidatorRemoval returns a RequestContent containing a ValidatorRemoval
func NewValidatorRemoval(address string) RequestContent {
	return RequestContent{
		Content: &RequestContent_ValidatorRemoval{
			ValidatorRemoval: &ValidatorRemoval{
				ValAddress: address,
			},
		},
	}
}

// Validate implements ValidatorRemoval validation
func (m ValidatorRemoval) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.ValAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrortypes.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}
