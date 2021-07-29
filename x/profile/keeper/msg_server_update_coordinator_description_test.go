package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription(t *testing.T) {
	var (
		addr        = sample.AccAddress()
		msgCoord    = msgCreateCoordinator()
		ctx, k, srv = setupMsgServer(t)
		wCtx        = sdk.WrapSDKContext(ctx)
	)
	if _, err := srv.CreateCoordinator(wCtx, &msgCoord); err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name string
		msg  types.MsgUpdateCoordinatorDescription
		err  error
	}{
		{
			name: "not found address",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: addr,
			},
			err: sdkerrors.Wrap(types.ErrCoordAddressNotFound, addr),
		}, {
			name: "update one value",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: msgCoord.Address,
				Description: &types.CoordinatorDescription{
					Identity: "update",
				},
			},
		}, {
			name: "update all values",
			msg: types.MsgUpdateCoordinatorDescription{
				Address: msgCoord.Address,
				Description: &types.CoordinatorDescription{
					Identity: "update",
					Website:  "update",
					Details:  "update",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var oldCoord types.Coordinator
			if tt.err == nil {
				coordByAddr, found := k.GetCoordinatorByAddress(ctx, tt.msg.Address)
				require.True(t, found, "coordinator by address not found")
				oldCoord = k.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			}

			_, err := srv.UpdateCoordinatorDescription(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			coordByAddr, found := k.GetCoordinatorByAddress(ctx, tt.msg.Address)
			coord := k.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			require.True(t, found, "coordinator id not found")
			require.EqualValues(t, tt.msg.Address, coord.Address)
			require.EqualValues(t, coordByAddr.CoordinatorId, coord.CoordinatorId)

			if len(tt.msg.Description.Identity) > 0 {
				require.EqualValues(t, tt.msg.Description.Identity, coord.Description.Identity)
			} else {
				require.EqualValues(t, oldCoord.Description.Identity, coord.Description.Identity)
			}

			if len(tt.msg.Description.Website) > 0 {
				require.EqualValues(t, tt.msg.Description.Website, coord.Description.Website)
			} else {
				require.EqualValues(t, oldCoord.Description.Website, coord.Description.Website)
			}

			if len(tt.msg.Description.Details) > 0 {
				require.EqualValues(t, tt.msg.Description.Details, coord.Description.Details)
			} else {
				require.EqualValues(t, oldCoord.Description.Details, coord.Description.Details)
			}
		})
	}
}
