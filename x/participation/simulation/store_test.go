package simulation_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	participationsim "github.com/tendermint/spn/x/participation/simulation"
	"testing"
)

func TestRandomAccWithBalance(t *testing.T) {
	var (
		ctx, tk, _ = testkeeper.NewTestSetup(t)
		r          = sample.Rand()
		accs       = simulation.RandomAccounts(r, 5)
	)

	// give one account balance
	newCoins := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)))
	err := tk.BankKeeper.MintCoins(ctx, minttypes.ModuleName, newCoins)
	require.NoError(t, err)
	err = tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, minttypes.ModuleName, accs[0].Address, newCoins)
	require.NoError(t, err)

	tests := []struct {
		name         string
		accounts     []simulation.Account
		desiredCoins sdk.Coins
		wantAccount  simulation.Account
		found        bool
	}{
		{
			name:     "no accounts with balance",
			accounts: accs[1:],
			found:    false,
		},
		{
			name:         "one account has balance",
			accounts:     accs,
			desiredCoins: newCoins,
			wantAccount:  accs[0],
			found:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, coins, found := participationsim.RandomAccWithBalance(ctx, r, tk.BankKeeper, tt.accounts, tt.desiredCoins)
			require.Equal(t, tt.found, found)
			if !tt.found {
				return
			}

			require.Equal(t, tt.wantAccount, got)
			require.Equal(t, tt.desiredCoins, coins)
		})
	}
}

func TestRandomAuction(t *testing.T) {

}

func TestRandomAccWithAvailableAllocations(t *testing.T) {

}
