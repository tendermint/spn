package keeper

import (
	"github.com/tendermint/spn/x/identity/types"
)

var _ types.QueryServer = Keeper{}
