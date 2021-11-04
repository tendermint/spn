// Package chainid defines helper methods to manipulate and verify chain id used in Cosmos-SDK blockchain genesis
package chainid

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	ChainIDSeparator   = "-"
	ChainNameMaxLength = 30
)

var ErrInvalidChainID = errors.New("invalid chain ID, format must be <chain-name>-<network-number>")

// NewGenesisChainID returns the genesis chain id from the chain name and the number
func NewGenesisChainID(chainName string, networkNumber uint64) string {
	return fmt.Sprintf("%s%s%d", chainName, ChainIDSeparator, networkNumber)
}

// ParseGenesisChainID returns the chain name and the number from the chain ID
// The chain ID follows the format <ChainName>-<Number>
// The function returns an error if the chain ID is invalid
func ParseGenesisChainID(genesisChainID string) (string, uint64, error) {
	parsed := strings.Split(genesisChainID, ChainIDSeparator)
	if len(parsed) != 2 {
		return "", 0, errors.Wrapf(ErrInvalidChainID, "%s: invalid format", genesisChainID)
	}
	if err := CheckChainName(parsed[0]); err != nil {
		return "", 0, errors.Wrapf(ErrInvalidChainID, "%s: invalid chain name %s", parsed[0], err.Error())
	}
	number, err := strconv.ParseUint(parsed[1], 10, 64)
	if err != nil {
		return "", 0, errors.Wrapf(ErrInvalidChainID, "%s: invalid network id %s", parsed[1], err.Error())
	}

	return parsed[0], number, nil
}

// CheckChainName verifies the name is valid as a chain name
func CheckChainName(chainName string) error {
	if len(chainName) == 0 {
		return errors.New("chain name can't be empty")
	}

	if len(chainName) > ChainNameMaxLength {
		return fmt.Errorf("chain name is too long, max length is %d", ChainNameMaxLength)
	}

	// Iterate characters
	for _, c := range chainName {
		if !isChainAuthorizedChar(c) {
			return errors.New("chain name must be lowercase alpha character")
		}
	}

	return nil
}

// isChainAuthorizedChar checks to ensure that all letters in the chain name are valid
func isChainAuthorizedChar(c rune) bool {
	return 'a' <= c && c <= 'z'
}
