package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	campaigntypes "github.com/tendermint/spn/x/campaign/types"
)

// Mint mints the specified coins into the account balance
func (tk TestKeepers) Mint(ctx sdk.Context, address string, coins sdk.Coins) {
	sdkAddr, err := sdk.AccAddressFromBech32(address)
	require.NoError(tk.T, err)
	require.NoError(tk.T, tk.BankKeeper.MintCoins(ctx, campaigntypes.ModuleName, coins))
	require.NoError(tk.T, tk.BankKeeper.SendCoinsFromModuleToAccount(ctx, campaigntypes.ModuleName, sdkAddr, coins))
}
