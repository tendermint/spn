package keeper_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/tendermint/spn/x/chat"
)

func TestAppendChannel(t *testing.T) {
	ctx, k := chat.MockContext()

	// Channel count is 0 at initialization
	require.Zero(t, k.GetChannelCount(ctx), "Channel count must be 0 at initialization")

	// Cannot find a non existing channel
	_, found := k.GetChannel(ctx, 0)
	require.False(t, found, "GetChannel should return found to as false if the channel doesn't exist")

	// A channel can be appended and retrieved
	newChannel := chat.MockChannel()
	k.AppendChannel(ctx, newChannel)
	retrieved, found := k.GetChannel(ctx, 0)
	require.True(t, found, "An appended channel should be retrieved")
	require.Equal(t, newChannel.Name, retrieved.Name, "GetChannel should retrieve the appended channel")
	require.Equal(t, int32(1), k.GetChannelCount(ctx), "Channel count must be 1 after a channel has been appended")

	// A second channel can be appended an retrieved
	newChannel = chat.MockChannel()
	k.AppendChannel(ctx, newChannel)
	retrieved, found = k.GetChannel(ctx, 1)
	require.True(t, found, "An appended channel should be retrieved")
	require.Equal(t, newChannel.Name, retrieved.Name, "GetChannel should retrieve the appended channel")
	require.Equal(t, int32(2), k.GetChannelCount(ctx), "Channel count must be 2 after two channels have been appended")
}
