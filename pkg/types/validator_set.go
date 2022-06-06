// Package types defines types to handle IBC handshakes in SPN modules
package types

import (
	"encoding/base64"
	"fmt"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"

	ibctmtypes "github.com/cosmos/ibc-go/v3/modules/light-clients/07-tendermint/types"
	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/types"
)

const TypeEd25519 = "tendermint/PubKeyEd25519"

// validatorSetFile represents dumped validator set from
// q tendermint-validator-set
type validatorSetFile struct {
	Validators []struct {
		ProposerPriority string `yaml:"proposer_priority"`
		VotingPower      string `yaml:"voting_power"`
		PubKey           struct {
			Type  string `yaml:"type"`
			Value string `yaml:"value"`
		} `yaml:"pub_key"`
	} `yaml:"validators"`
}

// ParseValidatorSetFromFile parses a YAML dumped Validator Set file and returns a new Validator Set
// TODO: Support other format than YAML if there are other format of dumped file
func ParseValidatorSetFromFile(filePath string) (vs ValidatorSet, err error) {
	// parse file
	var vsf validatorSetFile
	f, err := os.ReadFile(filePath)
	if err != nil {
		return vs, err
	}
	err = yaml.Unmarshal(f, &vsf)
	if err != nil {
		return vs, err
	}

	// convert
	for i, v := range vsf.Validators {
		proposerPriority, err := strconv.ParseInt(v.ProposerPriority, 10, 64)
		if err != nil {
			return vs, errors.Wrapf(err, "invalid validator %d proposer priority", i)
		}
		votingPower, err := strconv.ParseInt(v.VotingPower, 10, 64)
		if err != nil {
			return vs, errors.Wrapf(err, "invalid validator %d voting power", i)
		}
		vs.Validators = append(vs.Validators, NewValidator(v.PubKey.Value, proposerPriority, votingPower))
	}

	return
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

// NewValidatorSet returns a new Validator Set from a list of validators
func NewValidatorSet(validators ...Validator) ValidatorSet {
	return ValidatorSet{
		Validators: validators,
	}
}

// ToTendermintValidatorSet returns a new Tendermint Validator Set
func (vs ValidatorSet) ToTendermintValidatorSet() (valSet tmtypes.ValidatorSet, err error) {
	if len(vs.Validators) == 0 {
		return valSet, errors.New("empty validator set")
	}

	for i, v := range vs.Validators {
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

		// ValidateBasic() method of Validator Set type requires to have a non-nil proposer
		// We add the first validator as proposer to ensure the method succeeds
		// This value doesn't influence the hash verification of the validator set
		if valSet.Proposer == nil {
			valSet.Proposer = v
		}
	}

	valSet.Validators[0].PubKey.Address()
	return valSet, nil
}

// CheckValidatorSetHash checks the Tendermint validator set hash matches the Tendermint consensus state next validator set hash
func CheckValidatorSetHash(valSet tmtypes.ValidatorSet, consensusState ibctmtypes.ConsensusState) bool {
	nextValHash := base64.StdEncoding.EncodeToString(consensusState.NextValidatorsHash)
	valSetHash := base64.StdEncoding.EncodeToString(valSet.Hash())
	return nextValHash == valSetHash
}
