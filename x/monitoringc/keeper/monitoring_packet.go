package keeper

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	"github.com/tendermint/spn/x/monitoringc/types"
)

// OnRecvMonitoringPacket processes packet reception
func (k Keeper) OnRecvMonitoringPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data spntypes.MonitoringPacket,
) (packetAck spntypes.MonitoringPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// retrieve launch ID for channel ID
	lidFromCid, found := k.GetLaunchIDFromChannelID(ctx, packet.DestinationChannel)
	if !found {
		return packetAck, fmt.Errorf("no launch ID associated to channel ID %s", packet.DestinationChannel)
	}

	// save the latest received monitoring packet for documentation purpose
	k.SetMonitoringHistory(ctx, types.MonitoringHistory{
		LaunchID:               lidFromCid.LaunchID,
		LatestMonitoringPacket: data,
	})

	// distribute reward from the signature count
	err = k.rewardKeeper.DistributeRewards(
		ctx,
		lidFromCid.LaunchID,
		data.SignatureCounts,
		data.BlockHeight,
		true,
	)

	return packetAck, err
}

// OnAcknowledgementMonitoringPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementMonitoringPacket(
	_ sdk.Context,
	_ channeltypes.Packet,
	_ spntypes.MonitoringPacket,
	_ channeltypes.Acknowledgement,
) error {
	return errors.New("not implemented")
}

// OnTimeoutMonitoringPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutMonitoringPacket(
	_ sdk.Context,
	_ channeltypes.Packet,
	_ spntypes.MonitoringPacket,
) error {
	return errors.New("not implemented")
}
