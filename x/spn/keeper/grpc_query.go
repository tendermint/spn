package keeper

import (
	"github.com/tendermint/spn/x/spn/types"
)

var _ types.QueryServer = Keeper{}
