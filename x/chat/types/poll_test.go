package types_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
)

func TestNewPoll(t *testing.T) {
	// Can create a poll
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, err := types.NewPoll(pollOptions)
	require.NoError(t, err, "NewPoll should create a new poll")
	require.Equal(t, pollOptions, poll.Options, "NewPoll should create a new poll with options provided")
	require.Zero(t, len(poll.Votes), "NewPoll should create a new poll with no vote")
}

func TestNewVote(t *testing.T) {
	user := chat.MockUser()

	// Can create a vote with payload
	payload := chat.MockPayload()
	_, err := types.NewVote(user, 0, payload)
	require.NoError(t, err, "NewVote should create a new vote")
}

func TestAppendVote(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := chat.MockUser()

	// Allow to append a vote
	vote, _ := types.NewVote(user, 0, nil)
	err := poll.AppendVote(&vote)
	require.NoError(t, err, "AppendVote should append a vote")
	require.Equal(t, 1, len(poll.Votes), "AppendVote should append a vote")

	// Allow to append a vote from another user
	user2 := chat.MockUser()
	vote2, _ := types.NewVote(user2, 2, nil)
	err = poll.AppendVote(&vote2)
	require.NoError(t, err, "AppendVote should append a second vote")
	require.Equal(t, 2, len(poll.Votes), "AppendVote should append a second vote")

	// Prevent appending invalid valid
	user3 := chat.MockUser()
	vote3, _ := types.NewVote(user3, int32(len(pollOptions)), nil)
	err = poll.AppendVote(&vote3)
	require.Error(t, err, "AppendVote should prevent append an invalid vote")

	// Prevent appending a vote from the same user
	vote4, _ := types.NewVote(user, 0, nil)
	err = poll.AppendVote(&vote4)
	require.Error(t, err, "AppendVote should prevent append a vote from a user who already voted")
}

func TestHasUserVoted(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := chat.MockUser()

	// Should return false if no vote
	voted, err := poll.HasUserVoted(user)
	require.NoError(t, err, "HasUserVoted should return false")
	require.False(t, voted, "HasUserVoted should return false")

	// Should return true if voted
	vote, _ := types.NewVote(user, 0, nil)
	poll.AppendVote(&vote)
	voted, err = poll.HasUserVoted(user)
	require.NoError(t, err, "HasUserVoted with vote should return true")
	require.True(t, voted, "HasUserVoted with vote should return true")

}

func TestGetUserVote(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := chat.MockUser()

	// Should return error if no vote
	_, err := poll.GetUserVote(user)
	require.Error(t, err, "GetUserVote with no vote should return an error")

	// Should return the vote
	vote, _ := types.NewVote(user, 0, nil)
	poll.AppendVote(&vote)
	retrieved, err := poll.GetUserVote(user)
	require.NoError(t, err, "GetUserVote should return the vote")
	require.Equal(t, vote, *retrieved, "Not equal")

}
