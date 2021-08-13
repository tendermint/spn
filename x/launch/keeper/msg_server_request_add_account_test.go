package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestAddAccount(t *testing.T) {
	var (
		invalidChain, _            = sample.ChainID(0)
		coordAddr                  = sample.AccAddress()
		addr1                      = sample.AccAddress()
		addr2                      = sample.AccAddress()
		addr3                      = sample.AccAddress()
		k, pk, srv, _, sdkCtx, cdc = setupMsgServer(t)
		ctx                        = sdk.WrapSDKContext(sdkCtx)
	)

	coordID := pk.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordAddr,
	})
	chains := createNChainForCoordinator(k, sdkCtx, coordID, 4)
	chains[3].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[3])

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
			err: sdkerrors.Wrap(types.ErrChainNotFound, invalidChain),
		}, {
			name: "launch triggered chain",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[3].ChainID,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			err: sdkerrors.Wrap(types.ErrTriggeredLaunch, chains[3].ChainID),
		}, {
			name: "add chain 1 request 1",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[0].ChainID,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			wantID: 0,
		}, {
			name: "add chain 2 request 1",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[1].ChainID,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			wantID: 0,
		}, {
			name: "add chain 2 request 2",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[1].ChainID,
				Address: addr2,
				Coins:   sample.Coins(),
			},
			wantID: 1,
		}, {
			name: "add chain 3 request 1",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[2].ChainID,
				Address: addr1,
				Coins:   sample.Coins(),
			},
			wantID: 0,
		}, {
			name: "add chain 3 request 2",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[2].ChainID,
				Address: addr2,
				Coins:   sample.Coins(),
			},
			wantID: 1,
		}, {
			name: "add chain 3 request 3",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[2].ChainID,
				Address: addr3,
				Coins:   sample.Coins(),
			},
			wantID: 2,
		}, {
			name: "add coordinator account",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[2].ChainID,
				Address: coordAddr,
				Coins:   sample.Coins(),
			},
			wantApprove: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.RequestAddAccount(ctx, &tt.msg)
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			if !tt.wantApprove {
				request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, got.RequestID)
				require.True(t, found, "request not found")
				require.Equal(t, tt.wantID, request.RequestID)

				content, err := request.UnpackGenesisAccount(cdc)
				require.NoError(t, err)
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
