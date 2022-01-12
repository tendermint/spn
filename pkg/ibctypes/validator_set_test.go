package ibctypes_test

import (
	"encoding/base64"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/pkg/ibctypes"
	"os"
	"strconv"
	"testing"
)

func TestParseValSetFile(t *testing.T) {
	t.Run("parse a dumped validator set", func(t *testing.T) {
		validatorSetYAML := `validators:
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
		f, err := os.CreateTemp("", "spn_validator_set_test")
		require.NoError(t, err)
		t.Cleanup(func() {
			f.Close()
			os.Remove(f.Name())
		})
		_, err = f.WriteString(validatorSetYAML)
		require.NoError(t, err)

		vsf, err := ibctypes.ParseValSetFile(f.Name())
		require.NoError(t, err)
		require.Len(t, vsf.Validators, 2)
		require.EqualValues(t, ibctypes.Validator{
			ProposerPriority: "0",
			VotingPower:      "100",
			PubKey: ibctypes.PubKey{
				Type:  ibctypes.TypeEd25519,
				Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
			},
		}, vsf.Validators[0])
		require.EqualValues(t, ibctypes.Validator{
			ProposerPriority: "1",
			VotingPower:      "50",
			PubKey: ibctypes.PubKey{
				Type:  ibctypes.TypeEd25519,
				Value: "/hO27XpCRWr6bZKqOxdNyYdLB3sAG2dG9dYXrOfM2II=",
			},
		}, vsf.Validators[1])
	})

	t.Run("non-existent file", func(t *testing.T) {
		_, err := ibctypes.ParseValSetFile("/foo/bar/foobar")
		require.Error(t, err)
	})

	t.Run("invalid file", func(t *testing.T) {
		validatorSetYAML := `foo`
		f, err := os.CreateTemp("", "spn_validator_set_test")
		require.NoError(t, err)
		t.Cleanup(func() {
			f.Close()
			os.Remove(f.Name())
		})
		_, err = f.WriteString(validatorSetYAML)
		require.NoError(t, err)

		_, err = ibctypes.ParseValSetFile(f.Name())
		require.Error(t, err)
	})
}

func TestNewValidatorSet(t *testing.T) {
	tests := []struct {
		name       string
		validators []ibctypes.Validator
		wantErr    bool
	}{
		{
			name: "return a new validator set",
			validators: []ibctypes.Validator{
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
			},
		},
		{
			name:       "prevent empty validator set",
			wantErr:    true,
			validators: []ibctypes.Validator{},
		},
		{
			name:    "prevent other key than ED25519",
			wantErr: true,
			validators: []ibctypes.Validator{
				{
					VotingPower:      "100",
					ProposerPriority: "0",
					PubKey: ibctypes.PubKey{
						Type:  "foo",
						Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					},
				},
			},
		},
		{
			name:    "prevent non-base 64 public key",
			wantErr: true,
			validators: []ibctypes.Validator{
				ibctypes.NewValidator(
					"foo",
					0,
					100,
				),
			},
		},
		{
			name:    "prevent invalid voting power",
			wantErr: true,
			validators: []ibctypes.Validator{
				{
					VotingPower:      "foo",
					ProposerPriority: "0",
					PubKey: ibctypes.PubKey{
						Type:  ibctypes.TypeEd25519,
						Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					},
				},
			},
		},
		{
			name:    "prevent invalid proposer priority",
			wantErr: true,
			validators: []ibctypes.Validator{
				{
					VotingPower:      "100",
					ProposerPriority: "foo",
					PubKey: ibctypes.PubKey{
						Type:  ibctypes.TypeEd25519,
						Value: "fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ibctypes.NewValidatorSet(tt.validators)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			// parse all validators
			require.Len(t, got.Validators, len(tt.validators))
			for i, v := range got.Validators {
				require.EqualValues(t,
					tt.validators[i].VotingPower,
					strconv.Itoa(int(v.VotingPower)),
				)
				require.EqualValues(t,
					tt.validators[i].ProposerPriority,
					strconv.Itoa(int(v.ProposerPriority)),
				)
				require.EqualValues(t,
					tt.validators[i].PubKey.Value,
					base64.StdEncoding.EncodeToString(v.PubKey.Bytes()),
				)
			}
		})
	}
}

func TestCheckValidatorSet(t *testing.T) {
	// first pair
	valSet1, err := ibctypes.NewValidatorSet([]ibctypes.Validator{
		ibctypes.NewValidator("fYaox+q+N3XkGZdcQ5f3MH4/5J4oh6FRoYdW0vxRdIg=", 0, 100),
	})
	require.NoError(t, err)
	consensusState1, err := ibctypes.NewConsensusState(
		"2022-01-12T12:25:19.523109Z",
		"48C4C20AC5A7BD99A45AEBAB92E61F5667253A2C51CCCD84D20327D3CB8737C9",
		"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	)
	require.NoError(t, err)

	// second pair
	valSet2, err := ibctypes.NewValidatorSet([]ibctypes.Validator{
		ibctypes.NewValidator("rQMyKjkzXXUhYsAdII6fSlTkFdf24hiSPGrSCBub5Oc=", 0, 100),
	})
	require.NoError(t, err)
	consensusState2, err := ibctypes.NewConsensusState(
		"2022-01-12T14:15:12.981874Z",
		"65BD4CB5502F7C926228F4A929E4FAF07384B3E5A0EC553A4230B8AB5A1022ED",
		"47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
	)
	require.NoError(t, err)

	require.True(t, ibctypes.CheckValidatorSet(valSet1, consensusState1))
	require.True(t, ibctypes.CheckValidatorSet(valSet2, consensusState2))
	require.False(t, ibctypes.CheckValidatorSet(valSet1, consensusState2))
	require.False(t, ibctypes.CheckValidatorSet(valSet2, consensusState1))
}
