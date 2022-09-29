package types

import (
	"errors"
	"regexp"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func (m RequestContent) Validate(launchID uint64) error {
	switch requestContent := m.Content.(type) {
	case *RequestContent_GenesisAccount:
		return requestContent.GenesisAccount.Validate(launchID)
	case *RequestContent_VestingAccount:
		return requestContent.VestingAccount.Validate(launchID)
	case *RequestContent_GenesisValidator:
		return requestContent.GenesisValidator.Validate(launchID)
	case *RequestContent_AccountRemoval:
		return requestContent.AccountRemoval.Validate()
	case *RequestContent_ValidatorRemoval:
		return requestContent.ValidatorRemoval.Validate()
	case *RequestContent_ParamChange:
		return requestContent.ParamChange.Validate(launchID)
	default:
		return errors.New("unrecognized request content")
	}
}

func (m RequestContent) IsValidForMainnet() error {
	switch m.Content.(type) {
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
func (m GenesisAccount) Validate(launchID uint64) error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidGenesisAddress, err.Error())
	}

	if !m.Coins.IsValid() || m.Coins.Empty() {
		return sdkerrors.Wrap(ErrInvalidCoins, m.Address)
	}

	if m.LaunchID != launchID {
		return ErrInvalidLaunchID
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
func (m VestingAccount) Validate(launchID uint64) error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidVestingAddress, err.Error())
	}

	if err := m.VestingOptions.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidVestingOption, err.Error())
	}

	if m.LaunchID != launchID {
		return ErrInvalidLaunchID
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
func (m GenesisValidator) Validate(launchID uint64) error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidValidatorAddress, err.Error())
	}

	if len(m.GenTx) == 0 {
		return sdkerrors.Wrap(ErrInvalidGenTx, "empty gentx")
	}

	if len(m.ConsPubKey) == 0 {
		return sdkerrors.Wrap(ErrInvalidConsPubKey, "empty consensus public key")
	}

	if !m.SelfDelegation.IsValid() {
		return ErrInvalidSelfDelegation
	}

	if m.SelfDelegation.IsZero() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, "self delegation is zero")
	}

	if err := m.Peer.Validate(); err != nil {
		return sdkerrors.Wrap(ErrInvalidPeer, err.Error())
	}

	if m.LaunchID != launchID {
		return ErrInvalidLaunchID
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
		return sdkerrors.Wrap(ErrInvalidGenesisAddress, err.Error())
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
		return sdkerrors.Wrap(ErrInvalidValidatorAddress, err.Error())
	}
	return nil
}

// NewParamChange returns a RequestContent containing a ParamChange
func NewParamChange(launchID uint64, module, param string, value []byte) RequestContent {
	return RequestContent{
		Content: &RequestContent_ParamChange{
			ParamChange: &ParamChange{
				LaunchID: launchID,
				Module:   module,
				Param:    param,
				Value:    value,
			},
		},
	}
}

// Validate implements ParamChange validation
func (m ParamChange) Validate(launchID uint64) error {
	if m.Module == "" || m.Param == "" {
		return ErrInvalidRequestContent
	}

	if !isStringAlphabetic(m.Module) || !isStringAlphabetic(m.Param) {
		return ErrInvalidRequestContent
	}

	if m.LaunchID != launchID {
		return ErrInvalidLaunchID
	}

	return nil
}
