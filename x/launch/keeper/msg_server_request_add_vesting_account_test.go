package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestAddVestingAccount(t *testing.T) {
	var (
		invalidChain                = uint64(1000)
		coordAddr                   = sample.Address()
		addr1                       = sample.Address()
		addr2                       = sample.Address()
		addr3                       = sample.Address()
		addr4                       = sample.Address()
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
	chains[5].HasCampaign = true
	k.SetChain(sdkCtx, chains[5])

	tests := []struct {
		name        string
		msg         types.MsgRequestAddVestingAccount
		wantID      uint64
		wantApprove bool
		err         error
	}{
		{
			name: "invalid chain",
			msg:  sample.MsgRequestAddVestingAccount(sample.Address(), addr1, invalidChain),
			err:  types.ErrChainNotFound,
		},
		{
			name: "launch triggered chain",
			msg:  sample.MsgRequestAddVestingAccount(sample.Address(), addr1, chains[0].LaunchID),
			err:  types.ErrTriggeredLaunch,
		},
		{
			name: "coordinator not found",
			msg:  sample.MsgRequestAddVestingAccount(sample.Address(), addr1, chains[1].LaunchID),
			err:  types.ErrChainInactive,
		},
		{
			name:   "add chain 3 request 1",
			msg:    sample.MsgRequestAddVestingAccount(sample.Address(), addr1, chains[2].LaunchID),
			wantID: 1,
		},
		{
			name:   "add chain 4 request 1",
			msg:    sample.MsgRequestAddVestingAccount(sample.Address(), addr1, chains[3].LaunchID),
			wantID: 1,
		},
		{
			name:   "add chain 4 request 2",
			msg:    sample.MsgRequestAddVestingAccount(sample.Address(), addr2, chains[3].LaunchID),
			wantID: 2,
		},
		{
			name:   "add chain 5 request 1",
			msg:    sample.MsgRequestAddVestingAccount(sample.Address(), addr1, chains[4].LaunchID),
			wantID: 1,
		},
		{
			name:   "add chain 5 request 2",
			msg:    sample.MsgRequestAddVestingAccount(sample.Address(), addr2, chains[4].LaunchID),
			wantID: 2,
		},
		{
			name:   "add chain 5 request 3",
			msg:    sample.MsgRequestAddVestingAccount(sample.Address(), addr3, chains[4].LaunchID),
			wantID: 3,
		},
		{
			name:        "request from coordinator is pre-approved",
			msg:         sample.MsgRequestAddVestingAccount(coordAddr, addr4, chains[4].LaunchID),
			wantApprove: true,
		},
		{
			name: "failing request from coordinator",
			msg:  sample.MsgRequestAddVestingAccount(coordAddr, addr4, chains[4].LaunchID),
			err:  types.ErrAccountAlreadyExist,
		},
		{
			name: "is mainnet chain",
			msg:  sample.MsgRequestAddVestingAccount(coordAddr, sample.Address(), chains[5].LaunchID),
			err:  types.ErrAddMainnetVestingAccount,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.RequestAddVestingAccount(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			if !tt.wantApprove {
				request, found := k.GetRequest(sdkCtx, tt.msg.LaunchID, got.RequestID)
				require.True(t, found, "request not found")
				require.Equal(t, tt.wantID, request.RequestID)
				require.Equal(t, tt.msg.Creator, request.Creator)

				content := request.Content.GetVestingAccount()
				require.NotNil(t, content)
				require.Equal(t, tt.msg.Address, content.Address)
				require.Equal(t, tt.msg.LaunchID, content.LaunchID)
				require.Equal(t, tt.msg.StartingBalance, content.StartingBalance)
				require.Equal(t, tt.msg.Options.String(), content.VestingOptions.String())
			} else {
				_, found := k.GetVestingAccount(sdkCtx, tt.msg.LaunchID, tt.msg.Address)
				require.True(t, found, "vesting account not found")
			}
		})
	}
}
