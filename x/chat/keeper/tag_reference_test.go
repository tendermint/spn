package keeper_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/chat/types"

	"testing"
)

func TestGetTagReferencesFromChannel(t *testing.T) {
	ctx, k := spnmocks.MockChatContext()

	k.CreateChannel(ctx, spnmocks.MockChannel())
	message0ID := types.GetMessageIDFromChannelIDandIndex(0, 0)
	message1ID := types.GetMessageIDFromChannelIDandIndex(0, 1)
	message2ID := types.GetMessageIDFromChannelIDandIndex(0, 2)
	message3ID := types.GetMessageIDFromChannelIDandIndex(0, 3)
	message4ID := types.GetMessageIDFromChannelIDandIndex(1, 0)
	message5ID := types.GetMessageIDFromChannelIDandIndex(2, 0)

	// Cam retrieve references from tags in appended messages
	message := spnmocks.MockMessage(0)
	message.Tags = []string{"foo", "bar", "foo-bar"}
	k.AppendMessageToChannel(ctx, message)
	message.Tags = []string{"foo", "bar"}
	k.AppendMessageToChannel(ctx, message)
	message.Tags = []string{"bar"}
	k.AppendMessageToChannel(ctx, message)
	message.Tags = []string{"foo"}
	k.AppendMessageToChannel(ctx, message)
	message.Tags = []string{}
	k.AppendMessageToChannel(ctx, message)

	fooReferences := k.GetTagReferencesFromChannel(ctx, "foo", 0)
	require.Equal(t, 3, len(fooReferences))
	require.Equal(t, message0ID, fooReferences[0])
	require.Equal(t, message1ID, fooReferences[1])
	require.Equal(t, message3ID, fooReferences[2])

	barReferences := k.GetTagReferencesFromChannel(ctx, "bar", 0)
	require.Equal(t, 3, len(barReferences))
	require.Equal(t, message0ID, barReferences[0])
	require.Equal(t, message1ID, barReferences[1])
	require.Equal(t, message2ID, barReferences[2])

	foobarReferences := k.GetTagReferencesFromChannel(ctx, "foo-bar", 0)
	require.Equal(t, 1, len(foobarReferences))
	require.Equal(t, message0ID, foobarReferences[0])

	barfooReferences := k.GetTagReferencesFromChannel(ctx, "bar-foo", 0)
	require.Equal(t, 0, len(barfooReferences))

	// Can get all tag references in all channels
	k.CreateChannel(ctx, spnmocks.MockChannel())
	k.CreateChannel(ctx, spnmocks.MockChannel())
	message = spnmocks.MockMessage(1)
	message.Tags = []string{"foo"}
	k.AppendMessageToChannel(ctx, message)
	message = spnmocks.MockMessage(2)
	message.Tags = []string{"foo"}
	k.AppendMessageToChannel(ctx, message)

	fooReferences = k.GetAllTagReferences(ctx, "foo")
	require.Equal(t, 5, len(fooReferences))
	require.Equal(t, message0ID, fooReferences[0])
	require.Equal(t, message1ID, fooReferences[1])
	require.Equal(t, message3ID, fooReferences[2])
	require.Equal(t, message4ID, fooReferences[3])
	require.Equal(t, message5ID, fooReferences[4])
}
