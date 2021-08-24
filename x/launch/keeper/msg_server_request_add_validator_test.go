package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestAddValidator(t *testing.T) {
	var (
		invalidChain, _          = sample.ChainID(0)
		coordAddr                = sample.AccAddress()
		addr1                    = sample.AccAddress()
		addr2                    = sample.AccAddress()
		k, pk, srv, _, sdkCtx, _ = setupMsgServer(t)
		ctx                      = sdk.WrapSDKContext(sdkCtx)
	)

	coordID := pk.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordAddr,
	})
	chains := createNChainForCoordinator(k, sdkCtx, coordID, 4)
	chains[0].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	k.SetChain(sdkCtx, chains[1])

	for _, tc := range []struct {
		name        string
		msg         types.MsgRequestAddValidator
		wantID      uint64
		wantApprove bool
		valid       bool
	}{
		{
			name:  "invalid chain",
			msg:   sample.MsgRequestAddValidator(addr1, invalidChain),
			valid: false,
		}, {
			name:  "chain with triggered launch",
			msg:   sample.MsgRequestAddValidator(addr1, chains[0].ChainID),
			valid: false,
		}, {
			name:  "chain without coordinator",
			msg:   sample.MsgRequestAddValidator(addr1, chains[1].ChainID),
			valid: false,
		}, {
			name:   "request to a chain 3",
			msg:    sample.MsgRequestAddValidator(addr1, chains[2].ChainID),
			valid:  true,
			wantID: 0,
		}, {
			name:   "second request to a chain 3",
			msg:    sample.MsgRequestAddValidator(addr2, chains[2].ChainID),
			valid:  true,
			wantID: 1,
		}, {
			name:   "request to a chain 4",
			msg:    sample.MsgRequestAddValidator(addr1, chains[3].ChainID),
			valid:  true,
			wantID: 0,
		}, {
			name:        "add coordinator to a chain",
			msg:         sample.MsgRequestAddValidator(coordAddr, chains[3].ChainID),
			valid:       true,
			wantApprove: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := srv.RequestAddValidator(ctx, &tc.msg)
			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.wantID, got.RequestID)
			require.Equal(t, tc.wantApprove, got.AutoApproved)

			if !tc.wantApprove {
				request, found := k.GetRequest(sdkCtx, tc.msg.ChainID, got.RequestID)
				require.True(t, found, "request not found")
				require.Equal(t, tc.wantID, request.RequestID)

				content := request.Content.GetGenesisValidator()
				require.NotNil(t, content)
				require.Equal(t, tc.msg.ValAddress, content.Address)
				require.Equal(t, tc.msg.ChainID, content.ChainID)
				require.True(t, tc.msg.SelfDelegation.Equal(content.SelfDelegation))
				require.Equal(t, tc.msg.GenTx, content.GenTx)
				require.Equal(t, tc.msg.Peer, content.Peer)
				require.Equal(t, tc.msg.ConsPubKey, content.ConsPubKey)
			} else {
				_, found := k.GetGenesisValidator(sdkCtx, tc.msg.ChainID, tc.msg.ValAddress)
				require.True(t, found, "genesis validator not found")
			}
		})
	}
}
