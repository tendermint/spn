package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/tendermint/spn/testutil/sample"
)

// DelegateN creates N delegations from the same address
func (tk TestKeepers) DelegateN(ctx sdk.Context, address string, shareAmt int64, n int) ([]stakingtypes.Delegation, sdk.Dec) {

	items := make([]stakingtypes.Delegation, n)
	totalShares := sdk.ZeroDec()

	for i := range items {
		items[i] = tk.Delegate(ctx, address, shareAmt)
		totalShares = totalShares.Add(items[i].Shares)
	}

	return items, totalShares
}

// Delegate creates a sample delegation and sets it in the keeper
func (tk TestKeepers) Delegate(ctx sdk.Context, address string, amt int64) stakingtypes.Delegation {
	del := sample.Delegation(tk.T, address)
	del.Shares = sdk.NewDec(amt)
	tk.StakingKeeper.SetDelegation(ctx, del)
	return del
}
