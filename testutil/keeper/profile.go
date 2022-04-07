package keeper

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"math/rand"
)

// CreateCoordinator creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateCoordinator(ctx context.Context, r *rand.Rand) (id uint64, address string) {
	addr := sample.Address(r)
	res, err := tm.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     addr,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(tm.T, err)
	return res.CoordinatorID, addr
}
