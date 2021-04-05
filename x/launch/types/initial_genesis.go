package types

// NewInitialGenesisDefault returns an initial genesis that is the default genesis
func NewInitialGenesisDefault() *InitialGenesis {
	var ig InitialGenesis

	ig.Source = &InitialGenesis_DefaultGenesis{
		DefaultGenesis: &DefaultInitialGenesis{},
	}

	return &ig
}

// NewInitialGenesisURL returns an initial genesis from a genesis URL
func NewInitialGenesisURL(genesisURL GenesisURL) *InitialGenesis {
	var ig InitialGenesis

	ig.Source = &InitialGenesis_GenesisURL{
		GenesisURL: &genesisURL,
	}

	return &ig
}
