package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgSettleRequest(t *testing.T) {
	var (
		coordinator1             = sample.Coordinator()
		coordinator2             = sample.Coordinator()
		invalidChain, _          = sample.ChainID(0)
		k, pk, srv, _, sdkCtx, _ = setupMsgServer(t)
		ctx                      = sdk.WrapSDKContext(sdkCtx)
		requests                 = createNRequest(k, sdkCtx, 6)
	)

	coordinator1.CoordinatorId = pk.AppendCoordinator(sdkCtx, coordinator1)
	coordinator2.CoordinatorId = pk.AppendCoordinator(sdkCtx, coordinator2)

	chains := createNChainForCoordinator(k, sdkCtx, coordinator1.CoordinatorId, 2)
	chains[0].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[0])
	chains[1].ChainID = "foo"
	k.SetChain(sdkCtx, chains[1])

	tests := []struct {
		name string
		msg  types.MsgSettleRequest
		err  error
	}{
		{
			name: "invalid chain",
			msg: types.MsgSettleRequest{
				ChainID:     invalidChain,
				Coordinator: coordinator1.Address,
				RequestID:   requests[0].RequestID,
			},
			err: sdkerrors.Wrap(types.ErrChainNotFound, invalidChain),
		}, {
			name: "launch triggered chain",
			msg: types.MsgSettleRequest{
				ChainID:     chains[0].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[0].RequestID,
			},
			err: sdkerrors.Wrap(types.ErrTriggeredLaunch, chains[0].ChainID),
		}, {
			name: "no permission error",
			msg: types.MsgSettleRequest{
				ChainID:     chains[1].ChainID,
				Coordinator: coordinator2.Address,
				RequestID:   requests[0].RequestID,
			},
			err: sdkerrors.Wrap(types.ErrNoAddressPermission, coordinator2.Address),
		}, {
			name: "settle a invalid request",
			msg: types.MsgSettleRequest{
				ChainID:     chains[1].ChainID,
				Coordinator: pk.GetCoordinatorAddressFromID(sdkCtx, chains[1].CoordinatorID),
				RequestID:   99999999,
			},
			err: sdkerrors.Wrapf(types.ErrRequestNotFound,
				"request 99999999 for chain %s not found", chains[1].ChainID),
		}, {
			name: "add chain 1 request 1",
			msg: types.MsgSettleRequest{
				ChainID:     chains[1].ChainID,
				Coordinator: pk.GetCoordinatorAddressFromID(sdkCtx, chains[1].CoordinatorID),
				RequestID:   requests[0].RequestID,
			},
		}, {
			name: "add chain 1 request 2",
			msg: types.MsgSettleRequest{
				ChainID:     chains[1].ChainID,
				Coordinator: pk.GetCoordinatorAddressFromID(sdkCtx, chains[1].CoordinatorID),
				RequestID:   requests[1].RequestID,
			},
		}, {
			name: "add chain 1 request 3",
			msg: types.MsgSettleRequest{
				ChainID:     chains[1].ChainID,
				Coordinator: pk.GetCoordinatorAddressFromID(sdkCtx, chains[1].CoordinatorID),
				RequestID:   requests[2].RequestID,
			},
		}, {
			name: "add chain 2 request 5",
			msg: types.MsgSettleRequest{
				ChainID:     chains[1].ChainID,
				Coordinator: pk.GetCoordinatorAddressFromID(sdkCtx, chains[1].CoordinatorID),
				RequestID:   requests[5].RequestID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.SettleRequest(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)

			_, found := k.GetRequest(sdkCtx, tt.msg.ChainID, tt.msg.RequestID)
			require.False(t, found, "request not removed")
		})
	}
}
