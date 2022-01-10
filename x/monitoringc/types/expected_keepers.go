package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

type LaunchKeeper interface {
	// Methods imported from launch should be defined here
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ClientKeeper is imported to add the ability to create IBC Client from the module
type ClientKeeper interface {
	CreateClient(
		ctx sdk.Context, clientState exported.ClientState, consensusState exported.ConsensusState,
	) (string, error)
}
