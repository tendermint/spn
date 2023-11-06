package keeper

import (
	"context"
	"math/rand"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
)

// CreateCoordinator creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateChain(ctx context.Context, r *rand.Rand, coordAddress string, genesisURL string, hasProject bool, projectID uint64) uint64 {
	msgCreateChain := sample.MsgCreateChain(r, coordAddress, genesisURL, hasProject, projectID)
	res, err := tm.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(tm.T, err)
	return res.LaunchID
}
