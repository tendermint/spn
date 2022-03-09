package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	launchtypes "github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

type ProfileKeeper interface {
	GetValidator(ctx sdk.Context, address string) (profiletypes.Validator, bool)
	GetValidatorByOperatorAddress(ctx sdk.Context, operatorAddress string) (profiletypes.ValidatorByOperatorAddress, bool)
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (uint64, error)
}

type LaunchKeeper interface {
	GetChain(ctx sdk.Context, launchID uint64) (launchtypes.Chain, bool)
}

// AccountKeeper defines the expected account keeper used for simulations
type AccountKeeper interface {
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
}
