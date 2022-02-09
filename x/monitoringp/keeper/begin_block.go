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
	// TODO(489): define a proper value
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
			// TODO(490): implement correct address format
			monitoringInfo.SignatureCounts.AddSignature(vote.Validator.Address, valSetSize)
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
	if blockHeight < k.LastBlockHeight(ctx) {
		return nil
	}
	cid, cidFound := k.GetConnectionChannelID(ctx)
	if !cidFound {
		return nil
	}
	mi, miFound := k.GetMonitoringInfo(ctx)
	if !miFound || mi.Transmitted {
		return nil
	}

	// transmit signature packet
	err := k.TransmitMonitoringPacket(
		ctx,
		spntypes.MonitoringPacket{
			BlockHeight:     blockHeight,
			SignatureCounts: mi.SignatureCounts,
		},
		types.PortID,
		cid.ChannelID,
		clienttypes.ZeroHeight(),
		uint64(ctx.BlockTime().Add(MonitoringPacketTimeoutDelay).UnixNano()),
	)
	if err != nil {
		k.SetConsumerClientID(ctx, types.ConsumerClientID{
			ClientID: err.Error(),
		})
		return err
	}

	// signatures have been transmitted
	mi.Transmitted = true
	k.SetMonitoringInfo(ctx, mi)
	return nil
}
