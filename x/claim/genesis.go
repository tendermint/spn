package claim

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tendermint/spn/x/claim/keeper"
	"github.com/tendermint/spn/x/claim/types"
)

// InitGenesis initializes the claim module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the claimRecord
	for _, elem := range genState.ClaimRecords {
		k.SetClaimRecord(ctx, elem)
	}
	// Set all the mission
	for _, elem := range genState.Missions {
		k.SetMission(ctx, elem)
	}

	k.SetAirdropSupply(ctx, genState.AirdropSupply)

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the claim module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ClaimRecords = k.GetAllClaimRecord(ctx)
	genesis.Missions = k.GetAllMission(ctx)
	airdropSupply, found := k.GetAirdropSupply(ctx)
	if found {
		genesis.AirdropSupply = airdropSupply
	} else {
		// set to 0uspn otherwise
		genesis.AirdropSupply = types.DefaultGenesis().AirdropSupply
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
