package types

import (
	"errors"

	types "github.com/cosmos/cosmos-sdk/codec/types"
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
func (poll Poll) HasUserVoted(user *User) (bool, error) {
	// Decode into an addressable user
	chatUser, err := user.DecodeChatUser()
	if err != nil {
		return false, err
	}

	_, ok := poll.Votes[chatUser.Identifier()]
	return ok, nil
}

// GetUserVote retrieves the vote of a user
func (poll Poll) GetUserVote(user *User) (*Vote, error) {
	// Decode into an addressable user
	chatUser, err := user.DecodeChatUser()
	if err != nil {
		return nil, err
	}

	vote, ok := poll.Votes[chatUser.Identifier()]
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

	// Decode into an addressable user
	chatUser, err := vote.Creator.DecodeChatUser()
	if err != nil {
		return err
	}

	// Get a string representation of the address
	addressString := chatUser.Identifier()

	poll.Votes[addressString] = vote

	return nil
}

// Check if the options provided for the poll have the right format
func checkOptions(options []string) bool {
	// TODO: Check the format of options, the number of options
	return true
}

// NewVote create a new vote
func NewVote(
	creator User,
	value int32,
	metadata *types.Any,
) (Vote, error) {
	var vote Vote

	vote.Creator = &creator
	vote.Value = value
	vote.Metadata = metadata

	return vote, nil
}
