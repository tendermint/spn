package keeper

import (
	"github.com/tendermint/spn/x/project/types"
)

var _ types.QueryServer = Keeper{}
