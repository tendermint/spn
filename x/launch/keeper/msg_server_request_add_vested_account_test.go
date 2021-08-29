package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestAddVestedAccount(t *testing.T) {
	var (
		invalidChain             = uint64(1000)
		coordAddr                = sample.AccAddress()
		addr1                    = sample.AccAddress()
		addr2                    = sample.AccAddress()
		addr3                    = sample.AccAddress()
		k, pk, srv, _, sdkCtx, _ = setupMsgServer(t)
		ctx                      = sdk.WrapSDKContext(sdkCtx)
	)

	coordID := pk.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordAddr,
	})
	chains := createNChainForCoordinator(k, sdkCtx, coordID, 5)
	chains[0].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[0])
	delayedVesting := *types.NewDelayedVesting(sample.Coins(), 10000)
	chains[1].CoordinatorID = 99999
	k.SetChain(sdkCtx, chains[1])

	tests := []struct {
		name        string
		msg         types.MsgRequestAddVestedAccount
		wantID      uint64
		wantApprove bool
		err         error
	}{
		{
			name: "invalid chain",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         invalidChain,
				Address:         addr1,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			err: types.ErrChainNotFound,
		}, {
			name: "launch triggered chain",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[0].Id,
				Address:         addr1,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			err: types.ErrTriggeredLaunch,
		}, {
			name: "coordinator not found",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[1].Id,
				Address:         addr1,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			err: types.ErrChainInactive,
		}, {
			name: "add chain 3 request 1",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[2].Id,
				Address:         addr1,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			wantID: 0,
		}, {
			name: "add chain 4 request 1",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[3].Id,
				Address:         addr1,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			wantID: 0,
		}, {
			name: "add chain 4 request 2",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[3].Id,
				Address:         addr2,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			wantID: 1,
		}, {
			name: "add chain 5 request 1",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[4].Id,
				Address:         addr1,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			wantID: 0,
		}, {
			name: "add chain 5 request 2",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[4].Id,
				Address:         addr2,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			wantID: 1,
		}, {
			name: "add chain 5 request 3",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[4].Id,
				Address:         addr3,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			wantID: 2,
		}, {
			name: "add coordinator account",
			msg: types.MsgRequestAddVestedAccount{
				ChainID:         chains[4].Id,
				Address:         coordAddr,
				StartingBalance: sample.Coins(),
				Options:         delayedVesting,
			},
			wantApprove: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.RequestAddVestedAccount(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			if !tt.wantApprove {
				request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, got.RequestID)
				require.True(t, found, "request not found")
				require.Equal(t, tt.wantID, request.RequestID)

				content := request.Content.GetVestedAccount()
				require.NotNil(t, content)
				require.Equal(t, tt.msg.Address, content.Address)
				require.Equal(t, tt.msg.ChainID, content.ChainID)
				require.Equal(t, tt.msg.StartingBalance, content.StartingBalance)
				require.Equal(t, tt.msg.Options.String(), content.VestingOptions.String())
			} else {
				_, found := k.GetVestedAccount(sdkCtx, tt.msg.ChainID, tt.msg.Address)
				require.True(t, found, "vested account not found")
			}
		})
	}
}
