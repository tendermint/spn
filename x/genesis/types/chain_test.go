package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/x/genesis/types"
)

func TestNewChain(t *testing.T) {
	// Can create a chain
	chain, err := types.NewChain(
		"chainID",
		"creator",
		"sourceURL",
		"sourceHash",
		time.Now(),
	)

	// Can append peers to the chain

	// Prevent creating a chain with a invalid name

}
