package keeper

import (
	"github.com/tendermint/spn/x/genesis/types"
)

var _ types.QueryServer = Keeper{}
