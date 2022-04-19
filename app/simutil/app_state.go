package simutil

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	AuctionCoinDenom = "auction"
	MaxNumBonded     = 300
)

// CustomAppStateFn returns the initial application state using the simulation parameters.
func CustomAppStateFn(cdc codec.JSONCodec, simManager *module.SimulationManager) simtypes.AppStateFn {
	return func(r *rand.Rand, accs []simtypes.Account, config simtypes.Config,
	) (appState json.RawMessage, simAccs []simtypes.Account, chainID string, genesisTimestamp time.Time) {
		genesisTimestamp = simtypes.RandTimestamp(r)

		numAccs := int64(len(accs))
		genesisState := simapp.NewDefaultGenesisState(cdc)
		chainID = config.ChainID
		appParams := make(simtypes.AppParams)

		maxInitialCoin := math.MaxInt64 / (numAccs + MaxNumBonded)

		// generate a random amount of initial stake coins and a random initial
		// number of bonded accounts
		var initialStake, numInitiallyBonded int64
		appParams.GetOrGenerate(
			cdc, simappparams.StakePerAccount, &initialStake, r,
			func(r *rand.Rand) { initialStake = r.Int63n(maxInitialCoin) },
		)
		appParams.GetOrGenerate(
			cdc, simappparams.InitiallyBondedValidators, &numInitiallyBonded, r,
			func(r *rand.Rand) { numInitiallyBonded = r.Int63n(MaxNumBonded) },
		)

		if numInitiallyBonded > numAccs {
			numInitiallyBonded = numAccs
		}

		fmt.Printf(
			`Selected randomly generated parameters for simulated genesis:
{
  stake_per_account: "%d",
  initially_bonded_validators: "%d",
  num_accs: "%d"
}
`, initialStake, numInitiallyBonded, numAccs,
		)

		simState := &module.SimulationState{
			AppParams:    appParams,
			Cdc:          cdc,
			Rand:         r,
			GenState:     genesisState,
			Accounts:     accs,
			InitialStake: initialStake,
			NumBonded:    numInitiallyBonded,
			GenTimestamp: genesisTimestamp,
		}

		simManager.GenerateGenesisStates(simState)
		appState, err := json.Marshal(genesisState)
		if err != nil {
			panic(err)
		}

		rawState := make(map[string]json.RawMessage)
		err = json.Unmarshal(appState, &rawState)
		if err != nil {
			panic(err)
		}

		stakingStateBz, ok := rawState[stakingtypes.ModuleName]
		if !ok {
			panic("staking genesis state is missing")
		}

		stakingState := new(stakingtypes.GenesisState)
		err = cdc.UnmarshalJSON(stakingStateBz, stakingState)
		if err != nil {
			panic(err)
		}
		// compute not bonded balance
		notBondedTokens := sdk.ZeroInt()
		for _, val := range stakingState.Validators {
			if val.Status != stakingtypes.Unbonded {
				continue
			}
			notBondedTokens = notBondedTokens.Add(val.GetTokens())
		}
		notBondedCoins := sdk.NewCoin(stakingState.Params.BondDenom, notBondedTokens)
		// edit bank state to make it have the not bonded pool tokens
		bankStateBz, ok := rawState[banktypes.ModuleName]
		if !ok {
			panic("bank genesis state is missing")
		}
		bankState := new(banktypes.GenesisState)
		err = cdc.UnmarshalJSON(bankStateBz, bankState)
		if err != nil {
			panic(err)
		}

		// add auction coins randomly to accounts
		totalNewCoins := sdk.NewCoins()
		for i, balance := range bankState.Balances {
			if r.Int63n(100) < 20 {
				auctionCoinAmt := r.Int63n(maxInitialCoin)
				auctionCoin := sdk.NewCoin(AuctionCoinDenom, sdk.NewInt(auctionCoinAmt))
				newBalance := balance.Coins.Add(auctionCoin)
				bankState.Balances[i].Coins = newBalance
				totalNewCoins = totalNewCoins.Add(auctionCoin)
			}
		}

		// add new coins to the genesis supply
		bankState.Supply = bankState.Supply.Add(totalNewCoins...)

		stakingAddr := authtypes.NewModuleAddress(stakingtypes.NotBondedPoolName).String()
		var found bool
		for _, balance := range bankState.Balances {
			if balance.Address == stakingAddr {
				found = true
				break
			}
		}
		if !found {
			bankState.Balances = append(bankState.Balances, banktypes.Balance{
				Address: stakingAddr,
				Coins:   sdk.NewCoins(notBondedCoins),
			})
		}

		// override bank parameters to always enable transfers
		bankState.Params.SendEnabled = banktypes.SendEnabledParams{}
		bankState.Params.DefaultSendEnabled = true

		// change appState back
		rawState[stakingtypes.ModuleName] = cdc.MustMarshalJSON(stakingState)
		rawState[banktypes.ModuleName] = cdc.MustMarshalJSON(bankState)

		// replace appstate
		appState, err = json.Marshal(rawState)
		if err != nil {
			panic(err)
		}
		return appState, accs, chainID, genesisTimestamp
	}
}
