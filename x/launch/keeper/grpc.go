package keeper

import (
	"github.com/tendermint/spn/x/launch/types"
)

var _ types.QueryServer = Keeper{}
