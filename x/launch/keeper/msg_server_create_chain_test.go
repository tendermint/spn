package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func TestMsgCreateChain(t *testing.T) {
	k, _, srv, profileSrv, sdkCtx, _ := setupMsgServer(t)
	ctx := sdk.WrapSDKContext(sdkCtx)
	coordAddress := sample.AccAddress()

	// Create a coordinator
	msgCreateCoordinator := sample.MsgCreateCoordinator(coordAddress)
	res, err := profileSrv.CreateCoordinator(ctx, &msgCreateCoordinator)
	require.NoError(t, err)
	coordID := res.CoordinatorId

	for _, tc := range []struct {
		name          string
		msg           types.MsgCreateChain
		wantedChainID uint64
		err           error
	}{
		{
			name:          "valid message",
			msg:           sample.MsgCreateChain(coordAddress, ""),
			wantedChainID: 0,
		},
		{
			name:          "creates a unique chain ID",
			msg:           sample.MsgCreateChain(coordAddress, ""),
			wantedChainID: 1,
		},
		{
			name:          "valid message with genesis url",
			msg:           sample.MsgCreateChain(coordAddress, "foo.com"),
			wantedChainID: 2,
		},
		{
			name: "coordinator doesn't exist for the chain",
			msg:  sample.MsgCreateChain(sample.AccAddress(), ""),
			err:  profiletypes.ErrCoordAddressNotFound,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := srv.CreateChain(ctx, &tc.msg)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
				return
			}
			require.NoError(t, err)
			require.EqualValues(t, tc.wantedChainID, got.Id)

			// The chain must exist in the store
			chain, found := k.GetChain(sdkCtx, got.Id)
			require.True(t, found)
			require.EqualValues(t, coordID, chain.CoordinatorID)
			require.EqualValues(t, got.Id, chain.Id)
			require.EqualValues(t, tc.msg.GenesisChainID, chain.GenesisChainID)
			require.EqualValues(t, tc.msg.SourceURL, chain.SourceURL)
			require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)

			// Compare initial genesis
			if tc.msg.GenesisURL == "" {
				require.Equal(t, types.NewDefaultInitialGenesis(), chain.InitialGenesis)
			} else {
				require.Equal(
					t,
					types.NewGenesisURL(tc.msg.GenesisURL, tc.msg.GenesisHash),
					chain.InitialGenesis,
				)
			}
		})
	}
}
