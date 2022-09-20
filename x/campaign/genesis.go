package campaign

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/campaign/keeper"
	"github.com/tendermint/spn/x/campaign/types"
)

// InitGenesis initializes the campaign module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the campaign
	for _, elem := range genState.Campaigns {
		k.SetCampaign(ctx, elem)
	}
	// Set campaign counter
	k.SetCampaignCounter(ctx, genState.CampaignCounter)

	// Set all the campaignChains
	for _, elem := range genState.CampaignChains {
		k.SetCampaignChains(ctx, elem)
	}

	// Set all the mainnetAccount
	for _, elem := range genState.MainnetAccounts {
		k.SetMainnetAccount(ctx, elem)
	}

	k.SetParams(ctx, genState.Params)

	// set maximum shares constant value
	k.SetTotalShares(ctx, genState.TotalShares)

	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the campaign module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.Campaigns = k.GetAllCampaign(ctx)
	genesis.CampaignCounter = k.GetCampaignCounter(ctx)
	genesis.CampaignChains = k.GetAllCampaignChains(ctx)
	genesis.MainnetAccounts = k.GetAllMainnetAccount(ctx)
	genesis.Params = k.GetParams(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
