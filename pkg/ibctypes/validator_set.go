package ibctypes

import (
	"encoding/base64"
	"errors"
	"fmt"
	ibctmtypes "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/types"
	"gopkg.in/yaml.v2"
	"os"
	"strconv"
)

const TypeEd25519 = "tendermint/PubKeyEd25519"

// PubKey represents a public key in Validator
type PubKey struct {
	Type  string `yaml:"type"`
	Value string `yaml:"value"`
}

// Validator represents a validator in ValSetFile
type Validator struct {
	ProposerPriority string `yaml:"proposer_priority"`
	PubKey           PubKey `yaml:"pub_key"`
	VotingPower      string `yaml:"voting_power"`
}

// NewValidator returns a validator with a ed25519 public key
func NewValidator(pubKey string, proposerPriority int64, votingPower int64) Validator {
	return Validator{
		ProposerPriority: strconv.Itoa(int(proposerPriority)),
		VotingPower:      strconv.Itoa(int(votingPower)),
		PubKey: PubKey{
			Type:  TypeEd25519,
			Value: pubKey,
		},
	}
}

// ValSetFile represents a Validator Set dumped into a YAML file with command:
// appd q tendermint-validator-set n
type ValSetFile struct {
	Validators []Validator `yaml:"validators"`
}

// ParseValSetFile parses a YAML dumped Validator Set file and returns a new Tendermint Validator Set
func ParseValSetFile(filePath string) (vsf ValSetFile, err error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return vsf, err
	}

	err = yaml.Unmarshal(f, &vsf)
	return
}

// NewValidatorSet returns a new Tendermint Validator Set from a list of validators
func NewValidatorSet(validators []Validator) (valSet tmtypes.ValidatorSet, err error) {
	if len(validators) == 0 {
		return tmtypes.ValidatorSet{}, errors.New("empty validator set")
	}

	for i, v := range validators {
		// convert the public key
		if v.PubKey.Type != TypeEd25519 {
			return valSet, fmt.Errorf(
				"validator %d: invalid key type: %s only %s is supported",
				i,
				v.PubKey.Type,
				TypeEd25519,
			)
		}

		keyBase64, err := base64.StdEncoding.DecodeString(v.PubKey.Value)
		if err != nil {
			return valSet, fmt.Errorf("validator %d: invalid public key %s", i, err.Error())
		}

		// create in the right format and add the validator
		votingPower, err := strconv.ParseInt(v.VotingPower, 10, 64)
		if err != nil {
			return valSet, fmt.Errorf("validator %d: invalid voting power", i)
		}
		proposerPriority, err := strconv.ParseInt(v.ProposerPriority, 10, 64)
		if err != nil {
			return valSet, fmt.Errorf("validator %d: invalid proposer priority", i)
		}
		v := tmtypes.NewValidator(ed25519.PubKey(keyBase64), votingPower)
		v.ProposerPriority = proposerPriority

		valSet.Validators = append(valSet.Validators, v)
	}

	return valSet, nil
}

// CheckValidatorSetHash checks the validator set hash matches the consensus state next validator set hash
func CheckValidatorSetHash(valSet tmtypes.ValidatorSet, consensusState ibctmtypes.ConsensusState) bool {
	nextValHash := base64.StdEncoding.EncodeToString(consensusState.NextValidatorsHash)
	valSetHash := base64.StdEncoding.EncodeToString(valSet.Hash())
	return nextValHash == valSetHash
}
