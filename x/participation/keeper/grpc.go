package keeper

import (
	"github.com/tendermint/spn/x/participation/types"
)

var _ types.QueryServer = Keeper{}
