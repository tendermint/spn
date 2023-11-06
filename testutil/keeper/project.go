package keeper

import (
	"context"
	"github.com/stretchr/testify/require"
	"math/rand"

	"github.com/tendermint/spn/testutil/sample"
)

// CreateCoordinator creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateProject(ctx context.Context, r *rand.Rand, coordAddress string) (id uint64) {
	msgCreateProject := sample.MsgCreateProject(r, coordAddress)
	resProject, err := tm.ProjectSrv.CreateProject(ctx, &msgCreateProject)
	require.NoError(tm.T, err)
	return resProject.ProjectID
}
