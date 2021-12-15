package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgSettleRequest(t *testing.T) {
	var (
		addr1                       = sample.Address()
		addr2                       = sample.Address()
		addr3                       = sample.Address()
		addr4                       = sample.Address()
		addr5                       = sample.Address()
		addr6                       = sample.Address()
		addr7                       = sample.Address()
		addr8                       = sample.Address()
		addr9                       = sample.Address()
		coordinator1                = sample.Coordinator(sample.Address())
		coordinator2                = sample.Coordinator(sample.Address())
		invalidChain                = uint64(1000)
		k, pk, _, srv, _, _, sdkCtx = setupMsgServer(t)
		ctx                         = sdk.WrapSDKContext(sdkCtx)
	)

	coordinator1.CoordinatorID = pk.AppendCoordinator(sdkCtx, coordinator1)
	coordinator2.CoordinatorID = pk.AppendCoordinator(sdkCtx, coordinator2)

	chains := createNChainForCoordinator(k, sdkCtx, coordinator1.CoordinatorID, 3)
	chains[0].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	k.SetChain(sdkCtx, chains[1])

	requests := createRequestsFromSamples(k, sdkCtx, chains[2].LaunchID, []RequestSample{
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr1), Creator: addr1},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr2), Creator: addr2},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr3), Creator: addr3},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr4), Creator: addr4},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr5), Creator: addr5},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr6), Creator: addr6},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr7), Creator: addr7},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr8), Creator: addr8},
		{Content: sample.GenesisAccountContent(chains[2].LaunchID, addr9), Creator: addr9},
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
				LaunchID:  invalidChain,
				Settler:   coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "launch triggered chain",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[0].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "coordinator not found",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[1].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrChainInactive,
		},
		{
			name: "no permission error",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   coordinator2.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "approve an invalid request",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: 99999999,
				Approve:   true,
			},
			err: types.ErrRequestNotFound,
		},
		{
			name: "coordinator approves chain request 1",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: requests[0].RequestID,
				Approve:   true,
			},
			checkAddr: addr1,
		},
		{
			name: "coordinator approves chain request 2",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: requests[1].RequestID,
				Approve:   true,
			},
			checkAddr: addr2,
		},
		{
			name: "coordinator approves chain request 3",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: requests[2].RequestID,
				Approve:   true,
			},
			checkAddr: addr3,
		},
		{
			name: "coordinator approves chain request 4",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: requests[3].RequestID,
				Approve:   true,
			},
			checkAddr: addr4,
		},
		{
			name: "coordinator rejects chain request 5",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   coordinator1.Address,
				RequestID: requests[4].RequestID,
				Approve:   false,
			},
			checkAddr: addr5,
		},
		{
			name: "request creator rejects its own chain request 6",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   addr6,
				RequestID: requests[5].RequestID,
				Approve:   false,
			},
			checkAddr: addr6,
		},
		{
			name: "request creator rejects not its own chain request 7",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   addr6,
				RequestID: requests[6].RequestID,
				Approve:   false,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "request creator approves its own chain request 8",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   addr8,
				RequestID: requests[7].RequestID,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "request creator approves not its own chain request 9",
			msg: types.MsgSettleRequest{
				LaunchID:  chains[2].LaunchID,
				Settler:   addr8,
				RequestID: requests[8].RequestID,
				Approve:   true,
			},
			err: types.ErrNoAddressPermission,
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

			_, found := k.GetRequest(sdkCtx, tt.msg.LaunchID, tt.msg.RequestID)
			require.False(t, found, "request not removed")

			_, found = k.GetGenesisAccount(sdkCtx, tt.msg.LaunchID, tt.checkAddr)
			require.Equal(t, tt.msg.Approve, found, "request apply not performed")
		})
	}
}
