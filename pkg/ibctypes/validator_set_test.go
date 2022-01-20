package ibctypes_test

import (
	"encoding/base64"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/ibctypes"
)

func TestParseValidatorSetFromFile(t *testing.T) {
	fileFromContent := func(content string) string {
		f, err := os.CreateTemp("", "spn_validator_set_test")
		require.NoError(t, err)
		t.Cleanup(func() {
			os.Remove(f.Name())
		})
		_, err = f.WriteString(content)
		require.NoError(t, err)
		require.NoError(t, f.Close())
		return f.Name()
	}

	validValidatorSet := `validators:
- proposer_priority: "0"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=
  voting_power: "100"
- proposer_priority: "1"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: /hO27XpCRWr6bZKqOxdNyYdLB3sAG2dG9dYXrOfM2II=
  voting_power: "50"
`
	validatorSetInvalidVotingPower := `validators:
- proposer_priority: "0"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=
  voting_power: "foo""
- proposer_priority: "1"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: /hO27XpCRWr6bZKqOxdNyYdLB3sAG2dG9dYXrOfM2II=
  voting_power: "50"
`
	validatorSetInvalidProposerPriority := `validators:
- proposer_priority: "0"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=
  voting_power: "100""
- proposer_priority: "foo"
  pub_key:
    type: tendermint/PubKeyEd25519
    value: /hO27XpCRWr6bZKqOxdNyYdLB3sAG2dG9dYXrOfM2II=
  voting_power: "50"
`
	tests := []struct {
		name     string
		filename string
		expected ibctypes.ValidatorSet
		wantErr  bool
	}{
		{
			name:     "parse a dumped validator set",
			filename: fileFromContent(validValidatorSet),
			expected: ibctypes.NewValidatorSet(
				ibctypes.Validator{
					ProposerPriority: "0",
					VotingPower:      "100",
					PubKey: ibctypes.PubKey{
						Type:  ibctypes.TypeEd25519,
						Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					},
				},
				ibctypes.Validator{
					ProposerPriority: "1",
					VotingPower:      "50",
					PubKey: ibctypes.PubKey{
						Type:  ibctypes.TypeEd25519,
						Value: "/hO27XpCRWr6bZKqOxdNyYdLB3sAG2dG9dYXrOfM2II=",
					},
				},
			),
		},
		{
			name:     "invalid voting power",
			filename: fileFromContent(validatorSetInvalidVotingPower),
			wantErr:  true,
		},
		{
			name:     "invalid proposer priority",
			filename: fileFromContent(validatorSetInvalidProposerPriority),
			wantErr:  true,
		},
		{
			name:     "non-existent file",
			filename: "/foo/bar/foobar",
			wantErr:  true,
		},
		{
			name:     "invalid file",
			filename: fileFromContent("foo"),
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vs, err := ibctypes.ParseValidatorSetFromFile(tt.filename)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Len(t, vs.Validators, len(tt.expected.Validators))
			require.EqualValues(t, vs, tt.expected)
		})
	}
}

func TestValidatorSet_ToTendermintValidatorSet(t *testing.T) {
	tests := []struct {
		name         string
		validatorSet ibctypes.ValidatorSet
		wantErr      bool
	}{
		{
			name: "return a new validator set",
			validatorSet: ibctypes.NewValidatorSet(
				ibctypes.NewValidator(
					"fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					0,
					100,
				),
				ibctypes.NewValidator(
					"/hO27XpCRWr6bZKqOxdNyYdLB3sAG2dG9dYXrOfM2II=",
					1,
					50,
				),
			),
		},
		{
			name:         "prevent empty validator set",
			wantErr:      true,
			validatorSet: ibctypes.NewValidatorSet(),
		},
		{
			name:    "prevent other key than ED25519",
			wantErr: true,
			validatorSet: ibctypes.NewValidatorSet(
				ibctypes.Validator{
					VotingPower:      "100",
					ProposerPriority: "0",
					PubKey: ibctypes.PubKey{
						Type:  "foo",
						Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					},
				},
			),
		},
		{
			name:    "prevent non-base 64 public key",
			wantErr: true,
			validatorSet: ibctypes.NewValidatorSet(
				ibctypes.NewValidator(
					"foo",
					0,
					100,
				),
			),
		},
		{
			name:    "prevent invalid voting power",
			wantErr: true,
			validatorSet: ibctypes.NewValidatorSet(
				ibctypes.Validator{
					VotingPower:      "foo",
					ProposerPriority: "0",
					PubKey: ibctypes.PubKey{
						Type:  ibctypes.TypeEd25519,
						Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					},
				},
			),
		},
		{
			name:    "prevent invalid proposer priority",
			wantErr: true,
			validatorSet: ibctypes.NewValidatorSet(
				ibctypes.Validator{
					VotingPower:      "100",
					ProposerPriority: "foo",
					PubKey: ibctypes.PubKey{
						Type:  ibctypes.TypeEd25519,
						Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					},
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.validatorSet.ToTendermintValidatorSet()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, got.ValidateBasic(), "the converted type should be valid")
			// parse all validators
			require.Len(t, got.Validators, len(tt.validatorSet.Validators))
			for i, v := range got.Validators {
				require.EqualValues(t,
					tt.validatorSet.Validators[i].VotingPower,
					strconv.Itoa(int(v.VotingPower)),
				)
				require.EqualValues(t,
					tt.validatorSet.Validators[i].ProposerPriority,
					strconv.Itoa(int(v.ProposerPriority)),
				)
				require.EqualValues(t,
					tt.validatorSet.Validators[i].PubKey.Value,
					base64.StdEncoding.EncodeToString(v.PubKey.Bytes()),
				)
			}
		})
	}
}

func TestCheckValidatorSet(t *testing.T) {
	// first pair
	valSet1 := ibctypes.NewValidatorSet(
		ibctypes.NewValidator("fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=", 0, 100),
	)
	tmValSet1, err := valSet1.ToTendermintValidatorSet()
	require.NoError(t, err)
	consensusState1 := ibctypes.NewConsensusState(
		"2022-01-12T12:25:19.523109Z",
		"48C4C20AC5A7BD99A45AEBAB92E61F5667253A2C51CCCD84D20327D3CB8737C9",
		"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	)
	tmConsensusState1, err := consensusState1.ToTendermintConsensusState()
	require.NoError(t, err)

	// second pair
	valSet2 := ibctypes.NewValidatorSet(
		ibctypes.NewValidator("rQMyKjkzXXUhYsAdII6fSlTkFdf24hiSPGrSCBub5Oc=", 0, 100),
	)
	tmValSet2, err := valSet2.ToTendermintValidatorSet()
	require.NoError(t, err)
	consensusState2 := ibctypes.NewConsensusState(
		"2022-01-12T14:15:12.981874Z",
		"65BD4CB5502F7C926228F4A929E4FAF07384B3E5A0EC553A4230B8AB5A1022ED",
		"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	)
	tmConsensusState2, err := consensusState2.ToTendermintConsensusState()
	require.NoError(t, err)

	require.True(t, ibctypes.CheckValidatorSetHash(tmValSet1, tmConsensusState1))
	require.True(t, ibctypes.CheckValidatorSetHash(tmValSet2, tmConsensusState2))
	require.False(t, ibctypes.CheckValidatorSetHash(tmValSet1, tmConsensusState2))
	require.False(t, ibctypes.CheckValidatorSetHash(tmValSet2, tmConsensusState1))
}
