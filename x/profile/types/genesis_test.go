package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenesisStateValidate(t *testing.T) {
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
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 1, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
					{CoordinatorId: 2, Address: "spn1d6pd5nk08mu789q4msfpynsuha7yf4wcsvvspr"},
					{CoordinatorId: 3, Address: "spn1ktzsme3g0ag0236ngvkw62vy9tqrr3xysnhp3g"},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 1, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
					{CoordinatorId: 2, Address: "spn1d6pd5nk08mu789q4msfpynsuha7yf4wcsvvspr"},
					{CoordinatorId: 3, Address: "spn1ktzsme3g0ag0236ngvkw62vy9tqrr3xysnhp3g"},
				},
				CoordinatorCount: 4,
			},
		}, {
			name: "duplicated coordinator",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 1, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("duplicated index for coordinatorByAddress: spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"),
		}, {
			name: "duplicated coordinator id",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 0, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 0, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("duplicated id for coordinator: 0"),
		}, {
			name: "profile not associated with chain",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 1, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("coordinator address not found for CoordinatorByAddress: spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"),
		}, {
			name: "profile not associated with chain",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 1, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("coordinator address not found for coordinatorID: 1"),
		}, {
			name: "invalid coordinator id",
			genState: &GenesisState{
				CoordinatorByAddressList: []*CoordinatorByAddress{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 133, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
				},
				CoordinatorList: []*Coordinator{
					{CoordinatorId: 0, Address: "spn1c7gh3kejxm3pzl8fwe65665xncs24x5rl7a8sm"},
					{CoordinatorId: 133, Address: "spn12330zcy9yez37lzrkm6d7fedcu7hc279sgkh3c"},
				},
				CoordinatorCount: 2,
			},
			err: fmt.Errorf("coordinator id 133 should be lower or equal than the last id 2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				err := tt.genState.Validate()
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
