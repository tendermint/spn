package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgRevertLaunch(t *testing.T) {
	k, _, srv, profileSrv, sdkCtx, _ := setupMsgServer(t)

	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.AccAddress()
	coordAddress2 := sample.AccAddress()
	coordNoExist := sample.AccAddress()
	chainIDNoExist := uint64(1000)

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create chains
	msgCreateChain := sample.MsgCreateChain(coordAddress, "")
	res, err := srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	notLaunched := res.Id

	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	delayNotReached := res.Id
	chain, found := k.GetChain(sdkCtx, delayNotReached)
	require.True(t, found)
	chain.LaunchTriggered = true
	chain.LaunchTimestamp = testkeeper.ExampleTimestamp.Unix() - types.RevertDelay + 1
	k.SetChain(sdkCtx, chain)

	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	delayReached := res.Id
	chain, found = k.GetChain(sdkCtx, delayReached)
	require.True(t, found)
	chain.LaunchTriggered = true
	chain.LaunchTimestamp = testkeeper.ExampleTimestamp.Unix() - types.RevertDelay
	k.SetChain(sdkCtx, chain)

	for _, tc := range []struct {
		name string
		msg  types.MsgRevertLaunch
		err  error
	}{
		{
			name: "revert delay reached",
			msg:  *types.NewMsgRevertLaunch(coordAddress, delayReached),
		},
		{
			name: "revert delay not reached",
			msg:  *types.NewMsgRevertLaunch(coordAddress, delayNotReached),
			err:  types.ErrRevertDelayNotReached,
		},
		{
			name: "launch chain not launched",
			msg:  *types.NewMsgRevertLaunch(coordAddress, notLaunched),
			err:  types.ErrNotTriggeredLaunch,
		},
		{
			name: "non existent coordinator",
			msg:  *types.NewMsgRevertLaunch(coordNoExist, delayReached),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg:  *types.NewMsgRevertLaunch(coordAddress2, delayReached),
			err:  profiletypes.ErrCoordInvalid,
		},
		{
			name: "non existent chain id",
			msg:  *types.NewMsgRevertLaunch(coordAddress, chainIDNoExist),
			err:  types.ErrChainNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Send the message
			_, err := srv.RevertLaunch(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// Check value
			chain, found := k.GetChain(sdkCtx, tc.msg.ChainID)
			require.True(t, found)
			require.False(t, chain.LaunchTriggered)
			require.EqualValues(t, int64(0), chain.LaunchTimestamp)
		})
	}
}
