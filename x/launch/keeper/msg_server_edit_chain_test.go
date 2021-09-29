package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgEditChain(t *testing.T) {
	k, _, _, srv, _, profileSrv, sdkCtx := setupMsgServer(t)
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

	// Create a chain
	msgCreateChain := sample.MsgCreateChain(coordAddress, "", 0)
	res, err := srv.CreateChain(ctx, &msgCreateChain)
	require.NoError(t, err)
	chainID := res.Id

	for _, tc := range []struct {
		name string
		msg  types.MsgEditChain
		err  error
	}{
		{
			name: "edit genesis chain ID",
			msg: sample.MsgEditChain(coordAddress, chainID,
				true,
				false,
				false,
				false,
			),
		},
		{
			name: "edit source",
			msg: sample.MsgEditChain(coordAddress, chainID,
				false,
				true,
				false,
				false,
			),
		},
		{
			name: "edit initial genesis with default genesis",
			msg: sample.MsgEditChain(coordAddress, chainID,
				false,
				false,
				true,
				false,
			),
		},
		{
			name: "edit initial genesis with genesis url",
			msg: sample.MsgEditChain(coordAddress, chainID,
				false,
				false,
				true,
				true,
			),
		},
		{
			name: "edit source and initial genesis",
			msg: sample.MsgEditChain(coordAddress, chainID,
				false,
				true,
				true,
				true,
			),
		},
		{
			name: "non existent chain id",
			msg: sample.MsgEditChain(coordAddress, chainIDNoExist,
				false,
				true,
				false,
				false,
			),
			err: types.ErrChainNotFound,
		},
		{
			name: "non existent coordinator",
			msg: sample.MsgEditChain(coordNoExist, chainID,
				false,
				true,
				false,
				false,
			),
			err: profiletypes.ErrCoordAddressNotFound,
		},
		{
			name: "invalid coordinator",
			msg: sample.MsgEditChain(coordAddress2, chainID,
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
				previousChain, found = k.GetChain(sdkCtx, tc.msg.ChainID)
				require.True(t, found)
			}

			// Send the message
			_, err := srv.EditChain(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)

			// The chain must continue to exist in the store
			chain, found := k.GetChain(sdkCtx, tc.msg.ChainID)
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
