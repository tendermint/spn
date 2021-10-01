package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestAddAccount(t *testing.T) {
	var (
		invalidChain                = uint64(1000)
		coordAddr                   = sample.AccAddress()
		addr1                       = sample.AccAddress()
		addr2                       = sample.AccAddress()
		addr3                       = sample.AccAddress()
		k, pk, _, srv, _, _, sdkCtx = setupMsgServer(t)
		ctx                         = sdk.WrapSDKContext(sdkCtx)
	)

	coordID := pk.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordAddr,
	})
	chains := createNChainForCoordinator(k, sdkCtx, coordID, 6)
	chains[0].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	k.SetChain(sdkCtx, chains[1])
	chains[5].IsMainnet = true
	k.SetChain(sdkCtx, chains[5])

	tests := []struct {
		name        string
		msg         types.MsgRequestAddAccount
		wantID      uint64
		wantApprove bool
		err         error
	}{
		{
			name: "invalid chain",
			msg: types.MsgRequestAddAccount{
				ChainID: invalidChain,
				Address: sample.AccAddress(),
				Coins:   sample.Coins(),
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "launch triggered chain",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[0].Id,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "coordinator not found",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[1].Id,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			err: types.ErrChainInactive,
		},
		{
			name: "add chain 3 request 1",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[2].Id,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			wantID: 0,
		},
		{
			name: "add chain 4 request 1",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[3].Id,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			wantID: 0,
		},
		{
			name: "add chain 4 request 2",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[3].Id,
				Address: addr2,
				Coins:   sample.Coins(),
			},
			wantID: 1,
		},
		{
			name: "add chain 5 request 1",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[4].Id,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			wantID: 0,
		},
		{
			name: "add chain 5 request 2",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[4].Id,
				Address: addr2,
				Coins:   sample.Coins(),
			},
			wantID: 1,
		},
		{
			name: "add chain 5 request 3",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[4].Id,
				Address: addr3,
				Coins:   sample.Coins(),
			},
			wantID: 2,
		},
		{
			name: "request from coordinator is pre-approved",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[4].Id,
				Address: coordAddr,
				Coins:   sample.Coins(),
			},
			wantApprove: true,
		},
		{
			name: "failing request from coordinator",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[4].Id,
				Address: coordAddr,
				Coins:   sample.Coins(),
			},
			err: types.ErrAccountAlreadyExist,
		},
		{
			name: "is mainnet chain",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[5].Id,
				Address: coordAddr,
				Coins:   sample.Coins(),
			},
			err: types.ErrChainIsMainnet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.RequestAddAccount(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			if !tt.wantApprove {
				request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, got.RequestID)
				require.True(t, found, "request not found")
				require.Equal(t, tt.wantID, request.RequestID)

				content := request.Content.GetGenesisAccount()
				require.NotNil(t, content)
				require.Equal(t, tt.msg.Address, content.Address)
				require.Equal(t, tt.msg.ChainID, content.ChainID)
				require.Equal(t, tt.msg.Coins, content.Coins)
			} else {
				_, found := k.GetGenesisAccount(sdkCtx, tt.msg.ChainID, tt.msg.Address)
				require.True(t, found, "genesis account not found")
			}
		})
	}
}
