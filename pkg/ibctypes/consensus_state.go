package ibctypes

import (
	"encoding/base64"
	"encoding/hex"
	committypes "github.com/cosmos/ibc-go/modules/core/23-commitment/types"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"time"
)

func NewConsensusState(timestamp, nextValSetHash, rootHash string) (*ibctmtypes.ConsensusState, error) {
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return nil, err
	}
	nextValSetHashBytes, _ := hex.DecodeString(nextValSetHash)
	rootHashBase64, _ := base64.StdEncoding.DecodeString(rootHash)
	return ibctmtypes.NewConsensusState(
		t,
		committypes.NewMerkleRoot(rootHashBase64),
		nextValSetHashBytes,
	), nil
}