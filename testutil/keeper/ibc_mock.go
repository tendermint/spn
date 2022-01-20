package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
)

// ConnectionMock represents a mocked IBC connection keeper used for test purposes
type ConnectionMock struct {
	connections map[string]connectiontypes.ConnectionEnd
}

// SetConnection sets a connection for mocking purpose
func (c *ConnectionMock) SetConnection(connectionID string, connection connectiontypes.ConnectionEnd) {
	c.connections[connectionID] = connection
}

// GetConnection implements ConnectionKeeper
func (c ConnectionMock) GetConnection(_ sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool) {
	conn, ok := c.connections[connectionID]
	return conn, ok
}

// ChannelMock represents a mocked IBC channel keeper used for test purposes
type ChannelMock struct {
	channels map[string]channeltypes.Channel
}

// SetChannel sets a channel for mocking purpose
func (c ChannelMock) SetChannel(channelID string, channel channeltypes.Channel) {
	c.channels[channelID] = channel
}

// GetChannel implements ChannelKeeper
func (c ChannelMock) GetChannel(_ sdk.Context, _, channelID string) (channel channeltypes.Channel, found bool) {
	channel, ok := c.channels[channelID]
	return channel, ok
}

// GetNextSequenceSend implements ChannelKeeper
func (c ChannelMock) GetNextSequenceSend(_ sdk.Context, _, _ string) (uint64, bool) {
	return 0, false
}

// SendPacket implements ChannelKeeper
func (c ChannelMock) SendPacket(_ sdk.Context, _ *capabilitytypes.Capability, _ exported.PacketI) error {
	return nil
}

// WriteAcknowledgement implements ChannelKeeper
func (c ChannelMock) WriteAcknowledgement(_ sdk.Context, _ *capabilitytypes.Capability, _ exported.PacketI, _ []byte) error {
	return nil
}

// ChanCloseInit implements ChannelKeeper
func (c ChannelMock) ChanCloseInit(_ sdk.Context, _, _ string, _ *capabilitytypes.Capability) error {
	return nil
}
