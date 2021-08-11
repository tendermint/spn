package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestRemoveAccount(t *testing.T) {
	var (
		invalidChain, _            = sample.ChainID(0)
		addr1                      = sample.AccAddress()
		addr2                      = sample.AccAddress()
		addr3                      = sample.AccAddress()
		addr4                      = sample.AccAddress()
		k, pk, srv, _, sdkCtx, cdc = setupMsgServer(t)
		ctx                        = sdk.WrapSDKContext(sdkCtx)
		chains                     = createNChain(k, sdkCtx, 4)
	)
	chains[3].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[3])
	k.SetVestedAccount(sdkCtx, types.VestedAccount{
		ChainID: chains[2].ChainID,
		Address: addr4,
	})
	tests := []struct {
		name        string
		msg         types.MsgRequestRemoveAccount
		wantID      uint64
		wantApprove bool
		err         error
	}{
		{
			name: "invalid chain",
			msg: types.MsgRequestRemoveAccount{
				ChainID: invalidChain,
				Creator: addr1,
				Address: addr1,
			},
			err: sdkerrors.Wrap(types.ErrChainNotFound, invalidChain),
		}, {
			name: "launch triggered chain",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[3].ChainID,
				Creator: addr1,
				Address: addr1,
			},
			err: sdkerrors.Wrap(types.ErrTriggeredLaunch, addr1),
		}, {
			name: "no permission error",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[0].ChainID,
				Creator: addr1,
				Address: addr3,
			},
			err: sdkerrors.Wrap(types.ErrNoAddressPermission, addr1),
		}, {
			name: "add chain 1 request 1",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[0].ChainID,
				Creator: addr1,
				Address: addr1,
			},
			wantID: 0,
		}, {
			name: "add chain 1 request 2",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[1].ChainID,
				Creator: addr2,
				Address: addr2,
			},
			wantID: 0,
		}, {
			name: "add chain 1 request 3",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[1].ChainID,
				Creator: addr2,
				Address: addr2,
			},
			wantID: 1,
		}, {
			name: "add chain 2 request 1",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[2].ChainID,
				Creator: addr3,
				Address: addr3,
			},
			wantID: 0,
		}, {
			name: "remove chain 2 request 2",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[2].ChainID,
				Creator: addr3,
				Address: addr3,
			},
			wantID: 1,
		}, {
			name: "remove coordinator account",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[2].ChainID,
				Creator: pk.GetCoordinatorAddressFromID(sdkCtx, chains[2].CoordinatorID),
				Address: addr4,
			},
			wantApprove: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.RequestRemoveAccount(ctx, &tt.msg)
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

				content, err := request.UnpackAccountRemoval(cdc)
				require.NoError(t, err)
				require.Equal(t, tt.msg.Address, content.Address)
			} else {
				_, foundGenesis := k.GetGenesisAccount(sdkCtx, tt.msg.ChainID, tt.msg.Address)
				require.False(t, foundGenesis, "genesis account not removed")
				_, foundVested := k.GetVestedAccount(sdkCtx, tt.msg.ChainID, tt.msg.Address)
				require.False(t, foundVested, "vested account not removed")
			}
		})
	}
}
