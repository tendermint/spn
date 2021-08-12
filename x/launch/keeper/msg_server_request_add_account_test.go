package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddAccount(t *testing.T) {
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
		msg  types.MsgRequestAddAccount
		want uint64
		err  error
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
			want: 0,
		}, {
			name: "add chain 1 request 2",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[1].ChainID,
				Address: addr2,
				Coins:   sample.Coins(),
			},
			want: 0,
		}, {
			name: "add chain 1 request 3",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[1].ChainID,
				Address: addr2,
				Coins:   sample.Coins(),
			},
			want: 1,
		}, {
			name: "add chain 2 request 1",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[2].ChainID,
				Address: addr3,
				Coins:   sample.Coins(),
			},
			want: 0,
		}, {
			name: "add chain 2 request 2",
			msg: types.MsgRequestAddAccount{
				ChainID: chains[2].ChainID,
				Address: addr3,
				Coins:   sample.Coins(),
			},
			want: 1,
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

			request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, got.RequestID)
			require.True(t, found, "request not found")
			require.Equal(t, tt.want, request.RequestID)

			content, err := request.UnpackGenesisAccount(cdc)
			require.NoError(t, err)
			require.Equal(t, tt.msg.Address, content.Address)
			require.Equal(t, tt.msg.ChainID, content.ChainID)
			require.Equal(t, tt.msg.Coins, content.Coins)
		})
	}
}
