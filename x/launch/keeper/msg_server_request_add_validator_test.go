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
		invalidChain             = uint64(1000)
		coordAddr                = sample.AccAddress()
		addr1                    = sample.AccAddress()
		addr2                    = sample.AccAddress()
		k, _, pk, srv, _, sdkCtx = setupMsgServer(t)
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
		err         error
	}{
		{
			name: "invalid chain",
			msg:  sample.MsgRequestAddValidator(addr1, invalidChain),
			err:  types.ErrChainNotFound,
		},
		{
			name: "chain with triggered launch",
			msg:  sample.MsgRequestAddValidator(addr1, chains[0].Id),
			err:  types.ErrTriggeredLaunch,
		},
		{
			name: "chain without coordinator",
			msg:  sample.MsgRequestAddValidator(addr1, chains[1].Id),
			err:  types.ErrChainInactive,
		},
		{
			name:   "request to a chain 3",
			msg:    sample.MsgRequestAddValidator(addr1, chains[2].Id),
			wantID: 0,
		},
		{
			name:   "second request to a chain 3",
			msg:    sample.MsgRequestAddValidator(addr2, chains[2].Id),
			wantID: 1,
		},
		{
			name:   "request to a chain 4",
			msg:    sample.MsgRequestAddValidator(addr1, chains[3].Id),
			wantID: 0,
		},
		{
			name:        "request from coordinator is pre-approved",
			msg:         sample.MsgRequestAddValidator(coordAddr, chains[3].Id),
			wantApprove: true,
		},
		{
			name:        "failing request from coordinator",
			msg:         sample.MsgRequestAddValidator(coordAddr, chains[3].Id),
			err:         types.ErrValidatorAlreadyExist,
			wantApprove: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := srv.RequestAddValidator(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, tc.err, err)
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
