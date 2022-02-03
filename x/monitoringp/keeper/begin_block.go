package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/monitoringp/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ReportBlockSignatures gets signatures from blocks and update monitoring info
func (k Keeper) ReportBlockSignatures(ctx sdk.Context, lastCommit abci.LastCommitInfo, blockHeight int64) {
	// no report if last height is reached
	lastBlockHeight := k.LastBlockHeight(ctx)
	if blockHeight > lastBlockHeight {
		return
	}

	// get monitoring info
	monitoringInfo, found := k.GetMonitoringInfo(ctx)
	if !found {
		monitoringInfo = types.MonitoringInfo{
			SignatureCounts: spntypes.NewSignatureCounts(),
		}
	}

	// update signatures with voters that signed blocks
	valSetSize := int64(len(lastCommit.Votes))
	for _, vote := range lastCommit.Votes {
		if vote.SignedLastBlock {
			// TODO: implement correct address format
			monitoringInfo.SignatureCounts.AddSignature(string(vote.Validator.Address), valSetSize)
		}
	}

	// increment block count and save the monitoring info
	monitoringInfo.SignatureCounts.BlockCount++
	k.SetMonitoringInfo(ctx, monitoringInfo)
}