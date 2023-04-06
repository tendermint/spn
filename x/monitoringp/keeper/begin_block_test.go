package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/stretchr/testify/require"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestKeeper_ReportBlockSignatures(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)

	var (
		valFoo  = sample.Validator(t, r)
		valBar  = sample.Validator(t, r)
		valBaz  = sample.Validator(t, r)
		valFred = sample.Validator(t, r)
		valQux  = sample.Validator(t, r)

		consFoo  sdk.ConsAddress
		consBar  sdk.ConsAddress
		consBaz  sdk.ConsAddress
		consFred sdk.ConsAddress
		consQux  sdk.ConsAddress
		err      error
	)

	t.Run("should get consensus addresses", func(t *testing.T) {
		consFoo, err = valFoo.GetConsAddr()
		require.NoError(t, err)
		consBar, err = valBar.GetConsAddr()
		require.NoError(t, err)
		consBaz, err = valBaz.GetConsAddr()
		require.NoError(t, err)
		consFred, err = valFred.GetConsAddr()
		require.NoError(t, err)
		consQux, err = valQux.GetConsAddr()
		require.NoError(t, err)
	})

	// consensus address with no validator associated
	consNoValidator := sample.ConsAddress(r)

	// initialize staking validator set
	tk.StakingKeeper.SetValidator(ctx, valFoo)
	tk.StakingKeeper.SetValidator(ctx, valBar)
	tk.StakingKeeper.SetValidator(ctx, valBaz)
	tk.StakingKeeper.SetValidator(ctx, valFred)
	tk.StakingKeeper.SetValidator(ctx, valQux)

	t.Run("should set validators by consensus address", func(t *testing.T) {
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valFoo)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valBar)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valBaz)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valFred)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valQux)
		require.NoError(t, err)
	})

	tests := []struct {
		name                        string
		monitoringInfoExist         bool
		inputMonitoringInfo         types.MonitoringInfo
		lastBlockHeight             int64
		lastCommitInfo              abci.LastCommitInfo
		currentBlockHeight          int64
		expectedMonitoringInfoFound bool
		expectedMonitoringInfo      types.MonitoringInfo
		wantErr                     bool
	}{
		{
			name:                        "should not create monitoring info with lastBlockHeight reached",
			monitoringInfoExist:         false,
			lastBlockHeight:             10,
			currentBlockHeight:          11,
			expectedMonitoringInfoFound: false,
		},
		{
			name: "should not create monitoring info created " +
				"because counting skipped if blockHeight == 1",
			monitoringInfoExist:         false,
			lastBlockHeight:             1,
			currentBlockHeight:          1,
			expectedMonitoringInfoFound: false,
		},
		{
			name:                "should not update with lastBlockHeight reached",
			monitoringInfoExist: true,
			inputMonitoringInfo: tc.MonitoringInfo(10,
				tc.SignatureCount(t,
					valFoo.OperatorAddress,
					"1",
				),
				tc.SignatureCount(t,
					valBar.OperatorAddress,
					"2",
				),
			),
			lastBlockHeight: 10,
			lastCommitInfo: tc.LastCommitInfo(
				tc.Vote{
					Address: consFoo,
					Signed:  true,
				},
			),
			currentBlockHeight:          11,
			expectedMonitoringInfoFound: true,
			expectedMonitoringInfo: tc.MonitoringInfo(10,
				tc.SignatureCount(t,
					valFoo.OperatorAddress,
					"1",
				),
				tc.SignatureCount(t,
					valBar.OperatorAddress,
					"2",
				),
			),
		},
		{
			name: "should create structure if monitoring info doesn't exist with " +
				"block count to 1 and signatures from commit",
			monitoringInfoExist: false,
			lastBlockHeight:     10,
			lastCommitInfo: tc.LastCommitInfo(
				tc.Vote{
					Address: consFoo,
					Signed:  true,
				},
			),
			currentBlockHeight:          2,
			expectedMonitoringInfoFound: true,
			expectedMonitoringInfo: tc.MonitoringInfo(1,
				tc.SignatureCount(t,
					valFoo.OperatorAddress,
					"1",
				),
			),
		},
		{
			name:                "should update monitoring info following signatures in the last commit",
			monitoringInfoExist: true,
			inputMonitoringInfo: tc.MonitoringInfo(50,
				tc.SignatureCount(t,
					valFoo.OperatorAddress,
					"1",
				),
				tc.SignatureCount(t,
					valBar.OperatorAddress,
					"2",
				),
				tc.SignatureCount(t,
					valBaz.OperatorAddress,
					"3",
				),
			),
			lastBlockHeight: 10,
			lastCommitInfo: tc.LastCommitInfo(
				tc.Vote{
					Address: consFoo,
					Signed:  true,
				},
				tc.Vote{
					Address: consBar,
					Signed:  false,
				},
				tc.Vote{
					Address: consBaz,
					Signed:  true,
				},
				tc.Vote{
					Address: consQux,
					Signed:  false,
				},
				tc.Vote{
					Address: consFred,
					Signed:  true,
				},
			),
			currentBlockHeight:          2,
			expectedMonitoringInfoFound: true,
			expectedMonitoringInfo: tc.MonitoringInfo(51,
				tc.SignatureCount(t,
					valFoo.OperatorAddress,
					"1.2",
				),
				tc.SignatureCount(t,
					valBar.OperatorAddress,
					"2",
				),
				tc.SignatureCount(t,
					valBaz.OperatorAddress,
					"3.2",
				),
				tc.SignatureCount(t,
					valFred.OperatorAddress,
					"0.2",
				),
			),
		},
		{
			name:                "should prevent reporting signatures when a signer doesn't have an associated validator",
			monitoringInfoExist: true,
			inputMonitoringInfo: tc.MonitoringInfo(50,
				tc.SignatureCount(t,
					valFoo.OperatorAddress,
					"1",
				),
			),
			lastBlockHeight: 10,
			lastCommitInfo: tc.LastCommitInfo(
				tc.Vote{
					Address: consNoValidator,
					Signed:  true,
				},
			),
			currentBlockHeight: 2,
			wantErr:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set keeper values
			params := tk.MonitoringProviderKeeper.GetParams(ctx)
			params.LastBlockHeight = tt.lastBlockHeight
			tk.MonitoringProviderKeeper.SetParams(ctx, params)
			if tt.monitoringInfoExist {
				tk.MonitoringProviderKeeper.SetMonitoringInfo(ctx, tt.inputMonitoringInfo)
			} else {
				tk.MonitoringProviderKeeper.RemoveMonitoringInfo(ctx)
			}

			// report
			err := tk.MonitoringProviderKeeper.ReportBlockSignatures(ctx, tt.lastCommitInfo, tt.currentBlockHeight)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// check saved values
			monitoringInfo, found := tk.MonitoringProviderKeeper.GetMonitoringInfo(ctx)
			require.EqualValues(t, tt.expectedMonitoringInfoFound, found)
			require.EqualValues(t, tt.expectedMonitoringInfo, monitoringInfo)
		})
	}
}

func TestKeeper_TransmitSignatures(t *testing.T) {
	ctx, tk, _ := monitoringpKeeperWithFooClient(t)
	valFoo, valBar, valBaz, valFred, valQux := sample.Validator(t, r),
		sample.Validator(t, r),
		sample.Validator(t, r),
		sample.Validator(t, r),
		sample.Validator(t, r)

	// initialize staking validator set
	tk.StakingKeeper.SetValidator(ctx, valFoo)
	tk.StakingKeeper.SetValidator(ctx, valBar)
	tk.StakingKeeper.SetValidator(ctx, valBaz)
	tk.StakingKeeper.SetValidator(ctx, valFred)
	tk.StakingKeeper.SetValidator(ctx, valQux)

	t.Run("should set validators by consensus address", func(t *testing.T) {
		err := tk.StakingKeeper.SetValidatorByConsAddr(ctx, valFoo)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valBar)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valBaz)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valFred)
		require.NoError(t, err)
		err = tk.StakingKeeper.SetValidatorByConsAddr(ctx, valQux)
		require.NoError(t, err)
	})

	tests := []struct {
		name                        string
		monitoringInfoExist         bool
		inputMonitoringInfo         types.MonitoringInfo
		lastBlockHeight             int64
		currentBlockHeight          int64
		channelIDExist              bool
		channelID                   types.ConnectionChannelID
		expectedMonitoringInfoFound bool
		expectedMonitoringInfo      types.MonitoringInfo
		wantErr                     bool
	}{
		{
			name:                "should return monitoring info with channel not found",
			monitoringInfoExist: true,
			inputMonitoringInfo: tc.MonitoringInfo(1,
				tc.SignatureCount(t,
					valFoo.OperatorAddress,
					"1")),
			lastBlockHeight:    10,
			currentBlockHeight: 11,
			channelIDExist:     true,
			channelID:          types.ConnectionChannelID{ChannelID: "channelID"},
			wantErr:            true,
		},
		{
			name:                        "should return nil for currentBlockHeight < lastBlockHeight",
			monitoringInfoExist:         false,
			lastBlockHeight:             11,
			currentBlockHeight:          10,
			channelIDExist:              false,
			expectedMonitoringInfoFound: false,
		},
		{
			name:                        "should return nil for lastBlockHeight no channel ID set",
			monitoringInfoExist:         false,
			lastBlockHeight:             10,
			currentBlockHeight:          11,
			channelIDExist:              false,
			expectedMonitoringInfoFound: false,
		},
		{
			name:                        "should return nil for no monitoring info found",
			monitoringInfoExist:         false,
			lastBlockHeight:             10,
			currentBlockHeight:          11,
			channelIDExist:              true,
			channelID:                   types.ConnectionChannelID{ChannelID: "channelID"},
			expectedMonitoringInfoFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set keeper values
			params := tk.MonitoringProviderKeeper.GetParams(ctx)
			params.LastBlockHeight = tt.lastBlockHeight
			tk.MonitoringProviderKeeper.SetParams(ctx, params)
			if tt.monitoringInfoExist {
				tk.MonitoringProviderKeeper.SetMonitoringInfo(ctx, tt.inputMonitoringInfo)
			} else {
				tk.MonitoringProviderKeeper.RemoveMonitoringInfo(ctx)
			}

			if tt.channelIDExist {
				tk.MonitoringProviderKeeper.SetConnectionChannelID(ctx, tt.channelID)
			}

			// report
			// TODO check sequence in test
			_, err := tk.MonitoringProviderKeeper.TransmitSignatures(ctx, tt.currentBlockHeight)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// check saved values
			monitoringInfo, found := tk.MonitoringProviderKeeper.GetMonitoringInfo(ctx)
			require.EqualValues(t, tt.expectedMonitoringInfoFound, found)
			require.EqualValues(t, tt.expectedMonitoringInfo, monitoringInfo)
		})
	}
}
