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
func (tm TestMsgServers) CreateCoordinator(ctx context.Context, r *rand.Rand, optionalAddress ...string) (id uint64, address sdk.AccAddress) {
	var addr sdk.AccAddress
	if optionalAddress == nil {
		addr = sample.AccAddress(r)
	} else {
		addr = sdk.MustAccAddressFromBech32(optionalAddress[0])
	}
	res, err := tm.ProfileSrv.CreateCoordinator(ctx, &profiletypes.MsgCreateCoordinator{
		Address:     addr.String(),
		Description: sample.CoordinatorDescription(r),
	})
	require.NoError(tm.T, err)
	return res.CoordinatorID, addr
}
