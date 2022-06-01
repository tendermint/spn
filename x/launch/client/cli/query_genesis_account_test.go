package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/client/cli"
	"github.com/tendermint/spn/x/launch/types"
)

func networkWithGenesisAccountObjects(t *testing.T, n int) (*network.Network, []types.GenesisAccount) {
	t.Helper()
	r := sample.Rand()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		state.GenesisAccountList = append(
			state.GenesisAccountList,
			sample.GenesisAccount(r, 0, strconv.Itoa(i)),
		)
	}
	state.ChainList = append(state.ChainList, sample.Chain(r, 0, sample.Uint64(r)))
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.GenesisAccountList
}

func TestShowGenesisAccount(t *testing.T) {
	net, objs := networkWithGenesisAccountObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc       string
		idLaunchID string
		idAddress  string

		args []string
		err  error
		obj  types.GenesisAccount
	}{
		{
			desc:       "found",
			idLaunchID: strconv.Itoa(int(objs[0].LaunchID)),
			idAddress:  objs[0].Address,

			args: common,
			obj:  objs[0],
		},
		{
			desc:       "not found",
			idLaunchID: strconv.Itoa(100000),
			idAddress:  strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idLaunchID,
				tc.idAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowGenesisAccount(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetGenesisAccountResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.GenesisAccount)
				require.Equal(t, tc.obj, resp.GenesisAccount)
			}
		})
	}
}

func TestListGenesisAccount(t *testing.T) {
	net, objs := networkWithGenesisAccountObjects(t, 5)

	chainID := objs[0].LaunchID
	ctx := net.Validators[0].ClientCtx
	request := func(chainID uint64, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			strconv.Itoa(int(chainID)),
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(chainID, nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllGenesisAccountResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.GenesisAccount), step)
			require.Subset(t, objs, resp.GenesisAccount)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(chainID, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllGenesisAccountResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.GenesisAccount), step)
			require.Subset(t, objs, resp.GenesisAccount)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(chainID, nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisAccount(), args)
		require.NoError(t, err)
		var resp types.QueryAllGenesisAccountResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t, objs, resp.GenesisAccount)
	})
}
