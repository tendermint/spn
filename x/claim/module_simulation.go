package claim

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	claimsimulation "github.com/tendermint/spn/x/claim/simulation"
	"github.com/tendermint/spn/x/claim/types"
)

const (
	airdropDenom = "drop"

	opWeightMsgClaimInitial          = "op_weight_msg_claim_initial"
	defaultWeightMsgClaimInitial int = 50

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	claimRecords := make([]types.ClaimRecord, len(simState.Accounts))
	totalSupply := sdkmath.ZeroInt()
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()

		// fill claim records from simulation accounts
		accSupply := sdkmath.NewIntFromUint64(simState.Rand.Uint64() % 1000)
		claimRecords[i] = types.ClaimRecord{
			Claimable: accSupply,
			Address:   acc.Address.String(),
		}
		totalSupply = totalSupply.Add(accSupply)
	}

	// define some decimal numbers for mission weights
	dec1, err := sdk.NewDecFromStr("0.4")
	if err != nil {
		panic(err)
	}
	dec2, err := sdk.NewDecFromStr("0.3")
	if err != nil {
		panic(err)
	}

	claimGenesis := types.GenesisState{
		Params:        types.DefaultParams(),
		AirdropSupply: sdk.NewCoin(airdropDenom, totalSupply),
		Missions: []types.Mission{
			{
				MissionID:   0,
				Description: "initial claim",
				Weight:      dec1,
			},
			{
				MissionID:   1,
				Description: "mission 1",
				Weight:      dec2,
			},
			{
				MissionID:   2,
				Description: "mission 2",
				Weight:      dec2,
			},
		},
		InitialClaim: types.InitialClaim{
			Enabled:   true,
			MissionID: 0,
		},
		ClaimRecords: claimRecords,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&claimGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgClaimInitial int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgClaimInitial, &weightMsgClaimInitial, nil,
		func(_ *rand.Rand) {
			weightMsgClaimInitial = defaultWeightMsgClaimInitial
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgClaimInitial,
		claimsimulation.SimulateMsgClaimInitial(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
