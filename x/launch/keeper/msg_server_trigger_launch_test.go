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
	k, _, _, srv, profileSrv, _, sdkCtx := setupMsgServer(t)

	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.Address()
	coordAddress2 := sample.Address()
	disableCoordAddress := sample.Address()
	coordNoExist := sample.Address()
	chainIDNoExist := uint64(1000)

	launchTimeTooLow := types.DefaultMinLaunchTime - 1
	launchTimeTooHigh := types.DefaultMaxLaunchTime + 1

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	msgCreateCoordinator = sample.MsgCreateCoordinator(disableCoordAddress)
	_, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create chains
	msgCreateChain := sample.MsgCreateChain(coordAddress, "", false, 0)
	res, err := srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID := res.LaunchID

	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID2 := res.LaunchID

	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	alreadyLaunched := res.LaunchID

	// Create chains
	msgCreateChain = sample.MsgCreateChain(disableCoordAddress, "", false, 0)
	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	disableChainID := res.LaunchID

	msgDisableCoord := sample.MsgDisableCoordinator(disableCoordAddress)
	_, err = profileSrv.DisableCoordinator(ctx, &msgDisableCoord)
	require.NoError(t, err)

	// Set a chain as already launched
	chain, found := k.GetChain(sdkCtx, alreadyLaunched)
	require.True(t, found)
	chain.LaunchTriggered = true
	k.SetChain(sdkCtx, chain)

	for _, tc := range []struct {
		name string
		msg  types.MsgTriggerLaunch
		err  error
	}{
		{
			name: "launch chain not launched",
			msg:  sample.MsgTriggerLaunch(coordAddress, chainID),
		},
		{
			name: "non existent chain id",
			msg:  sample.MsgTriggerLaunch(coordAddress, chainIDNoExist),
			err:  types.ErrChainNotFound,
		},
		{
			name: "non existent coordinator",
			msg:  sample.MsgTriggerLaunch(coordNoExist, chainID2),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg:  sample.MsgTriggerLaunch(coordAddress2, chainID2),
			err:  profiletypes.ErrCoordInvalid,
		},
		{
			name: "chain launch already triggered",
			msg:  sample.MsgTriggerLaunch(coordAddress, alreadyLaunched),
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
		{
			name: "disable coordinator",
			msg:  sample.MsgTriggerLaunch(disableCoordAddress, disableChainID),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Send the message
			_, err := srv.TriggerLaunch(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// Check value
			chain, found := k.GetChain(sdkCtx, tc.msg.LaunchID)
			require.True(t, found)
			require.True(t, chain.LaunchTriggered)
			require.EqualValues(t, testkeeper.ExampleTimestamp.Unix()+int64(tc.msg.RemainingTime), chain.LaunchTimestamp)
		})
	}
}
