package types

const (
	// AccountAddressPrefix is the prefix used for account bech32 addresses
	AccountAddressPrefix = "spn"

	// Name is the name of the application
	Name = "spn"

	// DefaultChainID is the default chain ID used
	DefaultChainID = "spn-1"

	// DefaultUnbondingPeriod is the default unbonding time in seconds
	// 1814400 represents 21 days
	DefaultUnbondingPeriod = 1814400

	// MinimalUnbondingPeriod is the minimal unbonding time that can be set for a chain
	// it must greater than 1 because trusting period for the IBC client is unbonding period - 1
	// and trusting period can't be 0
	MinimalUnbondingPeriod = 2

	// DefaultRevisionHeight is the revision height used by default for creating the monitoring IBC client
	DefaultRevisionHeight = 1

	// TotalShareNumber is the default number of total share for an underlying supply asset
	TotalShareNumber = 100000
)
