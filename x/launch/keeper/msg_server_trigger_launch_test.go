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

	// TODO check params
	launchRemainingTime := uint64(1000)

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	if err != nil {
		t.Fatal(err)
	}
	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	if err != nil {
		t.Fatal(err)
	}

	// Create chains
	msgCreateChain := sample.MsgCreateChain(coordAddress, "foo", "")
	res, err := srv.CreateChain(ctx, &msgCreateChain)
	if err != nil {
		t.Fatal(err)
	}
	chainID := res.ChainID
	res, err = srv.CreateChain(ctx, &msgCreateChain)
	if err != nil {
		t.Fatal(err)
	}
	alreadyLaunched := res.ChainID

	// Set a chain as already launched
	chain, found := k.GetChain(sdkCtx, alreadyLaunched)
	if !found {
		t.Fatal(err)
	}
	chain.LaunchTriggered = true
	k.SetChain(sdkCtx, chain)

	for _, tc := range []struct{
		name  string
		msg   types.MsgTriggerLaunch
		valid bool
	} {
		{
			name: "launch chain not launched",
			msg: *types.NewMsgTriggerLaunch(coordAddress, chainID, launchRemainingTime),
			valid: true,
		},
		{
			name: "non existent chain id",
			msg: *types.NewMsgTriggerLaunch(coordAddress, chainIDNoExist, launchRemainingTime),
			valid: false,
		},
		{
			name: "non existent coordinator",
			msg: *types.NewMsgTriggerLaunch(coordNoExist, chainID, launchRemainingTime),
			valid: false,
		},
		{
			name: "invalid coordinator",
			msg: *types.NewMsgTriggerLaunch(coordAddress2, chainID, launchRemainingTime),
			valid: false,
		},
		{
			name: "chain launch already triggered",
			msg: *types.NewMsgTriggerLaunch(coordAddress, alreadyLaunched, launchRemainingTime),
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
		require.EqualValues(t, sampleTimestamp.Unix() + int64(tc.msg.RemainingTime), chain.LaunchTimestamp)
	}
}