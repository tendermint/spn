package keeper

import (
	"github.com/tendermint/spn/x/campaign/types"
)

var _ types.QueryServer = Keeper{}
