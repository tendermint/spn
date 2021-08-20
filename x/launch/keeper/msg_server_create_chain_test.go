package keeper_test

import (
	"testing"

	codec "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/types"
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
		wantedChainID string
		valid         bool
	}{
		{
			name:          "valid message",
			msg:           sample.MsgCreateChain(coordAddress, "foo", ""),
			wantedChainID: "foo-0",
			valid:         true,
		},
		{
			name:          "an existing chain name creates a unique chain ID",
			msg:           sample.MsgCreateChain(coordAddress, "foo", ""),
			wantedChainID: "foo-1",
			valid:         true,
		},
		{
			name:          "valid message with genesis url",
			msg:           sample.MsgCreateChain(coordAddress, "bar", "foo.com"),
			wantedChainID: "bar-0",
			valid:         true,
		},
		{
			name:          "coordinator doesn't exist for the chain",
			msg:           sample.MsgCreateChain(sample.AccAddress(), "foo", ""),
			wantedChainID: "",
			valid:         false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := srv.CreateChain(ctx, &tc.msg)
			if !tc.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.EqualValues(t, tc.wantedChainID, got.ChainID)

			// The chain must exist in the store
			chain, found := k.GetChain(sdkCtx, got.ChainID)
			require.True(t, found)
			require.EqualValues(t, coordID, chain.CoordinatorID)
			require.EqualValues(t, got.ChainID, chain.ChainID)
			require.EqualValues(t, tc.msg.SourceURL, chain.SourceURL)
			require.EqualValues(t, tc.msg.SourceHash, chain.SourceHash)

			// Compare initial genesis
			if tc.msg.GenesisURL == "" {
				// Empty structure are nullified for Any type when encoded
				expectedDefault, _ := codec.NewAnyWithValue(&types.DefaultInitialGenesis{})
				expectedDefault.Value = nil
				expectedDefault.ClearCachedValue()
				require.Equal(t, expectedDefault, chain.InitialGenesis)
			} else {
				expectedGenesisURL, _ := codec.NewAnyWithValue(&types.GenesisURL{
					Url:  tc.msg.GenesisURL,
					Hash: tc.msg.GenesisHash,
				})
				expectedGenesisURL.ClearCachedValue()
				require.Equal(
					t,
					expectedGenesisURL,
					chain.InitialGenesis,
				)
			}
		})
	}
}
