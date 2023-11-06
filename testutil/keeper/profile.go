package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"math/rand"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

// CreateCoordinator creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateCoordinator(ctx context.Context, r *rand.Rand) (id uint64, address sdk.AccAddress) {
	return tm.CreateCoordinatorWithAddr(ctx, r, sample.Address(r))
}

// CreateCoordinator creates a coordinator in the store and returns ID with associated address
func (tm TestMsgServers) CreateCoordinatorWithAddr(ctx context.Context, r *rand.Rand, address string) (uint64, sdk.AccAddress) {
	addr := sdk.MustAccAddressFromBech32(address)
	res, err := tm.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     address,
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(tm.T, err)
	return res.CoordinatorID, addr
}
