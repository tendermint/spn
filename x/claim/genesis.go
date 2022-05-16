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
	for _, elem := range genState.ClaimRecordList {
		k.SetClaimRecord(ctx, elem)
	}
	// Set all the mission
	for _, elem := range genState.MissionList {
		k.SetMission(ctx, elem)
	}

	// Set mission count
	k.SetMissionCount(ctx, genState.MissionCount)
	k.SetAirdropSupply(ctx, genState.AirdropSupply)

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the claim module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ClaimRecordList = k.GetAllClaimRecord(ctx)
	genesis.MissionList = k.GetAllMission(ctx)
	genesis.MissionCount = k.GetMissionCount(ctx)
	airdropSupply, found := k.GetAirdropSupply(ctx)
	if found {
		genesis.AirdropSupply = airdropSupply
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
