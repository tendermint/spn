package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"
)

type GenesisFile []byte

// NewGenesisFile returns a new genesis file from the bytes of the genesis
func NewGenesisFile(genesisBytes []byte) GenesisFile {
	return genesisBytes
}

// GetGenesisDoc returns the genesis in the tendermint/types.GenesisDoc format
func (g GenesisFile) GetGenesisDoc() (genesisDoc tmtypes.GenesisDoc, err error) {
	err = tmjson.Unmarshal(g, &genesisDoc)
	return genesisDoc, err
}

// SetChainID set the chain ID for the genesis file
func (g *GenesisFile) SetChainID(chainID string) error {
	// Unmarshal bytes
	var genesisObject tmtypes.GenesisDoc
	if err := tmjson.Unmarshal(*g, &genesisObject); err != nil {
		return err
	}

	genesisObject.ChainID = chainID

	// Marshal
	bz, err := tmjson.Marshal(genesisObject)
	if err != nil {
		return err
	}
	*g = bz

	return nil
}

// ValidateAndComplete checks that all necessary fields are present and fills in defaults for optional fields left empty
func (g *GenesisFile) ValidateAndComplete() error {
	// Unmarshal bytes
	var genesisObject tmtypes.GenesisDoc
	if err := tmjson.Unmarshal(*g, &genesisObject); err != nil {
		return sdkerrors.Wrap(ErrInvalidChain, err.Error())
	}

	// Validate and complete
	if err := genesisObject.ValidateAndComplete(); err != nil {
		return sdkerrors.Wrap(ErrInvalidChain, err.Error())
	}

	// Marshal
	bz, err := tmjson.Marshal(genesisObject)
	if err != nil {
		return err
	}
	*g = bz

	return nil
}

// ApplyProposals updates the genesis state from the application of a list of approved proposals
func (g *GenesisFile) ApplyProposals(cdc codec.JSONMarshaler, proposals []Proposal) error {
	// Unmarshal genesis
	var genesisDoc tmtypes.GenesisDoc
	err := tmjson.Unmarshal(*g, &genesisDoc)
	if err != nil {
		return err
	}

	// Apply proposals to the genesis
	for _, proposal := range proposals {
		err = applyProposal(cdc, &genesisDoc, proposal)
		if err != nil {
			return err
		}
	}

	// Marshal genesis
	bz, err := tmjson.Marshal(genesisDoc)
	if err != nil {
		return err
	}
	*g = bz

	return nil
}

// applyProposal updates the genesis state from the application of an approved proposal
func applyProposal(cdc codec.JSONMarshaler, genesisDoc *tmtypes.GenesisDoc, proposal Proposal) error {
	// Dispatch the proposal
	switch payload := proposal.Payload.(type) {
	case *Proposal_AddAccountPayload:
		return applyProposalAddAccount(cdc, genesisDoc, *payload.AddAccountPayload)
	case *Proposal_AddValidatorPayload:
		return applyProposalAddValidator(cdc, genesisDoc, *payload.AddValidatorPayload)
	default:
		return errors.New("invalid proposal")
	}
}

// applyProposalAddAccount updates the genesis state when an account is added
func applyProposalAddAccount(cdc codec.JSONMarshaler, genesisDoc *tmtypes.GenesisDoc, payload ProposalAddAccountPayload) error {
	// Create the account
	genAccount := authtypes.NewBaseAccount(
		payload.Address,
		nil,
		0,
		0,
	)

	// Create the balance for the account
	balance := banktypes.Balance{
		Address: payload.Address.String(),
		Coins:   payload.Coins.Sort(),
	}

	// Get the state for the app
	appState, err := genutiltypes.GenesisStateFromGenDoc(*genesisDoc)
	if err != nil {
		return fmt.Errorf("failed to get the application genesis state: %v", err)
	}

	// Add the account in the auth state
	cdcMarshaler := cdc.(codec.Marshaler)
	authGenState := authtypes.GetGenesisStateFromAppState(cdcMarshaler, appState)
	accs, err := authtypes.UnpackAccounts(authGenState.Accounts)
	if err != nil {
		return fmt.Errorf("failed to get accounts from any: %v", err)
	}
	if accs.Contains(payload.Address) {
		panic(fmt.Sprintf("the address %v is duplicated in approved proposals", payload.Address))
	}
	accs = append(accs, genAccount)
	accs = authtypes.SanitizeGenesisAccounts(accs)

	// Update the app state
	genAccs, err := authtypes.PackAccounts(accs)
	if err != nil {
		return fmt.Errorf("failed to convert accounts into any's: %v", err)
	}
	authGenState.Accounts = genAccs
	authGenStateBz, err := cdc.MarshalJSON(&authGenState)
	if err != nil {
		return fmt.Errorf("failed to marshal auth genesis state: %v", err)
	}
	appState[authtypes.ModuleName] = authGenStateBz

	// Add the account balance in the bank state
	bankGenState := banktypes.GetGenesisStateFromAppState(cdc, appState)
	bankGenState.Balances = append(bankGenState.Balances, balance)
	bankGenState.Balances = banktypes.SanitizeGenesisBalances(bankGenState.Balances)

	// Update the app state
	bankGenStateBz, err := cdc.MarshalJSON(bankGenState)
	if err != nil {
		return fmt.Errorf("failed to marshal bank genesis state: %v", err)
	}
	appState[banktypes.ModuleName] = bankGenStateBz

	// The the state of the app
	appStateJSON, err := json.Marshal(appState)
	if err != nil {
		return fmt.Errorf("failed to marshal application genesis state: %v", err)
	}
	genesisDoc.AppState = appStateJSON

	return nil
}

// applyProposalAddValidator updates the genesis state when a validator is added
func applyProposalAddValidator(cdc codec.JSONMarshaler, genesisDoc *tmtypes.GenesisDoc, payload ProposalAddValidatorPayload) error {
	// Get the state for the app
	appGenesisState, err := genutiltypes.GenesisStateFromGenDoc(*genesisDoc)
	if err != nil {
		return fmt.Errorf("failed to get the application genesis state: %v", err)
	}
	genesisState := genutiltypes.GetGenesisStateFromAppState(cdc, appGenesisState)

	// Encode the gentx
	genTxBz, err := cdc.MarshalJSON(payload.GenTx)
	if err != nil {
		return fmt.Errorf("failed to encode gentx: %v", err)
	}

	// Append the new gentx
	genesisState.GenTxs = append(genesisState.GenTxs, genTxBz)

	// Register the new state of the genesis
	appGenesisState = genutiltypes.SetGenesisStateInAppState(cdc, appGenesisState, genesisState)
	appState, err := json.MarshalIndent(appGenesisState, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode app genesis state: %v", err)
	}
	genesisDoc.AppState = appState

	return nil
}
