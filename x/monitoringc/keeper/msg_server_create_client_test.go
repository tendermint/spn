package keeper_test

import (
	"encoding/base64"
	"testing"
	"time"

	ibctmtypes "github.com/cosmos/ibc-go/v2/modules/light-clients/07-tendermint/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/testutil/sample"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/spn/x/monitoringc/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func Test_msgServer_CreateClient(t *testing.T) {
	var (
		coordAddr    = sample.Address()
		invalidChain = uint64(1000)

		monitoringKeeper, _, _, msgSrv, msgSrvProfile, msgSrvLaunch, ibcKeeper, sdkCtx = setupMsgServer(t)

		ctx = sdk.WrapSDKContext(sdkCtx)

		consPubKeyStr = "jP0v8F0e2kSAS367V/QAikddQPze+V36v7lhkv1Iqgg="
		cs            = spntypes.NewConsensusState(
			"2022-02-08T15:12:36.161481Z",
			"A13E761948413E405EA4F09BEC9F37632F739404108FE1635CB3529B61DA9FD7",
			"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		)
		vs = spntypes.NewValidatorSet(
			spntypes.NewValidator(consPubKeyStr, 0, 100),
		)
	)

	selfDelegation, err := sdk.ParseCoinNormalized("1000stake")
	require.NoError(t, err)
	consPubKey, err := base64.StdEncoding.DecodeString(consPubKeyStr)
	require.NoError(t, err)

	// create a coordinator and a chain with a genesis validator
	_, err = msgSrvProfile.CreateCoordinator(ctx, profiletypes.NewMsgCreateCoordinator(
		coordAddr,
		"",
		"",
		"",
	))
	require.NoError(t, err)
	resCreateChain, err := msgSrvLaunch.CreateChain(ctx, launchtypes.NewMsgCreateChain(
		coordAddr,
		"orbit-1",
		sample.String(10),
		sample.String(10),
		"",
		"",
		false,
		0,
		sample.Metadata(20),
	))
	require.NoError(t, err)
	_, err = msgSrvLaunch.RequestAddValidator(ctx, launchtypes.NewMsgRequestAddValidator(
		coordAddr,
		resCreateChain.LaunchID,
		sample.Address(),
		sample.Bytes(100),
		consPubKey,
		selfDelegation,
		sample.GenesisValidatorPeer(),
	))
	require.NoError(t, err)
	_, err = msgSrvLaunch.TriggerLaunch(ctx, launchtypes.NewMsgTriggerLaunch(
		coordAddr,
		resCreateChain.LaunchID,
		launchtypes.DefaultMinLaunchTime,
	))
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgCreateClient
		err  error
	}{
		{
			name: "chain doesn't exist",
			msg: *types.NewMsgCreateClient(
				sample.Address(),
				invalidChain,
				cs,
				vs,
				spntypes.DefaultUnbondingPeriod,
				spntypes.DefaultRevisionHeight,
			),
			err: launchtypes.ErrChainNotFound,
		},
		{
			name: "invalid validator set",
			msg: *types.NewMsgCreateClient(
				sample.Address(),
				resCreateChain.LaunchID,
				sample.ConsensusState(0),
				sample.ValidatorSet(1),
				spntypes.DefaultUnbondingPeriod,
				spntypes.DefaultRevisionHeight,
			),
			err: types.ErrInvalidValidatorSet,
		},
		{
			name: "verified client should be created",
			msg: *types.NewMsgCreateClient(
				sample.Address(),
				resCreateChain.LaunchID,
				cs,
				vs,
				spntypes.DefaultUnbondingPeriod,
				spntypes.DefaultRevisionHeight,
			),
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
			require.Contains(t, verifiedClients.ClientIDs, res.ClientID)

			launchIDFromClient, found := monitoringKeeper.GetLaunchIDFromVerifiedClientID(sdkCtx, res.ClientID)
			require.True(t, found, "launch ID should be registered for the verified client ID")
			require.EqualValues(t, res.ClientID, launchIDFromClient.ClientID)
			require.EqualValues(t, tt.msg.LaunchID, launchIDFromClient.LaunchID)

			// IBC client should be created
			clientState, found := ibcKeeper.ClientKeeper.GetClientState(sdkCtx, res.ClientID)
			require.True(t, found, "IBC consumer client state should be created")
			cs, ok := clientState.(*ibctmtypes.ClientState)
			require.True(t, ok)
			require.EqualValues(t, tt.msg.RevisionHeight, cs.LatestHeight.RevisionHeight)
			require.EqualValues(t, time.Second*time.Duration(tt.msg.UnbondingPeriod), cs.UnbondingPeriod)
		})
	}
}
