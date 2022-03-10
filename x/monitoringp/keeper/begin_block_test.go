package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringp/types"
)

func TestKeeper_ReportBlockSignatures(t *testing.T) {
	ctx, tk, _ := testkeeper.NewTestSetupWithMonitoringp(t)
	valFoo, valBar, valBaz, valFred, valQux := sample.Validator(t),
		sample.Validator(t),
		sample.Validator(t),
		sample.Validator(t),
		sample.Validator(t)
	consFoo, err := valFoo.GetConsAddr()
	require.NoError(t, err)
	consBar, err := valBar.GetConsAddr()
	require.NoError(t, err)
	consBaz, err := valBaz.GetConsAddr()
	require.NoError(t, err)
	consFred, err := valFred.GetConsAddr()
	require.NoError(t, err)
	consQux, err := valQux.GetConsAddr()
	require.NoError(t, err)

	// consensus address with no validator associated
	consNoValidator := sample.ConsAddress()

	// initialize staking validator set
	tk.StakingKeeper.SetValidator(ctx, valFoo)
	tk.StakingKeeper.SetValidator(ctx, valBar)
	tk.StakingKeeper.SetValidator(ctx, valBaz)
	tk.StakingKeeper.SetValidator(ctx, valFred)
	tk.StakingKeeper.SetValidator(ctx, valQux)
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
			name:                        "lastBlockHeight reached doesn't create a non existent monitoring info",
			monitoringInfoExist:         false,
			lastBlockHeight:             10,
			currentBlockHeight:          11,
			expectedMonitoringInfoFound: false,
		},
		{
			name:                        "no monitoring info created because counting skipped if blockHeight == 1",
			monitoringInfoExist:         false,
			lastBlockHeight:             1,
			currentBlockHeight:          1,
			expectedMonitoringInfoFound: false,
		},
		{
			name:                "lastBlockHeight reached doesn't update an existent monitoring info",
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
			name:                "if monitoring info doesn't exists, the structure is created with block count to 1 and signatures from commit",
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
			name:                "monitoring info should be updated following signatures in the last commit",
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
