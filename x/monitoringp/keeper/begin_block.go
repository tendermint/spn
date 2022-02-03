package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/monitoringp/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	// MonitoringPacketTimeoutDelay is the delay before a monitoring packet is timed out
	// The current value represents a year
	// TODO: define a proper value
	MonitoringPacketTimeoutDelay = time.Hour * 8760
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

// TransmitSignatures transmits over IBC the signatures to consumer if height is reached
// and signatures are not yet transmitted
func (k Keeper) TransmitSignatures(ctx sdk.Context, blockHeight int64) error {
	// check condition to transmit packet
	// IBC connection to consumer must be established
	// last block height must be reached
	// monitoring info must exist
	// signatures must not yet be transmitted
	cid, cidFound := k.GetConnectionChannelID(ctx)
	mi, miFound := k.GetMonitoringInfo(ctx)
	if !cidFound || blockHeight < k.LastBlockHeight(ctx) || !miFound || mi.Transmitted {
		return nil
	}

	// transmit signature packet
	err := k.TransmitMonitoringPacket(
		ctx,
		spntypes.MonitoringPacket{},
		types.PortID,
		cid.ChannelID,
		clienttypes.ZeroHeight(),
		uint64(ctx.BlockTime().Add(MonitoringPacketTimeoutDelay).Unix()),
	)
	if err != nil {
		return err
	}

	// signatures have been transmitted
	mi.Transmitted = true
	k.SetMonitoringInfo(ctx, mi)
	return nil
}
