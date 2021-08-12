package errors

import (
	"errors"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var ErrCritical = errors.New("CRITICAL: the state of the blockchain is inconsistent or an invariant is broken")

// Error handles and/or returns an error in case a critical error has been encountered:
// - Inconsistent state
// - Broken invariant
func Critical(description string) error {
	return sdkerrors.Wrap(ErrCritical, description)
}

// Criticalf extends a critical error with additional information.
//
// This function works like the Critical function with additional
// functionality of formatting the input as specified.
func Criticalf(format string, args ...interface{}) error {
	desc := fmt.Sprintf(format, args...)
	return Critical(desc)
}
