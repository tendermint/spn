package simutil

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	AuctionCoinDenom = "auction"
)

// CustomAppStateFn returns the initial application state using the simulation parameters.
func CustomAppStateFn(cdc codec.JSONCodec, simManager *module.SimulationManager) simtypes.AppStateFn {
	return func(r *rand.Rand, accs []simtypes.Account, config simtypes.Config,
	) (appState json.RawMessage, simAccs []simtypes.Account, chainID string, genesisTimestamp time.Time) {
		genesisTimestamp = simtypes.RandTimestamp(r)

		chainID = config.ChainID
		appParams := make(simtypes.AppParams)
		appState, simAccs = simapp.AppStateRandomizedFn(simManager, r, cdc, accs, genesisTimestamp, appParams)

		rawState := make(map[string]json.RawMessage)
		err := json.Unmarshal(appState, &rawState)
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

		// add more coins randomly to accounts
		totalNewCoins := sdk.NewCoins()
		for i, balance := range bankState.Balances {
			if r.Int63n(100) < 20 {
				auctionCoinAmt := r.Int63n(999_000_000) + 1_000_000
				bondCoinAmt := r.Int63n(999_000_000_000) + 1_000_000_000
				auctionCoin := sdk.NewCoin(AuctionCoinDenom, sdk.NewInt(auctionCoinAmt))
				bondCoin := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(bondCoinAmt))
				newBalance := balance.Coins.Add(auctionCoin, bondCoin)
				bankState.Balances[i].Coins = newBalance
				totalNewCoins = totalNewCoins.Add(auctionCoin, bondCoin)
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

		// change appState back
		rawState[stakingtypes.ModuleName] = cdc.MustMarshalJSON(stakingState)
		rawState[banktypes.ModuleName] = cdc.MustMarshalJSON(bankState)

		// replace appstate
		appState, err = json.Marshal(rawState)
		if err != nil {
			panic(err)
		}
		return appState, simAccs, chainID, genesisTimestamp
	}
}
