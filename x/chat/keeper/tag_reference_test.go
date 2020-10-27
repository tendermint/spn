package keeper_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"

	"testing"
)

func TestGetTagReferencesFromChannel(t *testing.T) {
	ctx, k := chat.MockContext()

	k.AppendChannel(ctx, chat.MockChannel())
	message0ID := types.GetMessageIDFromChannelIDandIndex(0, 0)
	message1ID := types.GetMessageIDFromChannelIDandIndex(0, 1)
	message2ID := types.GetMessageIDFromChannelIDandIndex(0, 2)
	message3ID := types.GetMessageIDFromChannelIDandIndex(0, 3)
	message4ID := types.GetMessageIDFromChannelIDandIndex(1, 0)
	message5ID := types.GetMessageIDFromChannelIDandIndex(2, 0)

	// Cam retrieve references from tags in appended messages
	message := chat.MockMessage(0)
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
	require.Equal(t, 3, len(fooReferences), "GetTagReferencesFromChannel should find 3 foo references")
	require.Equal(t, message0ID, fooReferences[0], "GetTagReferencesFromChannel should return reference message ID")
	require.Equal(t, message1ID, fooReferences[1], "GetTagReferencesFromChannel should return reference message ID")
	require.Equal(t, message3ID, fooReferences[2], "GetTagReferencesFromChannel should return reference message ID")

	barReferences := k.GetTagReferencesFromChannel(ctx, "bar", 0)
	require.Equal(t, 3, len(barReferences), "GetTagReferencesFromChannel should find 3 bar references")
	require.Equal(t, message0ID, barReferences[0], "GetTagReferencesFromChannel should return reference message ID")
	require.Equal(t, message1ID, barReferences[1], "GetTagReferencesFromChannel should return reference message ID")
	require.Equal(t, message2ID, barReferences[2], "GetTagReferencesFromChannel should return reference message ID")

	foobarReferences := k.GetTagReferencesFromChannel(ctx, "foo-bar", 0)
	require.Equal(t, 1, len(foobarReferences), "GetTagReferencesFromChannel should find 1 foo-bar reference")
	require.Equal(t, message0ID, foobarReferences[0], "GetTagReferencesFromChannel should return reference message ID")

	barfooReferences := k.GetTagReferencesFromChannel(ctx, "bar-foo", 0)
	require.Equal(t, 0, len(barfooReferences), "GetTagReferencesFromChannel should find 0 bar-foo reference")

	// Can get all tag references in all channels
	k.AppendChannel(ctx, chat.MockChannel())
	k.AppendChannel(ctx, chat.MockChannel())
	message = chat.MockMessage(1)
	message.Tags = []string{"foo"}
	k.AppendMessageToChannel(ctx, message)
	message = chat.MockMessage(2)
	message.Tags = []string{"foo"}
	k.AppendMessageToChannel(ctx, message)

	fooReferences = k.GetAllTagReferences(ctx, "foo")
	require.Equal(t, 5, len(fooReferences), "GetAllTagReferences should find 5 foo references")
	require.Equal(t, message0ID, fooReferences[0], "GetAllTagReferences should return reference message ID")
	require.Equal(t, message1ID, fooReferences[1], "GetAllTagReferences should return reference message ID")
	require.Equal(t, message3ID, fooReferences[2], "GetAllTagReferences should return reference message ID")
	require.Equal(t, message4ID, fooReferences[3], "GetAllTagReferences should return reference message ID")
	require.Equal(t, message5ID, fooReferences[4], "GetAllTagReferences should return reference message ID")
}
