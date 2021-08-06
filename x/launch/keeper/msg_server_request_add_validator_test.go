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
		addr3                     = sample.AccAddress()
		k, _, srv, _, sdkCtx, cdc = setupMsgServer(t)
		ctx                       = sdk.WrapSDKContext(sdkCtx)
		chains                    = createNChain(k, sdkCtx, 4)
	)
	chains[3].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[3])

	for _, tc := range []struct {
		name string
		msg  types.MsgRequestAddValidator
		want uint64
		valid  bool
	} {

	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := srv.RequestAddValidator(ctx, &tc.msg)
			if tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			request, found := k.GetRequest(sdkCtx, tc.msg.ChainID, got.RequestID)
			require.True(t, found, "request not found")
			require.Equal(t, tc.want, request.RequestID)

			content, err := request.UnpackGenesisAccount(cdc)
			require.NoError(t, err)
			require.Equal(t, tc.msg.Address, content.Address)
			require.Equal(t, tc.msg.ChainID, content.ChainID)
			require.Equal(t, tc.msg.Coins, content.Coins)
		})
	}
}