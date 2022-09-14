package types_test

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/profile/types"
)

func TestGenesisState_Validate(t *testing.T) {
	var (
		addr1   = sample.Address(r)
		addr2   = sample.Address(r)
		addr3   = sample.Address(r)
		opAddr1 = sample.Address(r)
		opAddr2 = sample.Address(r)
		opAddr3 = sample.Address(r)
	)
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "should validate default genesis",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "should validate valid genesis",
			genState: &types.GenesisState{
				Validators: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
					{Address: addr2, OperatorAddresses: []string{opAddr2}},
					{Address: addr3, OperatorAddresses: []string{opAddr3}},
				},
				ValidatorsByOperatorAddress: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
					{OperatorAddress: opAddr3, ValidatorAddress: addr3},
				},
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
					{CoordinatorID: 2, Address: addr3},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: true},
					{CoordinatorID: 2, Address: addr3, Active: true},
				},
				CoordinatorCounter: 4,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "should prevent validate an invalid genesis",
			genState: &types.GenesisState{
				Validators: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorsByOperatorAddress: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
				},
			},
			valid: false,
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
		addr1   = sample.Address(r)
		addr2   = sample.Address(r)
		addr3   = sample.Address(r)
		opAddr1 = sample.Address(r)
		opAddr2 = sample.Address(r)
		opAddr3 = sample.Address(r)
	)
	tests := []struct {
		name     string
		genState *types.GenesisState
		err      error
	}{
		{
			name:     "should validate default state",
			genState: types.DefaultGenesis(),
		},
		{
			name: "should validate genesis with valid validators",
			genState: &types.GenesisState{
				Validators: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
					{Address: addr2, OperatorAddresses: []string{opAddr2}},
					{Address: addr3, OperatorAddresses: []string{opAddr3}},
				},
				ValidatorsByOperatorAddress: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
					{OperatorAddress: opAddr3, ValidatorAddress: addr3},
				},
			},
		},
		{
			name: "should prevent validate duplicated validators",
			genState: &types.GenesisState{
				Validators: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorsByOperatorAddress: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
				},
			},
			err: errors.New("duplicated index for validator"),
		},
		{
			name: "should prevent validate duplicated validator by operator address",
			genState: &types.GenesisState{
				Validators: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorsByOperatorAddress: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
				},
			},
			err: errors.New("duplicated index for validatorByOperatorAddress"),
		},
		{
			name: "should prevent validate missing validator address",
			genState: &types.GenesisState{
				Validators: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorsByOperatorAddress: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
				},
			},
			err: errors.New("validator operator address not found for Validator"),
		},
		{
			name: "should prevent validator missing validator operator address",
			genState: &types.GenesisState{
				Validators: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{}},
					{Address: addr2, OperatorAddresses: []string{opAddr2}},
				},
				ValidatorsByOperatorAddress: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
				},
			},
			err: errors.New("operator address not found in the Validator operator address list"),
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
		addr1 = sample.Address(r)
		addr2 = sample.Address(r)
		addr3 = sample.Address(r)
		addr4 = sample.Address(r)
	)
	tests := []struct {
		name     string
		genState *types.GenesisState
		err      error
	}{
		{
			name:     "should validate default genesis state",
			genState: types.DefaultGenesis(),
		},
		{
			name: "should validate genesis with valid coordinators",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
					{CoordinatorID: 2, Address: addr3},
					{CoordinatorID: 3, Address: addr4},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: true},
					{CoordinatorID: 2, Address: addr3, Active: true},
					{CoordinatorID: 3, Address: addr4, Active: true},
				},
				CoordinatorCounter: 4,
			},
		},
		{
			name: "should validate genesis with valid inactive coordinators",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: false},
				},
				CoordinatorCounter: 2,
			},
		},
		{
			name: "should prevent validate duplicated coordinators",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr1},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("duplicated index for coordinatorByAddress"),
		},
		{
			name: "should prevent validate duplicated coordinator id",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 0, Address: addr2},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 0, Address: addr2, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("duplicated id for coordinator"),
		},
		{
			name: "should prevent validate coordinator without a coordinator by address",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("coordinator address not found for CoordinatorByAddress"),
		},
		{
			name: "should prevent validate coordinator by address without a coordinator",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("coordinator address not found for coordinatorID"),
		},
		{
			name: "should prevent validate coordinator id higher than coordinator id counter",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 133, Address: addr2},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 133, Address: addr2, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("coordinator id should be lower or equal than the last id"),
		},
		{
			name: "should prevent validate inactive coordinator associated to a coordinator by address",
			genState: &types.GenesisState{
				CoordinatorsByAddress: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
				},
				Coordinators: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: false},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("coordinator found by CoordinatorByAddress should not be inactive"),
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
