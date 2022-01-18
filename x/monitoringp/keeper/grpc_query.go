package keeper

import (
	"github.com/tendermint/spn/x/monitoringp/types"
)

var _ types.QueryServer = Keeper{}
