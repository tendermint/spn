package keeper

import (
	"github.com/tendermint/spn/x/profile/types"
)

var _ types.QueryServer = Keeper{}
