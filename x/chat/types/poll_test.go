package types_test

import (
	"github.com/stretchr/testify/require"
	spnmocks "github.com/tendermint/spn/internal/testing"
	"github.com/tendermint/spn/x/chat/types"

	"testing"
)

func TestNewPoll(t *testing.T) {
	// Can create a poll
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, err := types.NewPoll(pollOptions)
	require.NoError(t, err)
	require.Equal(t, pollOptions, poll.Options)
	require.Zero(t, len(poll.Votes))
}

func TestNewVote(t *testing.T) {
	user := spnmocks.MockUser()

	// Can create a vote with payload
	payload := spnmocks.MockPayload()
	_, err := types.NewVote(user, 0, payload)
	require.NoError(t, err)
}

func TestAppendVote(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := spnmocks.MockUser()

	// Allow to append a vote
	vote, _ := types.NewVote(user, 0, nil)
	err := poll.AppendVote(&vote)
	require.NoError(t, err)
	require.Equal(t, 1, len(poll.Votes))

	// Allow to append a vote from another user
	user2 := spnmocks.MockUser()
	vote2, _ := types.NewVote(user2, 2, nil)
	err = poll.AppendVote(&vote2)
	require.NoError(t, err)
	require.Equal(t, 2, len(poll.Votes))

	// Prevent appending invalid valid
	user3 := spnmocks.MockUser()
	vote3, _ := types.NewVote(user3, int32(len(pollOptions)), nil)
	err = poll.AppendVote(&vote3)
	require.Error(t, err)

	// Prevent appending a vote from the same user
	vote4, _ := types.NewVote(user, 0, nil)
	err = poll.AppendVote(&vote4)
	require.Error(t, err)
}

func TestHasUserVoted(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := spnmocks.MockUser()

	// Should return false if no vote
	voted, err := poll.HasUserVoted(user)
	require.NoError(t, err)
	require.False(t, voted)

	// Should return true if voted
	vote, _ := types.NewVote(user, 0, nil)
	poll.AppendVote(&vote)
	voted, err = poll.HasUserVoted(user)
	require.NoError(t, err)
	require.True(t, voted)

}

func TestGetUserVote(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := spnmocks.MockUser()

	// Should return error if no vote
	_, err := poll.GetUserVote(user)
	require.Error(t, err)

	// Should return the vote
	vote, _ := types.NewVote(user, 0, nil)
	poll.AppendVote(&vote)
	retrieved, err := poll.GetUserVote(user)
	require.NoError(t, err)
	require.Equal(t, vote, *retrieved)

}
