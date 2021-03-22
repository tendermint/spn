package types

import "errors"

// NewInitialGenesisDefault returns an initial genesis that is the default genesis
func NewInitialGenesisDefault() InitialGenesis {
	var ig InitialGenesis

	ig.Source = &InitialGenesis_DefaultGenesis{
		DefaultGenesis: nil,
	}

	return ig
}

// NewInitialGenesisURL returns an initial genesis from a genesis URL
func NewInitialGenesisURL(genesisURL GenesisURL) InitialGenesis {
	var ig InitialGenesis

	ig.Source = &InitialGenesis_GenesisURL{
		GenesisURL: &genesisURL,
	}

	return ig
}

// GetType returns the type of the initial genesis
func (ig InitialGenesis) GetType() (InitialGenesisType, error) {
	switch ig.Source.(type) {
	case *InitialGenesis_DefaultGenesis:
		return InitialGenesisType_DEFAULT, nil
	case *InitialGenesis_GenesisURL:
		return InitialGenesisType_URL, nil
	default:
		return InitialGenesisType_DEFAULT, errors.New("unrecognized initial genesis type")
	}
}

// GenesisURL returns the genesis URL of the initial genesis if it's its type
func (ig InitialGenesis) GenesisURL() (gURL GenesisURL, err error) {
	genesisURL, ok := ig.Source.(*InitialGenesis_GenesisURL)

	if !ok {
		return gURL, errors.New("initial genesis is not a genesis URL")
	}

	return *genesisURL.GenesisURL, nil
}

