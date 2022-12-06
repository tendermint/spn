package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	campaigntypes "github.com/tendermint/spn/x/campaign/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

type CampaignKeeper interface {
	GetParams(ctx sdk.Context) (params campaigntypes.Params)
	GetCampaign(ctx sdk.Context, id uint64) (campaigntypes.Campaign, bool)
	AddChainToCampaign(ctx sdk.Context, campaignID, launchID uint64) error
	GetAllCampaign(ctx sdk.Context) (list []campaigntypes.Campaign)
	GetCampaignChains(ctx sdk.Context, campaignID uint64) (val campaigntypes.CampaignChains, found bool)
	MainnetAccountBalanceAll(
		c context.Context,
		req *campaigntypes.QueryAllMainnetAccountBalanceRequest,
	) (*campaigntypes.QueryAllMainnetAccountBalanceResponse, error)
	MainnetAccountBalance(
		c context.Context,
		req *campaigntypes.QueryGetMainnetAccountBalanceRequest,
	) (*campaigntypes.QueryGetMainnetAccountBalanceResponse, error)
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
