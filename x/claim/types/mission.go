package types

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Validate checks the mission is valid
func (m Mission) Validate() error {

	if m.Weight.LT(sdk.ZeroDec()) || m.Weight.GT(sdk.OneDec()) {
		return errors.New("mission weight must be in range [0:1]")
	}

	return nil
}
