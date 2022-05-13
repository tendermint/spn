package keeper

import (
	"github.com/tendermint/spn/x/claim/types"
)

var _ types.QueryServer = Keeper{}
