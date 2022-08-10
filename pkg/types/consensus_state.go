// Package types defines types to handle IBC handshakes in SPN modules
package types

import (
	"encoding/base64"
	"encoding/hex"
	"os"
	"time"

	"gopkg.in/yaml.v2"

	committypes "github.com/cosmos/ibc-go/v5/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/v5/modules/light-clients/07-tendermint/types"
)

// consensusStateFile represents dumped consensus state from
// q ibc client self-consensus-state
type consensusStateFile struct {
	NextValidatorsHash string `yaml:"next_validators_hash"`
	Timestamp          string `yaml:"timestamp"`
	Root               struct {
		Hash string `yaml:"hash"`
	}
}

// RootHash returns the Merkle Root hash of the Consensus State
func (cs ConsensusState) RootHash() string {
	return cs.Root.Hash
}

// ParseConsensusStateFromFile parses a YAML dumped Consensus State file and
// returns a new Consensus State
// TODO: Improve method and support other format than YAML if there are other format of dumped file
func ParseConsensusStateFromFile(filePath string) (ConsensusState, error) {
	// parse file
	var csf consensusStateFile
	f, err := os.ReadFile(filePath)
	if err != nil {
		return ConsensusState{}, err
	}
	err = yaml.Unmarshal(f, &csf)

	// convert
	cs := NewConsensusState(csf.Timestamp, csf.NextValidatorsHash, csf.Root.Hash)
	return cs, err
}

// NewConsensusState initializes a new consensus state
func NewConsensusState(timestamp, nextValHash, rootHash string) ConsensusState {
	return ConsensusState{
		NextValidatorsHash: nextValHash,
		Timestamp:          timestamp,
		Root: MerkleRoot{
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
	nextValSetHashBytes, err := hex.DecodeString(cs.NextValidatorsHash)
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
