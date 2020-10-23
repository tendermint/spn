package types_test

import (
	"github.com/google/go-cmp/cmp"
	"github.com/tendermint/spn/x/chat"
	"github.com/tendermint/spn/x/chat/types"
	"testing"
)

func TestNewPoll(t *testing.T) {
	// Can create a poll
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, err := types.NewPoll(pollOptions)
	if err != nil {
		t.Errorf("NewPoll should create a new poll: %v", err)
	}
	if !cmp.Equal(pollOptions, pollOptions) {
		t.Errorf("NewPoll should create a new poll with options provded")
	}
	if len(poll.Votes) > 0 {
		t.Errorf("NewPoll should create a new poll with no vote")
	}
}

func TestAppendVote(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := chat.MockUser()

	// Allow to append a vote
	vote, _ := types.NewVote(user, 0, nil)
	err := poll.AppendVote(&vote)
	if err != nil {
		t.Errorf("AppendVote should append a vote: %v", err)
	}
	if len(poll.Votes) != 1 {
		t.Errorf("AppendVote should append a vote")
	}

	// Allow to append a vote from another user
	user2 := chat.MockUser()
	vote2, _ := types.NewVote(user2, 2, nil)
	err = poll.AppendVote(&vote2)
	if err != nil {
		t.Errorf("AppendVote should append a second vote: %v", err)
	}
	if len(poll.Votes) != 2 {
		t.Errorf("AppendVote should append a second vote")
	}

	// Prevent appending invalid valid
	user3 := chat.MockUser()
	vote3, _ := types.NewVote(user3, int32(len(pollOptions)), nil)
	err = poll.AppendVote(&vote3)
	if err == nil {
		t.Errorf("AppendVote should prevent append an invalid vote")
	}

	// Prevent appending a vote from the same user
	vote4, _ := types.NewVote(user, 0, nil)
	err = poll.AppendVote(&vote4)
	if err == nil {
		t.Errorf("AppendVote should prevent append a vote from a user who already voted")
	}
}

func TestHasUserVoted(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := chat.MockUser()

	// Should return false if no vote
	voted, err := poll.HasUserVoted(&user)
	if err != nil {
		t.Errorf("HasUserVoted should return false: %v", err)
	}
	if voted {
		t.Errorf("HasUserVoted should return false")
	}

	// Should return true if voted
	vote, _ := types.NewVote(user, 0, nil)
	poll.AppendVote(&vote)
	voted, err = poll.HasUserVoted(&user)
	if err != nil {
		t.Errorf("HasUserVoted with vote should return true: %v", err)
	}
	if !voted {
		t.Errorf("HasUserVoted with vote should return true")
	}
}

func TestGetUserVote(t *testing.T) {
	pollOptions := []string{"coffee", "tea", "water", "beer"}
	poll, _ := types.NewPoll(pollOptions)
	user := chat.MockUser()

	// Should return error if no vote
	_, err := poll.GetUserVote(&user)
	if err == nil {
		t.Errorf("GetUserVote with no vote should return an error")
	}

	// Should return the vote
	vote, _ := types.NewVote(user, 0, nil)
	poll.AppendVote(&vote)
	retrieved, err := poll.GetUserVote(&user)
	if err != nil {
		t.Errorf("GetUserVote should return the vote: %v", err)
	}
	if !cmp.Equal(*retrieved, vote) {
		t.Errorf("GetUserVote should return %v, got: %v", vote, retrieved)
	}
}
