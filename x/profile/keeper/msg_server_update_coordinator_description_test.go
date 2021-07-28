package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestMsgUpdateCoordinatorDescription(t *testing.T) {
	var (
		addr        = sample.AccAddress()
		msgCoord    = msgCreateCoordinator()
		ctx, k, srv = setupMsgServerAndKeeper(t)
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
			_, err := srv.UpdateCoordinatorDescription(wCtx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			coordByAddr, found := k.GetCoordinatorByAddress(ctx, tt.msg.Address)
			assert.True(t, found, "coordinator by address not found")

			coord := k.GetCoordinator(ctx, coordByAddr.CoordinatorId)
			assert.True(t, found, "coordinator id not found")
			assert.EqualValues(t, tt.msg.Address, coord.Address)
			assert.EqualValues(t, coordByAddr.CoordinatorId, coord.CoordinatorId)

			if len(tt.msg.Description.Identity) > 0 {
				assert.EqualValues(t, tt.msg.Description.Identity, coord.Description.Identity)
			}
			if len(tt.msg.Description.Website) > 0 {
				assert.EqualValues(t, tt.msg.Description.Website, coord.Description.Website)
			}
			if len(tt.msg.Description.Details) > 0 {
				assert.EqualValues(t, tt.msg.Description.Details, coord.Description.Details)
			}

		})
	}
}
