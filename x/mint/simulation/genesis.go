package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/mint/types"
)

// Simulation parameter constants
const (
	Inflation                 = "inflation"
	InflationRateChange       = "inflation_rate_change"
	InflationMax              = "inflation_max"
	InflationMin              = "inflation_min"
	GoalBonded                = "goal_bonded"
	DistributionProportions   = "distribution_proportions"
	DevelopmentFundRecipients = "development_fund_recipients"
)

// GenInflation randomized Inflation
func GenInflation(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(99)), 2)
}

// GenInflationRateChange randomized InflationRateChange
func GenInflationRateChange(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(int64(r.Intn(99)), 2)
}

// GenInflationMax randomized InflationMax
func GenInflationMax(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(20, 2)
}

// GenInflationMin randomized InflationMin
func GenInflationMin(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(7, 2)
}

// GenGoalBonded randomized GoalBonded
func GenGoalBonded(r *rand.Rand) sdk.Dec {
	return sdk.NewDecWithPrec(67, 2)
}

// GenDistributionProportions randomized DistributionProportions
func GenDistributionProportions(r *rand.Rand) types.DistributionProportions {
	staking := r.Int63n(99)
	left := int64(100) - staking
	incentives := r.Int63n(left)
	left -= incentives
	devs := r.Int63n(left)
	communityPool := left - devs

	return types.DistributionProportions{
		Staking:         sdk.NewDecWithPrec(staking, 2),
		Incentives:      sdk.NewDecWithPrec(incentives, 2),
		DevelopmentFund: sdk.NewDecWithPrec(devs, 2),
		CommunityPool:   sdk.NewDecWithPrec(communityPool, 2),
	}
}

func GenDevelopmentFundRecipients(r *rand.Rand) []types.WeightedAddress {
	numAddrs := r.Intn(51)
	addrs := make([]types.WeightedAddress, 0)
	remainWeight := sdk.NewDec(1)
	maxRandWeight := sdk.NewDecWithPrec(15, 3)
	minRandWeight := sdk.NewDecWithPrec(5, 3)
	for i := 0; i < numAddrs; i++ {
		// each address except the last can have a max of 2% weight and a min of 0.5%
		weight := simtypes.RandomDecAmount(r, maxRandWeight).Add(minRandWeight)
		if i == numAddrs-1 {
			// use residual weight if last address
			weight = remainWeight
		} else {
			remainWeight = remainWeight.Sub(weight)
		}
		wa := types.WeightedAddress{
			Address: sample.Address(r),
			Weight:  weight,
		}
		addrs = append(addrs, wa)
	}
	return addrs
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// minter
	var inflation sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, Inflation, &inflation, simState.Rand,
		func(r *rand.Rand) { inflation = GenInflation(r) },
	)

	// params
	var inflationRateChange sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, InflationRateChange, &inflationRateChange, simState.Rand,
		func(r *rand.Rand) { inflationRateChange = GenInflationRateChange(r) },
	)

	var inflationMax sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, InflationMax, &inflationMax, simState.Rand,
		func(r *rand.Rand) { inflationMax = GenInflationMax(r) },
	)

	var inflationMin sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, InflationMin, &inflationMin, simState.Rand,
		func(r *rand.Rand) { inflationMin = GenInflationMin(r) },
	)

	var goalBonded sdk.Dec
	simState.AppParams.GetOrGenerate(
		simState.Cdc, GoalBonded, &goalBonded, simState.Rand,
		func(r *rand.Rand) { goalBonded = GenGoalBonded(r) },
	)

	var distributionProportions types.DistributionProportions
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DistributionProportions, &distributionProportions, simState.Rand,
		func(r *rand.Rand) { distributionProportions = GenDistributionProportions(r) },
	)

	var developmentFundRecipients []types.WeightedAddress
	simState.AppParams.GetOrGenerate(
		simState.Cdc, DevelopmentFundRecipients, &developmentFundRecipients, simState.Rand,
		func(r *rand.Rand) { developmentFundRecipients = GenDevelopmentFundRecipients(r) },
	)

	mintDenom := sdk.DefaultBondDenom
	blocksPerYear := uint64(60 * 60 * 8766 / 5)
	params := types.NewParams(mintDenom, inflationRateChange, inflationMax, inflationMin, goalBonded, blocksPerYear, distributionProportions, developmentFundRecipients)

	mintGenesis := types.NewGenesisState(types.InitialMinter(inflation), params)

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
