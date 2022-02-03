package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	spntypes "github.com/tendermint/spn/pkg/types"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/x/monitoringp/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestKeeper_ReportBlockSignatures(t *testing.T) {
	k, _, ctx := testkeeper.MonitoringpKeeper(t)

	// simplified type for abci.VoteInfo for testing purpose
	type vote struct {
		address string
		signed  bool
	}

	// create LastCommitInfo from list of addresses and signed boolean
	lastCommit := func(votes ...vote) abci.LastCommitInfo {
		var lci abci.LastCommitInfo

		// add votes
		for _, vote := range votes {
			lci.Votes = append(lci.Votes, abci.VoteInfo{
				Validator: abci.Validator{
					Address: []byte(vote.address),
				},
				SignedLastBlock: vote.signed,
			})
		}

		return lci
	}

	// simplified type for spntypes.SignatureCount for testing purpose
	type signatureCount struct {
		consAddr          string
		relativeSignature string
	}

	// create MonitoringInfo from list of addresses and associated relative signatures
	monitoringInfo := func(blockCount uint64, signatureCount ...signatureCount) types.MonitoringInfo {
		var mi types.MonitoringInfo
		mi.SignatureCounts.BlockCount = blockCount

		// add signature count
		for _, sc := range signatureCount {
			sigDec, err := sdk.NewDecFromStr(sc.relativeSignature)
			require.NoError(t, err)
			mi.SignatureCounts.Counts = append(mi.SignatureCounts.Counts, spntypes.SignatureCount{
				ConsAddress:        []byte(sc.consAddr),
				RelativeSignatures: sigDec,
			})
		}

		return mi
	}

	tests := []struct {
		name                        string
		monitoringInfoExist         bool
		inputMonitoringInfo         types.MonitoringInfo
		lastBlockHeight             int64
		lastCommitInfo              abci.LastCommitInfo
		currentBlockHeight          int64
		expectedMonitoringInfoFound bool
		expectedMonitoringInfo      types.MonitoringInfo
	}{
		{
			name:                        "lastBlockHeight reached doesn't create a non existent monitoring info",
			monitoringInfoExist:         false,
			lastBlockHeight:             10,
			currentBlockHeight:          11,
			expectedMonitoringInfoFound: false,
		},
		{
			name:                "lastBlockHeight reached doesn't update an existent monitoring info",
			monitoringInfoExist: true,
			inputMonitoringInfo: monitoringInfo(10,
				signatureCount{
					"foo",
					"1",
				},
				signatureCount{
					"bar",
					"2",
				},
			),
			lastBlockHeight: 10,
			lastCommitInfo: lastCommit(
				vote{
					address: "foo",
					signed:  true,
				},
			),
			currentBlockHeight:          11,
			expectedMonitoringInfoFound: true,
			expectedMonitoringInfo: monitoringInfo(10,
				signatureCount{
					"foo",
					"1",
				},
				signatureCount{
					"bar",
					"2",
				},
			),
		},
		{
			name:                "if monitoring info doesn't exists, the structure is created with block count to 1 and signatures from commit",
			monitoringInfoExist: false,
			lastBlockHeight:     10,
			lastCommitInfo: lastCommit(
				vote{
					address: "foo",
					signed:  true,
				},
			),
			currentBlockHeight:          1,
			expectedMonitoringInfoFound: true,
			expectedMonitoringInfo: monitoringInfo(1,
				signatureCount{
					"foo",
					"1",
				},
			),
		},
		{
			name:                "monitoring info should be updated following signatures in the last commit",
			monitoringInfoExist: true,
			inputMonitoringInfo: monitoringInfo(50,
				signatureCount{
					"foo",
					"1",
				},
				signatureCount{
					"bar",
					"2",
				},
				signatureCount{
					"baz",
					"3",
				},
			),
			lastBlockHeight: 10,
			lastCommitInfo: lastCommit(
				vote{
					address: "foo",
					signed:  true,
				},
				vote{
					address: "bar",
					signed:  false,
				},
				vote{
					address: "baz",
					signed:  true,
				},
				vote{
					address: "qux",
					signed:  false,
				},
				vote{
					address: "fred",
					signed:  true,
				},
			),
			currentBlockHeight:          1,
			expectedMonitoringInfoFound: true,
			expectedMonitoringInfo: monitoringInfo(51,
				signatureCount{
					"foo",
					"1.2",
				},
				signatureCount{
					"bar",
					"2",
				},
				signatureCount{
					"baz",
					"3.2",
				},
				signatureCount{
					"fred",
					"0.2",
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set keeper values
			params := k.GetParams(ctx)
			params.LastBlockHeight = tt.lastBlockHeight
			k.SetParams(ctx, params)
			if tt.monitoringInfoExist {
				k.SetMonitoringInfo(ctx, tt.inputMonitoringInfo)
			} else {
				k.RemoveMonitoringInfo(ctx)
			}

			// report
			k.ReportBlockSignatures(ctx, tt.lastCommitInfo, tt.currentBlockHeight)

			// check saved values
			monitoringInfo, found := k.GetMonitoringInfo(ctx)
			require.EqualValues(t, tt.expectedMonitoringInfoFound, found)
			require.EqualValues(t, tt.expectedMonitoringInfo, monitoringInfo)
		})
	}
}
