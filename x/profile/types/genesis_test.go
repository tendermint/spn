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
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
					{Address: addr2, OperatorAddresses: []string{opAddr2}},
					{Address: addr3, OperatorAddresses: []string{opAddr3}},
				},
				ValidatorByOperatorAddressList: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
					{OperatorAddress: opAddr3, ValidatorAddress: addr3},
				},
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
					{CoordinatorID: 2, Address: addr3},
				},
				CoordinatorList: []types.Coordinator{
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
			desc: "invalid genesis state",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorByOperatorAddressList: []types.ValidatorByOperatorAddress{
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
			name:     "default is valid",
			genState: types.DefaultGenesis(),
		},
		{
			name: "valid custom genesis",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
					{Address: addr2, OperatorAddresses: []string{opAddr2}},
					{Address: addr3, OperatorAddresses: []string{opAddr3}},
				},
				ValidatorByOperatorAddressList: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
					{OperatorAddress: opAddr3, ValidatorAddress: addr3},
				},
			},
		},
		{
			name: "duplicated validator by address",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorByOperatorAddressList: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
				},
			},
			err: errors.New("duplicated index for validator"),
		},
		{
			name: "duplicated validator by operator address",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorByOperatorAddressList: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
				},
			},
			err: errors.New("duplicated index for validatorByOperatorAddress"),
		},
		{
			name: "missing validator address",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{opAddr1}},
				},
				ValidatorByOperatorAddressList: []types.ValidatorByOperatorAddress{
					{OperatorAddress: opAddr1, ValidatorAddress: addr1},
					{OperatorAddress: opAddr2, ValidatorAddress: addr2},
				},
			},
			err: errors.New("validator operator address not found for Validator"),
		},
		{
			name: "missing operator address in the validator list",
			genState: &types.GenesisState{
				ValidatorList: []types.Validator{
					{Address: addr1, OperatorAddresses: []string{}},
					{Address: addr2, OperatorAddresses: []string{opAddr2}},
				},
				ValidatorByOperatorAddressList: []types.ValidatorByOperatorAddress{
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
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: true},
					{CoordinatorID: 2, Address: addr3, Active: true},
					{CoordinatorID: 3, Address: addr4, Active: true},
				},
				CoordinatorCounter: 4,
			},
		},
		{
			name: "valid with inactive coordinator",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: false},
				},
				CoordinatorCounter: 2,
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
			err: errors.New("duplicated index for coordinatorByAddress"),
		},
		{
			name: "duplicated coordinator id",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 0, Address: addr2},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 0, Address: addr2, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("duplicated id for coordinator"),
		},
		{
			name: "coordinator without a coordinator by address",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 1, Address: addr2, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("coordinator address not found for CoordinatorByAddress"),
		},
		{
			name: "coordinator by address without a coordinator",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("coordinator address not found for coordinatorID"),
		},
		{
			name: "invalid coordinator id",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 133, Address: addr2},
				},
				CoordinatorList: []types.Coordinator{
					{CoordinatorID: 0, Address: addr1, Active: true},
					{CoordinatorID: 133, Address: addr2, Active: true},
				},
				CoordinatorCounter: 2,
			},
			err: errors.New("coordinator id should be lower or equal than the last id"),
		},
		{
			name: "inactive coordinator associated to CoordinatorByAddress",
			genState: &types.GenesisState{
				CoordinatorByAddressList: []types.CoordinatorByAddress{
					{CoordinatorID: 0, Address: addr1},
					{CoordinatorID: 1, Address: addr2},
				},
				CoordinatorList: []types.Coordinator{
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
