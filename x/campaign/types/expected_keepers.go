package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/tendermint/spn/x/profile/types"
)

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
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	IterateAllBalances(ctx sdk.Context, cb func(sdk.AccAddress, sdk.Coin) bool)
	IterateAccountBalances(ctx sdk.Context, addr sdk.AccAddress, cb func(sdk.Coin) bool)
}

type ProfileKeeper interface {
	GetAllCoordinator(ctx sdk.Context) []types.Coordinator
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, found bool)
	GetCoordinatorAddressFromID(ctx sdk.Context, id uint64) (string, bool)
}

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}
