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
		invalidChain, _           = sample.ChainID(0)
		addr1                     = sample.AccAddress()
		addr2                     = sample.AccAddress()
		addr3                     = sample.AccAddress()
		k, _, srv, _, sdkCtx, cdc = setupMsgServer(t)
		ctx                       = sdk.WrapSDKContext(sdkCtx)
		chains                    = createNChain(k, sdkCtx, 4)
	)
	chains[3].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[3])
	tests := []struct {
		name string
		msg  types.MsgRequestRemoveAccount
		want uint64
		err  error
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
			want: 0,
		}, {
			name: "add chain 1 request 2",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[1].ChainID,
				Creator: addr2,
				Address: addr2,
			},
			want: 0,
		}, {
			name: "add chain 1 request 3",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[1].ChainID,
				Creator: addr2,
				Address: addr2,
			},
			want: 1,
		}, {
			name: "add chain 2 request 1",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[2].ChainID,
				Creator: addr3,
				Address: addr3,
			},
			want: 0,
		}, {
			name: "add chain 2 request 2",
			msg: types.MsgRequestRemoveAccount{
				ChainID: chains[2].ChainID,
				Creator: addr3,
				Address: addr3,
			},
			want: 1,
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

			request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, got.RequestID)
			require.True(t, found, "request not found")
			require.Equal(t, tt.want, request.RequestID)

			content, err := request.UnpackAccountRemoval(cdc)
			require.NoError(t, err)
			require.Equal(t, tt.msg.Address, content.Address)
		})
	}
}
