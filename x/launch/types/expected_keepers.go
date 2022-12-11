package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	projecttypes "github.com/tendermint/spn/x/project/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

type ProjectKeeper interface {
	GetProject(ctx sdk.Context, id uint64) (projecttypes.Project, bool)
	AddChainToProject(ctx sdk.Context, projectID, launchID uint64) error
	GetAllProject(ctx sdk.Context) (list []projecttypes.Project)
	GetProjectChains(ctx sdk.Context, projectID uint64) (val projecttypes.ProjectChains, found bool)
	MainnetAccountBalanceAll(
		c context.Context,
		req *projecttypes.QueryAllMainnetAccountBalanceRequest,
	) (*projecttypes.QueryAllMainnetAccountBalanceResponse, error)
	MainnetAccountBalance(
		c context.Context,
		req *projecttypes.QueryGetMainnetAccountBalanceRequest,
	) (*projecttypes.QueryGetMainnetAccountBalanceResponse, error)
}

type MonitoringConsumerKeeper interface {
	ClearVerifiedClientIDs(ctx sdk.Context, launchID uint64)
}

type ProfileKeeper interface {
	CoordinatorIDFromAddress(ctx sdk.Context, address string) (id uint64, err error)
	GetCoordinator(ctx sdk.Context, id uint64) (val profiletypes.Coordinator, found bool)
}

type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	ValidateBalance(ctx sdk.Context, addr sdk.AccAddress) error
	HasBalance(ctx sdk.Context, addr sdk.AccAddress, amt sdk.Coin) bool
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	GetAccountsBalances(ctx sdk.Context) []banktypes.Balance
}

type DistributionKeeper interface {
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}
