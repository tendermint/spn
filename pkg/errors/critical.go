package errors

import (
	"errors"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var ErrCritical = errors.New("CRITICAL: the state of the blockchain is inconsistent or an invariant is broken")

// Critical handles and/or returns an error in case a critical error has been encountered:
// - Inconsistent state
// - Broken invariant
func Critical(description string) error {
	return sdkerrors.Wrap(ErrCritical, description)
}
