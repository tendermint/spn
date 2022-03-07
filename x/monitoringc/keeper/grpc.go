package keeper

import (
	"github.com/tendermint/spn/x/monitoringc/types"
)

var _ types.QueryServer = Keeper{}
