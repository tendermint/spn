package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	"math/rand"
)

// CreateCoordinator creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateCoordinator(ctx context.Context, r *rand.Rand) (id uint64, address sdk.AccAddress) {
	addr := sample.AccAddress(r)
	res, err := tm.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     addr.String(),
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(tm.T, err)
	return res.CoordinatorID, addr
}
