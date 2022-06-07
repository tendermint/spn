package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"
	"strings"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/tendermint/spn/x/mint/types"
)

const (
	keyInflationRateChange       = "InflationRateChange"
	keyInflationMax              = "InflationMax"
	keyInflationMin              = "InflationMin"
	keyGoalBonded                = "GoalBonded"
	keyDistributionProportions   = "DistributionProportions"
	keyDevelopmentFundRecipients = "DevelopmentFundRecipients"
)

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyInflationRateChange,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenInflationRateChange(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyInflationMax,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenInflationMax(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyInflationMin,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenInflationMin(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyGoalBonded,
			func(r *rand.Rand) string {
				return fmt.Sprintf("\"%s\"", GenGoalBonded(r))
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyDistributionProportions,
			func(r *rand.Rand) string {
				proportions := GenDistributionProportions(r)
				return fmt.Sprintf("{\"staking\":\"%s\",\"incentives\":\"%s\",\"development_fund\":\"%s\",\"community_pool\":\"%s\"}",
					proportions.Staking.String(), proportions.Incentives.String(), proportions.DevelopmentFund.String(), proportions.CommunityPool.String())
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyDevelopmentFundRecipients,
			func(r *rand.Rand) string {
				weightedAddrs := GenDevelopmentFundRecipients(r)
				weightedAddrsStr := make([]string, 0)
				for _, wa := range weightedAddrs {
					s := fmt.Sprintf("{\"address\":\"%s\",\"weight\":\"%s\"}", wa.Address, wa.Weight.String())
					weightedAddrsStr = append(weightedAddrsStr, s)
				}
				return fmt.Sprintf("[%s]", strings.Join(weightedAddrsStr, ","))
			},
		),
	}
}
