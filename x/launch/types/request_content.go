package types

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (m RequestContent) Validate() error {
	switch requestContent := m.Content.(type) {
	case *RequestContent_GenesisAccount:
		return requestContent.GenesisAccount.Validate()
	case *RequestContent_VestedAccount:
		return requestContent.VestedAccount.Validate()
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

// NewGenesisAccount returns a RequestContent containing an GenesisAccount
func NewGenesisAccount(chainID, address string, coins sdk.Coins) RequestContent {
	return RequestContent{
		Content: &RequestContent_GenesisAccount{
			GenesisAccount: &GenesisAccount{
				ChainID: chainID,
				Address: address,
				Coins: coins,
			},
		},
	}
}

// Validate implements GenesisAccount validation
func (m GenesisAccount) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	_, _, err = ParseChainID(m.ChainID)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidChainID, m.ChainID)
	}

	if !m.Coins.IsValid() || m.Coins.Empty() {
		return sdkerrors.Wrap(ErrInvalidCoins, m.Address)
	}
	return nil
}

// NewVestedAccount returns a RequestContent containing a VestedAccount
func NewVestedAccount(chainID, address string, startingBalance sdk.Coins, vestingOptions VestingOptions) RequestContent {
	return RequestContent{
		Content: &RequestContent_VestedAccount{
			VestedAccount: &VestedAccount{
				ChainID: chainID,
				Address: address,
				StartingBalance: startingBalance,
				VestingOptions: vestingOptions,
			},
		},
	}
}

// Validate implements VestedAccount validation
func (m VestedAccount) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	_, _, err = ParseChainID(m.ChainID)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidChainID, m.ChainID)
	}

	if !m.StartingBalance.IsValid() {
		return sdkerrors.Wrap(ErrInvalidSelfDelegation, m.StartingBalance.String())
	}

	if err := m.VestingOptions.Validate(); err != nil {
		return sdkerrors.Wrapf(ErrInvalidVestingOption, err.Error())
	}
	return nil
}

// NewGenesisValidator returns a RequestContent containing a GenesisValidator
func NewGenesisValidator(
	chainID,
	address string,
	genTx,
	consPubKey []byte,
	selfDelegation sdk.Coin,
	peer string,
	) RequestContent {
	return RequestContent{
		Content: &RequestContent_GenesisValidator{
			GenesisValidator: &GenesisValidator{
				ChainID:         chainID,
				Address:         address,
				GenTx: genTx,
				ConsPubKey: consPubKey,
				SelfDelegation: selfDelegation,
				Peer: peer,
			},
		},
	}
}

// Validate implements GenesisValidator validation
func (m GenesisValidator) Validate() error {
	_, err := sdk.AccAddressFromBech32(m.Address)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
	}
	_, _, err = ParseChainID(m.ChainID)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidChainID, m.ChainID)
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

	if m.Peer == "" {
		return sdkerrors.Wrap(ErrInvalidPeer, "empty peer")
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid account address (%s)", err)
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid validator address (%s)", err)
	}
	return nil
}

