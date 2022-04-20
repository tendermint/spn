package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
)

func TestTestKeepers_Mint(t *testing.T) {
	sdkCtx, tk, _ := testkeeper.NewTestSetup(t)
	r := sample.Rand()
	ctx := sdk.WrapSDKContext(sdkCtx)
	address := sample.Address(r)
	coins, otherCoins := sample.Coins(r), sample.Coins(r)

	getBalances := func(address string) sdk.Coins {
		res, err := tk.BankKeeper.AllBalances(ctx, &banktypes.QueryAllBalancesRequest{
			Address: address,
		})
		require.NoError(t, err)
		require.NotNil(t, res)
		return res.Balances
	}

	// should create the account
	tk.Mint(sdkCtx, address, coins)
	require.True(t, getBalances(address).IsEqual(coins))

	// should add the minted coins in the balance
	previousBalance := getBalances(address)
	tk.Mint(sdkCtx, address, otherCoins)
	require.True(t, getBalances(address).IsEqual(previousBalance.Add(otherCoins...)))
}
