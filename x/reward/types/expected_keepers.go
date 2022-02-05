package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

type AuthKeeper interface {
	// Methods imported from auth should be defined here
}

type ProfileKeeper interface {
	GetValidator(ctx sdk.Context, address string) (val profiletypes.Validator, found bool)
}

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
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}
