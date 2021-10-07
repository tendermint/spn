package types

import (
	"github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type AccountKeeper interface {
	GetAccount(ctx types.Context, addr types.AccAddress) authtypes.AccountI
}

type BankKeeper interface {
	SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins
}
