package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

// ChannelMock represents a mocked IBC channel keeper used for testnet purpose
type ChannelMock struct {

}

// GetChannel implements ChannelKeeper
func (c ChannelMock) GetChannel(ctx sdk.Context, srcPort, srcChan string) (channel channeltypes.Channel, found bool) {
	return channeltypes.Channel{}, false
}

// GetNextSequenceSend implements ChannelKeeper
func (c ChannelMock) GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool){
	return 0, false
}

// SendPacket implements ChannelKeeper
func (c ChannelMock) SendPacket(ctx sdk.Context, channelCap *capabilitytypes.Capability, packet exported.PacketI) error{
	return nil
}

// WriteAcknowledgement implements ChannelKeeper
func (c ChannelMock) WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI, acknowledgement []byte) error{
	return nil
}

// ChanCloseInit implements ChannelKeeper
func (c ChannelMock) ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error{
	return nil
}