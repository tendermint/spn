package keeper

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	abci "github.com/tendermint/tendermint/abci/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/monitoringp/types"
)

const (
	// MonitoringPacketTimeoutDelay is the delay before a monitoring packet is timed out
	// The timeout is set to one year
	// This is an arbitrarily chosen value that should never be reached in practice
	MonitoringPacketTimeoutDelay = time.Hour * 8760
)

// ReportBlockSignatures gets signatures from blocks and update monitoring info
func (k Keeper) ReportBlockSignatures(ctx sdk.Context, lastCommit abci.LastCommitInfo, blockHeight int64) error {
	// skip first block because it is not signed
	if blockHeight == 1 {
		return nil
	}

	// no report if last height is reached
	lastBlockHeight := k.LastBlockHeight(ctx)
	if blockHeight > lastBlockHeight {
		return nil
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
			// get the operator address from the consensus address
			val, found := k.stakingKeeper.GetValidatorByConsAddr(ctx, vote.Validator.Address)
			if !found {
				return fmt.Errorf("validator from consensus address %s not found", vote.Validator.Address)
			}

			monitoringInfo.SignatureCounts.AddSignature(val.OperatorAddress, valSetSize)
		}
	}

	// increment block count and save the monitoring info
	monitoringInfo.SignatureCounts.BlockCount++
	k.SetMonitoringInfo(ctx, monitoringInfo)

	return nil
}

// TransmitSignatures transmits over IBC the signatures to consumer if height is reached
// and signatures are not yet transmitted
func (k Keeper) TransmitSignatures(ctx sdk.Context, blockHeight int64) (sequence uint64, err error) {
	// check condition to transmit packet
	// IBC connection to consumer must be established
	// last block height must be reached
	// monitoring info must exist
	// signatures must not yet be transmitted
	if blockHeight < k.LastBlockHeight(ctx) {
		return 0, nil
	}
	cid, cidFound := k.GetConnectionChannelID(ctx)
	if !cidFound {
		return 0, nil
	}
	mi, miFound := k.GetMonitoringInfo(ctx)
	if !miFound || mi.Transmitted {
		return 0, nil
	}

	// transmit signature packet
	sequence, err = k.TransmitMonitoringPacket(
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
		return 0, err
	}

	// signatures have been transmitted
	mi.Transmitted = true
	k.SetMonitoringInfo(ctx, mi)
	return sequence, nil
}
