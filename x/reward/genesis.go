package reward

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/reward/keeper"
	"github.com/tendermint/spn/x/reward/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the rewardPool
	for _, elem := range genState.RewardPoolList {
		k.SetRewardPool(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.RewardPoolList = k.GetAllRewardPool(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
