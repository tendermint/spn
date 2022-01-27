package keeper

import (
	"github.com/tendermint/spn/x/reward/types"
)

var _ types.QueryServer = Keeper{}
