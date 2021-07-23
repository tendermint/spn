package types

import (
	"fmt"
	"github.com/tendermint/spn/testutil/sample"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenesisStateValidateValidator(t *testing.T) {
	var (
		consAddr1 = sample.AccAddress()
		consAddr2 = sample.AccAddress()
		consAddr3 = sample.AccAddress()
		consAddr4 = sample.AccAddress()
		addr1     = sample.AccAddress()
		addr2     = sample.AccAddress()
		addr3     = sample.AccAddress()
		addr4     = sample.AccAddress()
	)
	tests := []struct {
		name     string
		genState *GenesisState
		err      error
	}{
		{
			name:     "default is valid",
			genState: DefaultGenesis(),
		}, {
			name: "valid custom genesis",
			genState: &GenesisState{
				ConsensusKeyNonceList: []*ConsensusKeyNonce{
					{ConsAddress: consAddr1, Nonce: 1},
					{ConsAddress: consAddr2, Nonce: 2},
					{ConsAddress: consAddr3, Nonce: 3},
					{ConsAddress: consAddr4, Nonce: 4},
				},
				ValidatorByAddressList: []*ValidatorByAddress{
					{Address: addr1, ConsensusAddress: consAddr1},
					{Address: addr2, ConsensusAddress: consAddr2},
					{Address: addr3, ConsensusAddress: consAddr3},
					{Address: addr4, ConsensusAddress: consAddr4},
				},
				ValidatorByConsAddressList: []*ValidatorByConsAddress{
					{ConsAddress: consAddr1, Address: addr1},
					{ConsAddress: consAddr2, Address: addr2},
					{ConsAddress: consAddr3, Address: addr3},
					{ConsAddress: consAddr4, Address: addr4},
				},
			},
		}, {
			name: "duplicated consensus key",
			genState: &GenesisState{
				ConsensusKeyNonceList: []*ConsensusKeyNonce{
					{ConsAddress: consAddr1, Nonce: 1},
					{ConsAddress: consAddr1, Nonce: 1},
				},
			},
			err: fmt.Errorf("duplicated index for consensusKeyNonce: %s", consAddr1),
		}, {
			name: "duplicated validator by address",
			genState: &GenesisState{
				ValidatorByAddressList: []*ValidatorByAddress{
					{Address: addr1, ConsensusAddress: consAddr1},
					{Address: addr1, ConsensusAddress: consAddr1},
				},
			},
			err: fmt.Errorf("duplicated index for validatorByAddress: %s", addr1),
		}, {
			name: "duplicated validator by address",
			genState: &GenesisState{
				ConsensusKeyNonceList: []*ConsensusKeyNonce{
					{ConsAddress: consAddr1, Nonce: 1},
					{ConsAddress: consAddr2, Nonce: 2},
				},
				ValidatorByAddressList: []*ValidatorByAddress{
					{Address: addr1, ConsensusAddress: consAddr1},
					{Address: addr1, ConsensusAddress: consAddr2},
				},
			},
			err: fmt.Errorf("consesus key not found for address: %s", consAddr1),
		}, {
			name: "valid custom genesis",
			genState: &GenesisState{
				ConsensusKeyNonceList: []*ConsensusKeyNonce{
					{ConsAddress: consAddr1, Nonce: 1},
					{ConsAddress: consAddr2, Nonce: 2},
				},
				ValidatorByAddressList: []*ValidatorByAddress{
					{Address: addr1, ConsensusAddress: consAddr1},
					{Address: addr2, ConsensusAddress: consAddr2},
				},
				ValidatorByConsAddressList: []*ValidatorByConsAddress{
					{ConsAddress: consAddr1, Address: addr1},
					{ConsAddress: consAddr1, Address: addr1},
				},
			},
			err: fmt.Errorf("duplicated index for validatorByConsAddress: %s", consAddr1),
		}, {
			name: "valid custom genesis",
			genState: &GenesisState{
				ConsensusKeyNonceList: []*ConsensusKeyNonce{
					{ConsAddress: consAddr1, Nonce: 1},
					{ConsAddress: consAddr2, Nonce: 2},
				},
				ValidatorByAddressList: []*ValidatorByAddress{
					{Address: addr1, ConsensusAddress: consAddr1},
					{Address: addr2, ConsensusAddress: consAddr2},
				},
				ValidatorByConsAddressList: []*ValidatorByConsAddress{
					{ConsAddress: consAddr1, Address: addr1},
					{ConsAddress: consAddr2, Address: addr2},
					{ConsAddress: consAddr3, Address: addr3},
				},
			},
			err: fmt.Errorf("validator not found for address: %s", consAddr3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.genState.validateValidator()
				if tt.err != nil {
					require.Error(t, err)
					assert.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.NoError(t, err)
			})
		})
	}
}

func TestGenesisStateValidateCoordinator(t *testing.T) {
	var (
		addr1 = sample.AccAddress()
		addr2 = sample.AccAddress()
		addr3 = sample.AccAddress()
		addr4 = sample.AccAddress()
	)
	tests := []struct {
		name     string
		genState *GenesisState
		err      error
	}{
		{
			name:     "default is valid",
			genState: DefaultGenesis(),
		}, {
			name: "valid custom genesis",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 10, Address: addr1},
					{CoordinatorId: 11, Address: addr2},
					{CoordinatorId: 12, Address: addr3},
					{CoordinatorId: 13, Address: addr4},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 10, Address: addr1},
					{CoordinatorId: 11, Address: addr2},
					{CoordinatorId: 12, Address: addr3},
					{CoordinatorId: 13, Address: addr4},
				},
			},
		}, {
			name: "duplicated coordinator",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 10, Address: addr1},
					{CoordinatorId: 10, Address: addr1},
				},
			},
			err: fmt.Errorf("duplicated index for coordinatorByAddress: %s", addr1),
		}, {
			name: "profile associated with chain",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 10, Address: addr1},
					{CoordinatorId: 10, Address: addr2},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 10, Address: addr1},
					{CoordinatorId: 10, Address: addr2},
				},
			},
			err: fmt.Errorf("duplicated id for coordinator: 10"),
		}, {
			name: "profile not associated with chain",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 10, Address: addr1},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 10, Address: addr1},
					{CoordinatorId: 11, Address: addr2},
				},
			},
			err: fmt.Errorf("coordinator address not found for CoordinatorByAddress: %s", addr2),
		}, {
			name: "profile not associated with chain",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 10, Address: addr1},
					{CoordinatorId: 11, Address: addr2},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 10, Address: addr1},
				},
			},
			err: fmt.Errorf("coordinator address not found for coordinatorID: 11"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.genState.validateCoordinator()
				if tt.err != nil {
					require.Error(t, err)
					assert.Equal(t, tt.err.Error(), err.Error())
					return
				}
				require.NoError(t, err)
			})
		})
	}
}
