package cli_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/participation/client/cli"
	"github.com/tendermint/spn/x/participation/types"
)

func networkWithDelegations(t *testing.T) (*network.Network, uint64) {
	t.Helper()
	cfg := network.DefaultConfig()

	// validator amount delegated to self
	totalShares := cfg.BondedTokens.ToDec()

	allocationPrice := types.AllocationPrice{Bonded: sdk.NewInt(rand.Int63n(100))}
	participationState := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &participationState))
	participationState.Params.AllocationPrice = allocationPrice
	buf, err := cfg.Codec.MarshalJSON(&participationState)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf

	return network.New(t, cfg), uint64(totalShares.Quo(allocationPrice.Bonded.ToDec()).TruncateInt64())
}

func TestShowTotalAllocations(t *testing.T) {
	net, totalAlloc := networkWithDelegations(t)

	ctx := net.Validators[0].ClientCtx

	delAddr := net.Validators[0].Address.String()

	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc       string
		delAddress string
		args       []string
		err        error
		alloc      uint64
	}{
		{
			desc:       "found",
			delAddress: delAddr,
			args:       common,
			alloc:      totalAlloc,
		},
		{
			desc:       "not found",
			delAddress: strconv.Itoa(100000),
			args:       common,
			err:        status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.delAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowTotalAllocations(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetTotalAllocationsResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.Equal(t,
					nullify.Fill(&tc.alloc),
					nullify.Fill(&resp.TotalAllocations),
				)
			}
		})
	}
}
