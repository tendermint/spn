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

// MerkeRool represents a Merkel Root in ConsensusStateFile
type MerkeRool struct {
	Hash string `yaml:"hash"`
}

// ConsensusStateFile represents a Consensus State dumped into a YAML file with command:
// appd q ibc client self-consensus-state
type ConsensusStateFile struct {
	NextValHash string    `yaml:"next_validators_hash"`
	Root        MerkeRool `yaml:"root"`
	Timestamp   string    `yaml:"timestamp"`
}

// RootHash returns the Merkle Root hash of the Consensus State
func (csf ConsensusStateFile) RootHash() string {
	return csf.Root.Hash
}

// ParseConsensusStateFile parses a YAML dumped Consensus State file and
// returns a new IBC Tendermint Consensus State
func ParseConsensusStateFile(filePath string) (csf ConsensusStateFile, err error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return csf, err
	}

	err = yaml.Unmarshal(f, &csf)
	return
}

// NewConsensusState returns a new IBC Tendermint Consensus State from string values
func NewConsensusState(timestamp, nextValSetHash, rootHash string) (ibctmtypes.ConsensusState, error) {
	// parse the RFC3339 timestamp format
	t, err := time.Parse(time.RFC3339Nano, timestamp)
	if err != nil {
		return ibctmtypes.ConsensusState{}, err
	}

	// decode validator set
	nextValSetHashBytes, err := hex.DecodeString(nextValSetHash)
	if err != nil {
		return ibctmtypes.ConsensusState{}, err
	}

	// decode root hash
	rootHashBase64, err := base64.StdEncoding.DecodeString(rootHash)
	if err != nil {
		return ibctmtypes.ConsensusState{}, err
	}
	return *ibctmtypes.NewConsensusState(
		t,
		committypes.NewMerkleRoot(rootHashBase64),
		nextValSetHashBytes,
	), nil
}
