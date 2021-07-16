package keeper

import (
	"github.com/tendermint/spn/x/account/types"
)

var _ types.QueryServer = Keeper{}
