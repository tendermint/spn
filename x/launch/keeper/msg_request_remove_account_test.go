package keeper_test

import (
	"testing"

	testkeeper "github.com/tendermint/spn/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestRemoveAccount(t *testing.T) {
	var (
		invalidChain     = uint64(1000)
		coordAddr        = sample.Address(r)
		coordDisableAddr = sample.Address(r)
		addr1            = sample.Address(r)
		addr2            = sample.Address(r)
		addr3            = sample.Address(r)
		addr4            = sample.Address(r)
		sdkCtx, tk, ts   = testkeeper.NewTestSetup(t)
		ctx              = sdk.WrapSDKContext(sdkCtx)
	)

	coordID := tk.ProfileKeeper.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordAddr,
		Active:  true,
	})
	chains := createNChainForCoordinator(tk.LaunchKeeper, sdkCtx, coordID, 6)
	chains[0].LaunchTriggered = true
	tk.LaunchKeeper.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	tk.LaunchKeeper.SetChain(sdkCtx, chains[1])
	chains[5].IsMainnet = true
	chains[5].HasCampaign = true
	tk.LaunchKeeper.SetChain(sdkCtx, chains[5])

	tk.LaunchKeeper.SetVestingAccount(sdkCtx, types.VestingAccount{LaunchID: chains[3].LaunchID, Address: addr1})
	tk.LaunchKeeper.SetVestingAccount(sdkCtx, types.VestingAccount{LaunchID: chains[4].LaunchID, Address: addr2})
	tk.LaunchKeeper.SetVestingAccount(sdkCtx, types.VestingAccount{LaunchID: chains[4].LaunchID, Address: addr4})

	coordDisableID := tk.ProfileKeeper.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordDisableAddr,
		Active:  false,
	})
	disableChain := createNChainForCoordinator(tk.LaunchKeeper, sdkCtx, coordDisableID, 1)

	tests := []struct {
		name        string
		msg         types.MsgRequestRemoveAccount
		wantID      uint64
		wantApprove bool
		err         error
	}{
		{
			name: "should prevent requesting account removal for a non existing chain",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: invalidChain,
				Creator:  addr1,
				Address:  addr1,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "should prevent requesting account removal for a launch triggered chain",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[0].LaunchID,
				Creator:  addr1,
				Address:  addr1,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "should prevent requesting account removal for a chain where coordinator not found",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[1].LaunchID,
				Creator:  addr1,
				Address:  addr1,
			},
			err: types.ErrChainInactive,
		},
		{
			name: "should prevent requesting account removal without address permission",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[2].LaunchID,
				Creator:  addr1,
				Address:  addr3,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "should allow requesting account removal to an existing chain",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[2].LaunchID,
				Creator:  addr1,
				Address:  addr1,
			},
			wantID: 1,
		},
		{
			name: "should allow requesting account removal to a second chain",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[3].LaunchID,
				Creator:  coordAddr,
				Address:  addr1,
			},
			wantApprove: true,
			wantID:      1,
		},
		{
			name: "should allow requesting second account removal to a second chain",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[3].LaunchID,
				Creator:  addr2,
				Address:  addr2,
			},
			wantID: 2,
		},
		{
			name: "should allow requesting account removal to a third chain",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[4].LaunchID,
				Creator:  addr1,
				Address:  addr1,
			},
			wantID: 1,
		},
		{
			name: "should allow requesting and approving an account removal from the coordinator",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[4].LaunchID,
				Creator:  coordAddr,
				Address:  addr4,
			},
			wantApprove: true,
			wantID:      2,
		},
		{
			name: "should prevent requesting account removal from coordinator if no account to remove",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[4].LaunchID,
				Creator:  coordAddr,
				Address:  addr4,
			},
			err: types.ErrAccountNotFound,
		},
		{
			name: "should prevent requesting account removal for a mainnet chain",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: chains[5].LaunchID,
				Creator:  coordAddr,
				Address:  addr1,
			},
			err: types.ErrRemoveMainnetAccount,
		},
		{
			name: "should prevent requesting account removal for a chain where the coordinator of the chain is disabled",
			msg: types.MsgRequestRemoveAccount{
				LaunchID: disableChain[0].LaunchID,
				Creator:  sample.Address(r),
				Address:  sample.Address(r),
			},
			err: profiletypes.ErrCoordInactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.LaunchSrv.RequestRemoveAccount(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			request, found := tk.LaunchKeeper.GetRequest(sdkCtx, tt.msg.LaunchID, got.RequestID)
			require.True(t, found, "request not found")
			require.Equal(t, tt.wantID, request.RequestID)
			content := request.Content.GetAccountRemoval()
			require.NotNil(t, content)
			require.Equal(t, tt.msg.Address, content.Address)

			if !tt.wantApprove {
				require.Equal(t, types.Request_PENDING, request.Status)
			} else {
				_, foundGenesis := tk.LaunchKeeper.GetGenesisAccount(sdkCtx, tt.msg.LaunchID, tt.msg.Address)
				require.False(t, foundGenesis, "genesis account not removed")
				_, foundVesting := tk.LaunchKeeper.GetVestingAccount(sdkCtx, tt.msg.LaunchID, tt.msg.Address)
				require.False(t, foundVesting, "vesting account not removed")
				require.Equal(t, types.Request_APPROVED, request.Status)
			}
		})
	}
}
