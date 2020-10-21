package keeper

import (
	"github.com/tendermint/spn/x/chat/types"
)

var _ types.QueryServer = Keeper{}
