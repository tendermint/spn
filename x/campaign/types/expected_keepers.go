package types

import sdk "github.com/cosmos/cosmos-sdk/types"

type LaunchKeeper interface {
	CreateNewChain(
		ctx sdk.Context,
		coordinatorID uint64,
		genesisChainID,
		sourceURL,
		sourceHash,
		genesisURL,
		genesisHash string,
		hasCampaign bool,
		campaignID uint64,
		isMainnet bool,
	) (uint64, error)
}

type BankKeeper interface {
	// Methods imported from bank should be defined here
}

type ProfileKeeper interface {
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
}
