// Package ibctypes defines to handle IBC handshakes in SPN modules
package ibctypes

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"time"

	committypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"gopkg.in/yaml.v2"
)

// RootHash returns the Merkle Root hash of the Consensus State
func (cs ConsensusState) RootHash() string {
	return cs.Root.Hash
}

// ParseConsensusStateFromFile parses a YAML dumped Consensus State file and
// returns a new Consensus State
func ParseConsensusStateFromFile(filePath string) (csf ConsensusState, err error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return csf, err
	}

	err = yaml.Unmarshal(f, &csf)
	return
}

// NewConsensusState initializes a new consensus state
func NewConsensusState(timestamp, nextValHash, rootHash string) ConsensusState {
	return ConsensusState{
		NextValHash: nextValHash,
		Timestamp:   timestamp,
		Root: MerkeRool{
			Hash: rootHash,
		},
	}
}

// ToTendermintConsensusState returns a new IBC Tendermint Consensus State
func (cs ConsensusState) ToTendermintConsensusState() (ibctmtypes.ConsensusState, error) {
	// parse the RFC3339 timestamp format
	t, err := time.Parse(time.RFC3339Nano, cs.Timestamp)
	if err != nil {
		return ibctmtypes.ConsensusState{}, err
	}

	// decode validator set
	nextValSetHashBytes, err := hex.DecodeString(cs.NextValHash)
	if err != nil {
		return ibctmtypes.ConsensusState{}, err
	}

	// decode root hash
	rootHashBase64, err := base64.StdEncoding.DecodeString(cs.RootHash())
	if err != nil {
		return ibctmtypes.ConsensusState{}, err
	}
	return *ibctmtypes.NewConsensusState(
		t,
		committypes.NewMerkleRoot(rootHashBase64),
		nextValSetHashBytes,
	), nil
}
