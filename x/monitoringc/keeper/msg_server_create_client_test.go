package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringc/types"
	"testing"
)

func Test_msgServer_CreateClient(t *testing.T) {
	var (
		coordinator1                = sample.Coordinator(sample.Address())
		coordinator2                = sample.Coordinator(sample.Address())
		invalidChain                = uint64(1000)
		monitoringKeeper, profileKeeper, launchKeeper, msgSrv, msgSrvProfile, msgSrvLaunch, ibcKeeper, sdkCtx = setupMsgServer(t)
		ctx = sdk.WrapSDKContext(sdkCtx)
	)

	tests := []struct {
		name    string
		msg  types.MsgCreateClient
		err error
	}{
		{

		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := msgSrv.CreateClient(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			// verify the client is created
			verifiedClients, found := monitoringKeeper.GetVerifiedClientID(sdkCtx, tt.msg.LaunchID)
			require.True(t, found, "verified client ID should be added in the list")
			require.EqualValues(t, tt.msg.LaunchID, verifiedClients.LaunchID)
			require.Contains(t, res.ClientID, verifiedClients.ClientIDs)

			launchIDFromClient, found := monitoringKeeper.GetLaunchIDFromVerifiedClientID(sdkCtx, res.ClientID)
			require.True(t, found, "launch ID should be registered for the verified client ID")
			require.EqualValues(t, res.ClientID, launchIDFromClient.ClientID)
			require.EqualValues(t, tt.msg.LaunchID, launchIDFromClient.LaunchID)

			// IBC client should be created
			_, found = ibcKeeper.ClientKeeper.GetClientState(sdkCtx, res.ClientID)
			require.True(t, found, "IBC consumer client state should be created")
		})
	}
}
