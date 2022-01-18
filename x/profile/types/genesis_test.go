package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc:     "valid genesis state",
			genState: &types.GenesisState{
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestGenesisStateValidateValidator(t *testing.T) {
	var (
		addr1 = sample.Address()
		addr2 = sample.Address()
		addr3 = sample.Address()
	)
	tests := []struct {
		name     string
		genState *types.GenesisState
		err      error
	}{
		{
			name:     "default is valid",
			genState: types.DefaultGenesis(),
		},
		{
			name: "valid custom genesis",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, ConsensusAddress: addr1},
					{Address: addr2, ConsensusAddress: addr2},
					{Address: addr3, ConsensusAddress: addr3},
				},
				ValidatorByConsAddressList: []types.ValidatorByConsAddress{
					{ConsensusAddress: addr1, ValidatorAddress: addr1},
					{ConsensusAddress: addr2, ValidatorAddress: addr2},
					{ConsensusAddress: addr3, ValidatorAddress: addr3},
				},
			},
		},
		{
			name: "duplicated validator by address",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, ConsensusAddress: addr1},
					{Address: addr1, ConsensusAddress: addr1},
				},
				ValidatorByConsAddressList: []types.ValidatorByConsAddress{
					{ConsensusAddress: addr1, ValidatorAddress: addr1},
					{ConsensusAddress: addr2, ValidatorAddress: addr2},
				},
			},
			err: fmt.Errorf("duplicated index for validator: %s", addr1),
		},
		{
			name: "duplicated validator by consensus address",
			genState: &types.GenesisState{
				ValidatorByConsAddressList: []types.ValidatorByConsAddress{
					{ConsensusAddress: addr1, ValidatorAddress: addr1},
					{ConsensusAddress: addr1, ValidatorAddress: addr1},
				},
			},
			err: fmt.Errorf("duplicated index for validatorByConsAddress: %s", addr1),
		},
		{
			name: "missing validator by cons address",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, ConsensusAddress: addr1},
					{Address: addr2, ConsensusAddress: addr2},
					{Address: addr3, ConsensusAddress: addr3},
				},
				ValidatorByConsAddressList: []types.ValidatorByConsAddress{
					{ConsensusAddress: addr1, ValidatorAddress: addr1},
					{ConsensusAddress: addr2, ValidatorAddress: addr2},
				},
			},
			err: fmt.Errorf("validator consensus address not found for ValidatorByConsAddress: %s", addr3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.genState.ValidateValidators()
				if tt.err != nil {
					require.Error(t, err)
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.NoError(t, err)
			})
		})
	}
}

func TestGenesisStateValidateCoordinator(t *testing.T) {
	var (
		addr1 = sample.Address()
		addr2 = sample.Address()
		addr3 = sample.Address()
		addr4 = sample.Address()
	)
	tests := []struct {
		name     string
		genState *types.GenesisState
		err      error
	}{
		{
			name:     "default is valid",
			genState: types.DefaultGenesis(),
		},
		{
			name: "valid custom genesis",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
					{CoordinatorID: 2, Address: addr3},
					{CoordinatorID: 3, Address: addr4},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
					{CoordinatorID: 2, Address: addr3},
					{CoordinatorID: 3, Address: addr4},
				},
				CoordinatorCounter: 4,
			},
		},
		{
			name: "duplicated coordinator",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr1},
				},
				CoordinatorCounter: 2,
			},
			err: fmt.Errorf("duplicated index for coordinatorByAddress: %s", addr1),
		},
		{
			name: "duplicated coordinator id",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 0, Address: addr2},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 0, Address: addr2},
				},
				CoordinatorCounter: 2,
			},
			err: fmt.Errorf("duplicated id for coordinator: 0"),
		},
		{
			name: "profile not associated with chain",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
				},
				CoordinatorCounter: 2,
			},
			err: fmt.Errorf("coordinator address not found for CoordinatorByAddress: %s", addr2),
		},
		{
			name: "profile not associated with chain",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1},
				},
				CoordinatorCounter: 2,
			},
			err: fmt.Errorf("coordinator address not found for coordinatorID: 1"),
		},
		{
			name: "invalid coordinator id",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 133, Address: addr2},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 133, Address: addr2},
				},
				CoordinatorCounter: 2,
			},
			err: fmt.Errorf("coordinator id 133 should be lower or equal than the last id 2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.genState.ValidateCoordinators()
				if tt.err != nil {
					require.Error(t, err)
					require.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.NoError(t, err)
			})
		})
	}
}
