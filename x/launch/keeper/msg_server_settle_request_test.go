package keeper_test

import (
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgSettleRequest(t *testing.T) {
	var (
		addr1                    = sample.AccAddress()
		addr2                    = sample.AccAddress()
		addr3                    = sample.AccAddress()
		addr4                    = sample.AccAddress()
		addr5                    = sample.AccAddress()
		coordinator1             = sample.Coordinator()
		coordinator2             = sample.Coordinator()
		invalidChain, _          = sample.ChainID(0)
		k, pk, srv, _, sdkCtx, _ = setupMsgServer(t)
		ctx                      = sdk.WrapSDKContext(sdkCtx)
		invalidContent, _        = codectypes.NewAnyWithValue(&types.Request{})
	)

	coordinator1.CoordinatorId = pk.AppendCoordinator(sdkCtx, coordinator1)
	coordinator2.CoordinatorId = pk.AppendCoordinator(sdkCtx, coordinator2)

	chains := createNChainForCoordinator(k, sdkCtx, coordinator1.CoordinatorId, 3)
	chains[0].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	k.SetChain(sdkCtx, chains[1])

	requests := createRequests(k, sdkCtx, chains[2].ChainID, []*codectypes.Any{
		invalidContent,
		sample.GenesisAccountContent(chains[2].ChainID, addr1),
		sample.GenesisAccountContent(chains[2].ChainID, addr2),
		sample.GenesisAccountContent(chains[2].ChainID, addr3),
		sample.GenesisAccountContent(chains[2].ChainID, addr4),
		sample.GenesisAccountContent(chains[2].ChainID, addr5),
	})

	tests := []struct {
		name      string
		msg       types.MsgSettleRequest
		checkAddr string
		err       error
	}{
		{
			name: "invalid chain",
			msg: types.MsgSettleRequest{
				ChainID:     invalidChain,
				Coordinator: coordinator1.Address,
				RequestID:   requests[0].RequestID,
				Approve:     true,
			},
			err: types.ErrChainNotFound,
		}, {
			name: "launch triggered chain",
			msg: types.MsgSettleRequest{
				ChainID:     chains[0].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[0].RequestID,
				Approve:     true,
			},
			err: types.ErrTriggeredLaunch,
		}, {
			name: "coordinator not found",
			msg: types.MsgSettleRequest{
				ChainID:     chains[1].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[0].RequestID,
				Approve:     true,
			},
			err: types.ErrChainInactive,
		}, {
			name: "no permission error",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator2.Address,
				RequestID:   requests[0].RequestID,
				Approve:     true,
			},
			err: types.ErrNoAddressPermission,
		}, {
			name: "approve an invalid request",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   99999999,
				Approve:     true,
			},
			err: types.ErrRequestNotFound,
		}, {
			name: "invalid request content",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[0].RequestID,
				Approve:     true,
			},
			err: spnerrors.ErrCritical,
		}, {
			name: "approve chain request 1",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[1].RequestID,
				Approve:     true,
			},
			checkAddr: addr1,
		}, {
			name: "approve chain request 2",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[2].RequestID,
				Approve:     true,
			},
			checkAddr: addr2,
		}, {
			name: "approve chain request 3",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[3].RequestID,
				Approve:     true,
			},
			checkAddr: addr3,
		}, {
			name: "approve chain request 4",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[4].RequestID,
				Approve:     true,
			},
			checkAddr: addr4,
		}, {
			name: "reject chain request 5",
			msg: types.MsgSettleRequest{
				ChainID:     chains[2].ChainID,
				Coordinator: coordinator1.Address,
				RequestID:   requests[5].RequestID,
				Approve:     false,
			},
			checkAddr: addr5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := srv.SettleRequest(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			_, found := k.GetRequest(sdkCtx, tt.msg.ChainID, tt.msg.RequestID)
			require.False(t, found, "request not removed")

			_, found = k.GetGenesisAccount(sdkCtx, tt.msg.ChainID, tt.checkAddr)
			require.Equal(t, tt.msg.Approve, found, "request apply not performed")
		})
	}
}
