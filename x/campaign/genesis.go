package campaign

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the campaign
	for _, elem := range genState.CampaignList {
		k.SetCampaign(ctx, elem)
	}
	// Set campaign counter
	k.SetCampaignCounter(ctx, genState.CampaignCounter)

	// Set all the campaignChains
	for _, elem := range genState.CampaignChainsList {
		k.SetCampaignChains(ctx, elem)

	}

	// Set all the mainnetAccount
	for _, elem := range genState.MainnetAccountList {
		k.SetMainnetAccount(ctx, elem)
	}

	// Set all the mainnetVestingAccount
	for _, elem := range genState.MainnetVestingAccountList {
		k.SetMainnetVestingAccount(ctx, elem)
	}

	k.SetParams(ctx, genState.Params)

	// set maximum shares constant value
	k.SetTotalShares(ctx, genState.TotalShares)

	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.CampaignList = k.GetAllCampaign(ctx)
	genesis.CampaignCounter = k.GetCampaignCounter(ctx)
	genesis.CampaignChainsList = k.GetAllCampaignChains(ctx)
	genesis.MainnetAccountList = k.GetAllMainnetAccount(ctx)
	genesis.MainnetVestingAccountList = k.GetAllMainnetVestingAccount(ctx)
	genesis.Params = k.GetParams(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
