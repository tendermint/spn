package keeper_test

import (
	"github.com/stretchr/testify/require"
	"testing"

	spnmocks "github.com/tendermint/spn/internal/testing"
)

func TestCreateChannel(t *testing.T) {
	ctx, k := spnmocks.MockChatContext()

	// Channel count is 0 at initialization
	require.Zero(t, k.GetChannelCount(ctx))

	// Cannot find a non existing channel
	_, found := k.GetChannel(ctx, 0)
	require.False(t, found)

	// A channel can be appended and retrieved
	newChannel := spnmocks.MockChannel()
	k.CreateChannel(ctx, newChannel)
	retrieved, found := k.GetChannel(ctx, 0)
	require.True(t, found)
	require.Equal(t, newChannel.Title, retrieved.Title)
	require.Equal(t, int32(1), k.GetChannelCount(ctx))

	// A second channel can be appended an retrieved
	newChannel = spnmocks.MockChannel()
	k.CreateChannel(ctx, newChannel)
	retrieved, found = k.GetChannel(ctx, 1)
	require.True(t, found)
	require.Equal(t, newChannel.Title, retrieved.Title)

	// Prevent a invalid user to create a channel
	newChannel = spnmocks.MockChannel()
	newChannel.Creator = "invalid_identifier"
	err := k.CreateChannel(ctx, newChannel)
	require.Error(t, err)

	// Can retrieve all the channels
	channels := k.GetAllChannels(ctx)
	require.Equal(t, 2, len(channels))
}
