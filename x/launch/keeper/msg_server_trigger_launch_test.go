package keeper

import (
	"github.com/stretchr/testify/require"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
)

func TestMsgTriggerLaunch(t *testing.T) {
	k, _, srv, profileSrv, sdkCtx, _ := setupMsgServer(t)

	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.AccAddress()
	coordAddress2 := sample.AccAddress()
	coordNoExist := sample.AccAddress()
	chainIDNoExist, _ := sample.ChainID(0)

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
	msgCreateChain := sample.MsgCreateChain(coordAddress, "foo", "")
	res, err := srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)

	chainID := res.ChainID
	res, err = srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	alreadyLaunched := res.ChainID

	// Set a chain as already launched
	chain, found := k.GetChain(sdkCtx, alreadyLaunched)
	require.True(t, found)
	chain.LaunchTriggered = true
	k.SetChain(sdkCtx, chain)

	for _, tc := range []struct {
		name  string
		msg   types.MsgTriggerLaunch
		valid bool
	}{
		{
			name:  "launch chain not launched",
			msg:   *types.NewMsgTriggerLaunch(coordAddress, chainID, launchTime),
			valid: true,
		},
		{
			name:  "non existent chain id",
			msg:   *types.NewMsgTriggerLaunch(coordAddress, chainIDNoExist, launchTime),
			valid: false,
		},
		{
			name:  "non existent coordinator",
			msg:   *types.NewMsgTriggerLaunch(coordNoExist, chainID, launchTime),
			valid: false,
		},
		{
			name:  "invalid coordinator",
			msg:   *types.NewMsgTriggerLaunch(coordAddress2, chainID, launchTime),
			valid: false,
		},
		{
			name:  "chain launch already triggered",
			msg:   *types.NewMsgTriggerLaunch(coordAddress, alreadyLaunched, launchTime),
			valid: false,
		},
		{
			name:  "launch time too low",
			msg:   *types.NewMsgTriggerLaunch(coordAddress, chainID, launchTimeTooLow),
			valid: false,
		},
		{
			name:  "launch time too high",
			msg:   *types.NewMsgTriggerLaunch(coordAddress, chainID, launchTimeTooHigh),
			valid: false,
		},
	} {
		// Send the message
		_, err := srv.TriggerLaunch(ctx, &tc.msg)
		if !tc.valid {
			require.Error(t, err)
			return
		}
		require.NoError(t, err)

		// Check value
		chain, found := k.GetChain(sdkCtx, tc.msg.ChainID)
		require.True(t, found)
		require.True(t, chain.LaunchTriggered)
		require.EqualValues(t, sampleTimestamp.Unix()+int64(tc.msg.RemainingTime), chain.LaunchTimestamp)
	}
}
