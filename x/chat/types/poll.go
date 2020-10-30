package types

import (
	"errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewPoll creates a new poll
func NewPoll(options []string) (Poll, error) {
	var poll Poll

	if !checkOptions(options) {
		return poll, sdkerrors.Wrap(ErrInvalidPoll, "invalid options")
	}
	poll.Options = options
	poll.Votes = make(map[string]*Vote)

	return poll, nil
}

// HasUserVoted checks if a user voted for a poll
func (poll Poll) HasUserVoted(user string) (bool, error) {
	_, ok := poll.Votes[user]
	return ok, nil
}

// GetUserVote retrieves the vote of a user
func (poll Poll) GetUserVote(user string) (*Vote, error) {
	vote, ok := poll.Votes[user]
	if !ok {
		return nil, errors.New("No vote found")
	}

	return vote, nil
}

// AppendVote appends a vote into the poll
func (poll *Poll) AppendVote(vote *Vote) error {
	// Check if the vote value is valid
	if vote.Value >= int32(len(poll.Options)) {
		return errors.New("The vote value is not valid")
	}

	// Check if the user already voted
	hasVoted, err := poll.HasUserVoted(vote.Creator)
	if err != nil {
		return err
	}
	if hasVoted {
		return errors.New("The user already voted")
	}

	// Protobuf reset the map to nil if it has no value, therefore we must always check if it is initialized
	if poll.Votes == nil {
		poll.Votes = make(map[string]*Vote)
	}
	poll.Votes[vote.Creator] = vote

	return nil
}

// Check if the options provided for the poll have the right format
func checkOptions(options []string) bool {
	// TODO: Check the format of options, the number of options
	return true
}

// NewVote create a new vote
func NewVote(
	creator string,
	value int32,
	payload []byte,
) (Vote, error) {
	var vote Vote

	vote.Creator = creator
	vote.Value = value
	vote.Payload = payload

	return vote, nil
}
