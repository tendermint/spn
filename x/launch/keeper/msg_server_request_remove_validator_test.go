package keeper

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRequestRemoveValidator(t *testing.T) {
	var (
		invalidChain, _            = sample.ChainID(0)
		coordAddr                  = sample.AccAddress()
		addr1                      = sample.AccAddress()
		addr2                      = sample.AccAddress()
		addr3                      = sample.AccAddress()
		addr4                      = sample.AccAddress()
		k, pk, srv, _, sdkCtx, cdc = setupMsgServer(t)
		ctx                        = sdk.WrapSDKContext(sdkCtx)
	)

	coordID := pk.AppendCoordinator(sdkCtx, profiletypes.Coordinator{
		Address: coordAddr,
	})
	chains := createNChainForCoordinator(k, sdkCtx, coordID, 4)
	chains[3].LaunchTriggered = true
	k.SetChain(sdkCtx, chains[3])
	k.SetGenesisValidator(sdkCtx, types.GenesisValidator{ChainID: chains[1].ChainID, Address: addr1})
	k.SetGenesisValidator(sdkCtx, types.GenesisValidator{ChainID: chains[2].ChainID, Address: addr2})
	k.SetGenesisValidator(sdkCtx, types.GenesisValidator{ChainID: chains[2].ChainID, Address: addr4})

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
				ChainID:          invalidChain,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			err: sdkerrors.Wrap(types.ErrChainNotFound, invalidChain),
		}, {
			name: "launch triggered chain",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[3].ChainID,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			err: sdkerrors.Wrap(types.ErrTriggeredLaunch, chains[3].ChainID),
		}, {
			name: "no permission error",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[0].ChainID,
				Creator:          addr1,
				ValidatorAddress: addr3,
			},
			err: sdkerrors.Wrap(types.ErrNoAddressPermission, addr1),
		}, {
			name: "add chain 1 request 1",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[0].ChainID,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			wantID: 0,
		}, {
			name: "add chain 2 request 1",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[1].ChainID,
				Creator:          coordAddr,
				ValidatorAddress: addr1,
			},
			wantApprove: true,
		}, {
			name: "add chain 2 request 2",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[1].ChainID,
				Creator:          addr2,
				ValidatorAddress: addr2,
			},
			wantID: 0,
		}, {
			name: "add chain 3 request 1",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[2].ChainID,
				Creator:          addr1,
				ValidatorAddress: addr1,
			},
			wantID: 0,
		}, {
			name: "add chain 3 request 2",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[2].ChainID,
				Creator:          coordAddr,
				ValidatorAddress: addr2,
			},
			wantApprove: true,
		}, {
			name: "add chain 3 request 3",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[2].ChainID,
				Creator:          addr3,
				ValidatorAddress: addr3,
			},
			wantID: 1,
		}, {
			name: "add coordinator validator",
			msg: types.MsgRequestRemoveValidator{
				ChainID:          chains[2].ChainID,
				Creator:          coordAddr,
				ValidatorAddress: addr4,
			},
			wantApprove: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.RequestRemoveValidator(ctx, &tt.msg)
			if tt.err != nil {
				require.Error(t, err)
				require.Equal(t, tt.err.Error(), err.Error())
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.wantID, got.RequestID)
			require.Equal(t, tt.wantApprove, got.AutoApproved)

			if !tt.wantApprove {
				request, found := k.GetRequest(sdkCtx, tt.msg.ChainID, got.RequestID)
				require.True(t, found, "request not found")
				require.Equal(t, tt.wantID, request.RequestID)

				content, err := request.UnpackValidatorRemoval(cdc)
				require.NoError(t, err)
				require.Equal(t, tt.msg.ValidatorAddress, content.ValAddress)
			} else {
				_, found := k.GetGenesisValidator(sdkCtx, tt.msg.ChainID, tt.msg.ValidatorAddress)
				require.False(t, found, "genesis validator not removed")
			}
		})
	}
}
