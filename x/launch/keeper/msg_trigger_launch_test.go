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

func TestMsgTriggerLaunch(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)

	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.Address(r)
	coordAddress2 := sample.Address(r)
	coordNoExist := sample.Address(r)
	chainIDNoExist := uint64(1000)

	launchTimeTooLow := types.DefaultMinLaunchTime - 1
	launchTimeTooHigh := types.DefaultMaxLaunchTime + 1

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create chains
	msgCreateChain := sample.MsgCreateChain(r, coordAddress, "", false, 0)
	res, err := ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID := res.LaunchID

	res, err = ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID2 := res.LaunchID

	res, err = ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	alreadyLaunched := res.LaunchID

	// Set a chain as already launched
	chain, found := tk.LaunchKeeper.GetChain(sdkCtx, alreadyLaunched)
	require.True(t, found)
	chain.LaunchTriggered = true
	tk.LaunchKeeper.SetChain(sdkCtx, chain)

	for _, tc := range []struct {
		name string
		msg  types.MsgTriggerLaunch
		err  error
	}{
		{
			name: "launch chain not launched",
			msg:  sample.MsgTriggerLaunch(r, coordAddress, chainID),
		},
		{
			name: "non existent chain id",
			msg:  sample.MsgTriggerLaunch(r, coordAddress, chainIDNoExist),
			err:  types.ErrChainNotFound,
		},
		{
			name: "non existent coordinator",
			msg:  sample.MsgTriggerLaunch(r, coordNoExist, chainID2),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg:  sample.MsgTriggerLaunch(r, coordAddress2, chainID2),
			err:  profiletypes.ErrCoordInvalid,
		},
		{
			name: "chain launch already triggered",
			msg:  sample.MsgTriggerLaunch(r, coordAddress, alreadyLaunched),
			err:  types.ErrTriggeredLaunch,
		},
		{
			name: "launch time too low",
			msg:  *types.NewMsgTriggerLaunch(coordAddress, chainID2, launchTimeTooLow),
			err:  types.ErrLaunchTimeTooLow,
		},
		{
			name: "launch time too high",
			msg:  *types.NewMsgTriggerLaunch(coordAddress, chainID2, launchTimeTooHigh),
			err:  types.ErrLaunchTimeTooHigh,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Send the message
			_, err := ts.LaunchSrv.TriggerLaunch(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// Check value
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, tc.msg.LaunchID)
			require.True(t, found)
			require.True(t, chain.LaunchTriggered)
			require.EqualValues(t, testkeeper.ExampleTimestamp.Unix()+tc.msg.RemainingTime, chain.LaunchTimestamp)
			require.EqualValues(t, testkeeper.ExampleHeight, chain.ConsumerRevisionHeight)
		})
	}
}
