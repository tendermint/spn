package keeper_test

import (
	"encoding/base64"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ibctmtypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	ignterrors "github.com/ignite/modules/pkg/errors"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	launchtypes "github.com/tendermint/spn/x/launch/types"
	"github.com/tendermint/spn/x/monitoringc/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
)

func Test_msgServer_CreateClient(t *testing.T) {
	var (
		coordAddr    = sample.Address(r)
		invalidChain = uint64(1000)

		sdkCtx, tk, ts = testkeeper.NewTestSetup(t)
		ctx            = sdk.WrapSDKContext(sdkCtx)

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
	resCoord, err := ts.ProfileSrv.CreateCoordinator(ctx, profiletypes.NewMsgCreateCoordinator(
		coordAddr,
		"",
		"",
		"",
	))
	require.NoError(t, err)
	initialGenesis := launchtypes.NewDefaultInitialGenesis()
	resCreateChain, err := ts.LaunchSrv.CreateChain(ctx, launchtypes.NewMsgCreateChain(
		coordAddr,
		"orbit-1",
		sample.String(r, 10),
		sample.String(r, 10),
		initialGenesis,
		false,
		0,
		sample.Coins(r),
		sample.Metadata(r, 20),
	))
	require.NoError(t, err)
	chainWithInvalidChainID := sample.Chain(r, resCreateChain.LaunchID+1, resCoord.CoordinatorID)
	chainWithInvalidChainID.GenesisChainID = "invalid_chain_id"
	tk.LaunchKeeper.SetChain(sdkCtx, chainWithInvalidChainID)
	_, err = ts.LaunchSrv.SendRequest(ctx, launchtypes.NewMsgSendRequest(
		coordAddr,
		resCreateChain.LaunchID,
		launchtypes.NewGenesisValidator(
			resCreateChain.LaunchID,
			sample.Address(r),
			sample.Bytes(r, 100),
			consPubKey,
			selfDelegation,
			sample.GenesisValidatorPeer(r),
		),
	))
	require.NoError(t, err)
	_, err = ts.LaunchSrv.TriggerLaunch(ctx, launchtypes.NewMsgTriggerLaunch(
		coordAddr,
		resCreateChain.LaunchID,
		sdkCtx.BlockTime().Add(launchtypes.DefaultMinLaunchTime),
	))
	require.NoError(t, err)

	tests := []struct {
		name string
		msg  types.MsgCreateClient
		err  error
	}{
		{
			name: "invalid chain ID",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				chainWithInvalidChainID.LaunchID,
				cs,
				vs,
				spntypes.DefaultUnbondingPeriod,
				spntypes.DefaultRevisionHeight,
			),
			err: types.ErrInvalidClientState,
		},
		{
			name: "invalid client state",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				resCreateChain.LaunchID,
				cs,
				vs,
				0,
				spntypes.DefaultRevisionHeight,
			),
			err: types.ErrInvalidClientState,
		},
		{
			name: "invalid consensus state",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				resCreateChain.LaunchID,
				spntypes.NewConsensusState(
					"",
					"",
					"",
				),
				vs,
				spntypes.DefaultUnbondingPeriod,
				spntypes.DefaultRevisionHeight,
			),
			err: ignterrors.ErrCritical,
		},
		{
			name: "chain doesn't exist",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				invalidChain,
				cs,
				vs,
				spntypes.DefaultUnbondingPeriod,
				spntypes.DefaultRevisionHeight,
			),
			err: launchtypes.ErrChainNotFound,
		},
		{
			name: "empty validator set",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
				resCreateChain.LaunchID,
				sample.ConsensusState(0),
				spntypes.ValidatorSet{},
				spntypes.DefaultUnbondingPeriod,
				spntypes.DefaultRevisionHeight,
			),
			err: ignterrors.ErrCritical,
		},
		{
			name: "invalid validator set",
			msg: *types.NewMsgCreateClient(
				sample.Address(r),
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
				sample.Address(r),
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
			res, err := ts.MonitoringcSrv.CreateClient(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorIs(t, tt.err, err)
				return
			}
			require.NoError(t, err)

			// verify the client is created
			verifiedClients, found := tk.MonitoringConsumerKeeper.GetVerifiedClientID(sdkCtx, tt.msg.LaunchID)
			require.True(t, found, "verified client ID should be added in the list")
			require.EqualValues(t, tt.msg.LaunchID, verifiedClients.LaunchID)
			require.Contains(t, verifiedClients.ClientIDs, res.ClientID)

			launchIDFromClient, found := tk.MonitoringConsumerKeeper.GetLaunchIDFromVerifiedClientID(sdkCtx, res.ClientID)
			require.True(t, found, "launch ID should be registered for the verified client ID")
			require.EqualValues(t, res.ClientID, launchIDFromClient.ClientID)
			require.EqualValues(t, tt.msg.LaunchID, launchIDFromClient.LaunchID)

			// IBC client should be created
			clientState, found := tk.IBCKeeper.ClientKeeper.GetClientState(sdkCtx, res.ClientID)
			require.True(t, found, "IBC consumer client state should be created")
			cs, ok := clientState.(*ibctmtypes.ClientState)
			require.True(t, ok)
			require.EqualValues(t, tt.msg.RevisionHeight, cs.LatestHeight.RevisionHeight)
			require.EqualValues(t, time.Second*time.Duration(tt.msg.UnbondingPeriod), cs.UnbondingPeriod)
		})
	}
}
