package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	connectiontypes "github.com/cosmos/ibc-go/v3/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v3/modules/core/exported"
)

// Connection is an IBC connection end associated to a connection ID
type Connection struct {
	ConnID string
	Conn   connectiontypes.ConnectionEnd
}

// ConnectionMock represents a mocked IBC connection keeper used for test purposes
type ConnectionMock struct {
	connections map[string]connectiontypes.ConnectionEnd
}

// NewConnectionMock initializes a new connection mock
func NewConnectionMock(conns []Connection) (c ConnectionMock) {
	c.connections = make(map[string]connectiontypes.ConnectionEnd)
	for _, conn := range conns {
		c.connections[conn.ConnID] = conn.Conn
	}
	return
}

// GetConnection implements ConnectionKeeper
func (c ConnectionMock) GetConnection(_ sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool) {
	conn, ok := c.connections[connectionID]
	return conn, ok
}

// Channel is an IBC channel end associated to a channel ID
type Channel struct {
	ChannelID string
	Channel   channeltypes.Channel
}

// ChannelMock represents a mocked IBC channel keeper used for test purposes
type ChannelMock struct {
	channels map[string]channeltypes.Channel
}

// NewChannelMock initializes a new channel mock
func NewChannelMock(channels []Channel) (c ChannelMock) {
	c.channels = make(map[string]channeltypes.Channel)
	for _, channel := range channels {
		c.channels[channel.ChannelID] = channel.Channel
	}
	return
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
