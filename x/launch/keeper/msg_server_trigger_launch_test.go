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
	coordAddress := sample.AccAddress()
	coordAddress2 := sample.AccAddress()
	coordNoExist := sample.AccAddress()
	chainIDNoExist := uint64(1000)

	launchTime := types.DefaultMinLaunchTime
	launchTimeTooLow := types.DefaultMinLaunchTime - 1
	launchTimeTooHigh := types.DefaultMaxLaunchTime + 1

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create chains
	msgCreateChain := sample.MsgCreateChain(coordAddress, "", 0)
	res, err := srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID := res.Id

	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID2 := res.Id

	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	alreadyLaunched := res.Id

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
			msg:  *types.NewMsgTriggerLaunch(coordAddress, chainID, launchTime),
		},
		{
			name: "non existent chain id",
			msg:  *types.NewMsgTriggerLaunch(coordAddress, chainIDNoExist, launchTime),
			err:  types.ErrChainNotFound,
		},
		{
			name: "non existent coordinator",
			msg:  *types.NewMsgTriggerLaunch(coordNoExist, chainID2, launchTime),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg:  *types.NewMsgTriggerLaunch(coordAddress2, chainID2, launchTime),
			err:  profiletypes.ErrCoordInvalid,
		},
		{
			name: "chain launch already triggered",
			msg:  *types.NewMsgTriggerLaunch(coordAddress, alreadyLaunched, launchTime),
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
			_, err := srv.TriggerLaunch(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// Check value
			chain, found := k.GetChain(sdkCtx, tc.msg.ChainID)
			require.True(t, found)
			require.True(t, chain.LaunchTriggered)
			require.EqualValues(t, testkeeper.ExampleTimestamp.Unix()+int64(tc.msg.RemainingTime), chain.LaunchTimestamp)
		})
	}
}
