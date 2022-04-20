package participation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

// InitGenesis initializes the participation module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the usedAllocations
	for _, elem := range genState.UsedAllocationsList {
		k.SetUsedAllocations(ctx, elem)
	}
	// Set all the auctionUsedAllocations
	for _, elem := range genState.AuctionUsedAllocationsList {
		k.SetAuctionUsedAllocations(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the participation module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.UsedAllocationsList = k.GetAllUsedAllocations(ctx)
	genesis.AuctionUsedAllocationsList = k.GetAllAuctionUsedAllocations(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
