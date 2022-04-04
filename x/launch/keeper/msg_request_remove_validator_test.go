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

func TestMsgRequestRemoveValidator(t *testing.T) {
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
	chains := createNChainForCoordinator(tk.LaunchKeeper, sdkCtx, coordID, 5)
	chains[0].LaunchTriggered = true
	tk.LaunchKeeper.SetChain(sdkCtx, chains[0])
	chains[1].CoordinatorID = 99999
	tk.LaunchKeeper.SetChain(sdkCtx, chains[1])

	coordDisableID := tk.ProfileKeeper.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordDisableAddr,
		Active:  false,
	})
	disableChain := createNChainForCoordinator(tk.LaunchKeeper, sdkCtx, coordDisableID, 1)

	tk.LaunchKeeper.SetGenesisValidator(sdkCtx, types.GenesisValidator{LaunchID: chains[3].LaunchID, Address: addr1})
	tk.LaunchKeeper.SetGenesisValidator(sdkCtx, types.GenesisValidator{LaunchID: chains[4].LaunchID, Address: addr2})
	tk.LaunchKeeper.SetGenesisValidator(sdkCtx, types.GenesisValidator{LaunchID: chains[4].LaunchID, Address: addr4})

	tests := []struct {
		name        string
		msg         types.MsgRequestRemoveValidator
		wantID      uint64
		wantApprove bool
		err         error
	}{
		{
			name: "invalid chain",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         invalidChain,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			err: types.ErrChainNotFound,
		},
		{
			name: "launch triggered chain",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[0].LaunchID,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			err: types.ErrTriggeredLaunch,
		},
		{
			name: "coordinator not found",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[1].LaunchID,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			err: types.ErrChainInactive,
		},
		{
			name: "no permission error",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[2].LaunchID,
				Creator:          addr1,
				ValidatorAddress: addr3,
			},
			err: types.ErrNoAddressPermission,
		},
		{
			name: "add chain 3 request 1",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[2].LaunchID,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			wantID: 1,
		},
		{
			name: "add chain 3 request 1",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[3].LaunchID,
				Creator:          coordAddr,
				ValidatorAddress: addr1,
			},
			wantApprove: true,
		},
		{
			name: "add chain 4 request 2",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[3].LaunchID,
				Creator:          addr2,
				ValidatorAddress: addr2,
			},
			wantID: 1,
		},
		{
			name: "add chain 5 request 1",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[4].LaunchID,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			wantID: 1,
		},
		{
			name: "add chain 5 request 2",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[4].LaunchID,
				Creator:          coordAddr,
				ValidatorAddress: addr2,
			},
			wantApprove: true,
		},
		{
			name: "add chain 5 request 3",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[4].LaunchID,
				Creator:          addr3,
				ValidatorAddress: addr3,
			},
			wantID: 2,
		},
		{
			name: "request from coordinator is pre-approved",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[4].LaunchID,
				Creator:          coordAddr,
				ValidatorAddress: addr4,
			},
			wantApprove: true,
		},
		{
			name: "failing request from coordinator",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         chains[4].LaunchID,
				Creator:          coordAddr,
				ValidatorAddress: addr4,
			},
			err: types.ErrValidatorNotFound,
		},
		{
			name: "fail if the coordinator of the chain is disabled",
			msg: types.MsgRequestRemoveValidator{
				LaunchID:         disableChain[0].LaunchID,
				Creator:          sample.Address(r),
				ValidatorAddress: sample.Address(r),
			}, err: profiletypes.ErrCoordInactive,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ts.LaunchSrv.RequestRemoveValidator(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			if !tt.wantApprove {
				request, found := tk.LaunchKeeper.GetRequest(sdkCtx, tt.msg.LaunchID, got.RequestID)
				require.True(t, found, "request not found")
				require.Equal(t, tt.wantID, request.RequestID)
				require.Equal(t, types.Request_PENDING, request.Status)

				content := request.Content.GetValidatorRemoval()
				require.NotNil(t, content)
				require.Equal(t, tt.msg.ValidatorAddress, content.ValAddress)
			} else {
				_, found := tk.LaunchKeeper.GetGenesisValidator(sdkCtx, tt.msg.LaunchID, tt.msg.ValidatorAddress)
				require.False(t, found, "genesis validator not removed")
			}
		})
	}
}
