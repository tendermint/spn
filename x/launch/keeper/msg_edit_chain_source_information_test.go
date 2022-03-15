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

func TestMsgEditChainSourceInformation(t *testing.T) {
	sdkCtx, tk, ts := testkeeper.NewTestSetup(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.Address()
	coordAddress2 := sample.Address()
	coordNoExist := sample.Address()
	launchIDNoExist := uint64(1000)

	// Create coordinators
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	_, err := ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	msgCreateCoordinator = sample.MsgCreateCoordinator(coordAddress2)
	_, err = ts.ProfileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)

	// Create a chain
	msgCreateChain := sample.MsgCreateChain(coordAddress, "", false, 0)
	res, err := ts.LaunchSrv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	launchID := res.LaunchID

	for _, tc := range []struct {
		name string
		msg  types.MsgEditChainSourceInformation
		err  error
	}{
		{
			name: "edit genesis chain ID",
			msg: sample.MsgEditChainSourceInformation(coordAddress, launchID,
				true,
				false,
				false,
				false,
			),
		},
		{
			name: "edit source",
			msg: sample.MsgEditChainSourceInformation(coordAddress, launchID,
				false,
				true,
				false,
				false,
			),
		},
		{
			name: "edit initial genesis with default genesis",
			msg: sample.MsgEditChainSourceInformation(coordAddress, launchID,
				false,
				false,
				true,
				false,
			),
		},
		{
			name: "edit initial genesis with genesis url",
			msg: sample.MsgEditChainSourceInformation(coordAddress, launchID,
				false,
				false,
				true,
				true,
			),
		},
		{
			name: "edit source and initial genesis",
			msg: sample.MsgEditChainSourceInformation(coordAddress, launchID,
				false,
				true,
				true,
				true,
			),
		},
		{
			name: "non existent launch id",
			msg: sample.MsgEditChainSourceInformation(coordAddress, launchIDNoExist,
				false,
				true,
				false,
				false,
			),
			err: types.ErrChainNotFound,
		},
		{
			name: "non existent coordinator",
			msg: sample.MsgEditChainSourceInformation(coordNoExist, launchID,
				false,
				true,
				false,
				false,
			),
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg: sample.MsgEditChainSourceInformation(coordAddress2, launchID,
				false,
				true,
				false,
				false,
			),
			err: profiletypes.ErrCoordInvalid,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// Fetch the previous state of the chain to perform checks
			var previousChain types.Chain
			var found bool
			if tc.err == nil {
				previousChain, found = tk.LaunchKeeper.GetChain(sdkCtx, tc.msg.LaunchID)
				require.True(t, found)
			}

			// Send the message
			_, err := ts.LaunchSrv.EditChainSourceInformation(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// The chain must continue to exist in the store
			chain, found := tk.LaunchKeeper.GetChain(sdkCtx, tc.msg.LaunchID)
			require.True(t, found)

			// Unchanged values
			require.EqualValues(t, previousChain.CoordinatorID, chain.CoordinatorID)
			require.EqualValues(t, previousChain.CreatedAt, chain.CreatedAt)
			require.EqualValues(t, previousChain.LaunchTimestamp, chain.LaunchTimestamp)
			require.EqualValues(t, previousChain.LaunchTriggered, chain.LaunchTriggered)

			// Compare changed values
			if tc.msg.GenesisChainID != "" {
				require.EqualValues(t, tc.msg.GenesisChainID, chain.GenesisChainID)
			} else {
				require.EqualValues(t, previousChain.GenesisChainID, chain.GenesisChainID)
			}
			if tc.msg.SourceURL != "" {
				require.EqualValues(t, tc.msg.SourceURL, chain.SourceURL)
				require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)
			} else {
				require.EqualValues(t, previousChain.SourceURL, chain.SourceURL)
				require.EqualValues(t, previousChain.SourceHash, chain.SourceHash)
			}

			if tc.msg.InitialGenesis != nil {
				require.EqualValues(t, *tc.msg.InitialGenesis, chain.InitialGenesis)
			} else {
				require.EqualValues(t, previousChain.InitialGenesis, chain.InitialGenesis)
			}
		})
	}
}
