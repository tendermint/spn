package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgRequestAddValidator(t *testing.T) {
	var (
		invalidChain, _           = sample.ChainID(0)
		addr1                     = sample.AccAddress()
		addr2                     = sample.AccAddress()
		k, _, srv, _, sdkCtx, cdc = setupMsgServer(t)
		ctx                       = sdk.WrapSDKContext(sdkCtx)
		chains                    = createNChain(k, sdkCtx, 3)
	)
	chains[2].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[2])

	for _, tc := range []struct {
		name string
		msg  types.MsgRequestAddValidator
		want uint64
		valid  bool
	} {
		{
			name: "invalid chain",
			msg: sample.MsgRequestAddValidator(addr1, invalidChain),
			valid: false,
		},
		{
			name: "chain with triggered launch",
			msg: sample.MsgRequestAddValidator(addr1, chains[2].ChainID),
			valid: false,
		},
		{
			name: "request to a chain 1",
			msg: sample.MsgRequestAddValidator(addr1, chains[0].ChainID),
			valid: true,
			want: uint64(0),
		},
		{
			name: "second request to a chain 1",
			msg: sample.MsgRequestAddValidator(addr2, chains[0].ChainID),
			valid: true,
			want: uint64(1),
		},
		{
			name: "request to a chain 2",
			msg: sample.MsgRequestAddValidator(addr1, chains[1].ChainID),
			valid: true,
			want: uint64(0),
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := srv.RequestAddValidator(ctx, &tc.msg)
			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			request, found := k.GetRequest(sdkCtx, tc.msg.ChainID, got.RequestID)
			require.True(t, found, "request not found")
			require.Equal(t, tc.want, request.RequestID)

			content, err := request.UnpackGenesisValidator(cdc)
			require.NoError(t, err)
			require.Equal(t, tc.msg.ValAddress, content.Address)
			require.Equal(t, tc.msg.ChainID, content.ChainID)
			require.True(t, tc.msg.SelfDelegation.Equal(content.SelfDelegation))
			require.Equal(t, tc.msg.GenTx, content.GenTx)
			require.Equal(t, tc.msg.Peer, content.Peer)
			require.Equal(t, tc.msg.ConsPubKey, content.ConsPubKey)
		})
	}
}