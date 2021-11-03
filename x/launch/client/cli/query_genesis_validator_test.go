package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/launch/client/cli"
	"github.com/tendermint/spn/x/launch/types"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func networkWithGenesisValidatorObjects(t *testing.T, n int) (*network.Network, []types.GenesisValidator) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	for i := 0; i < n; i++ {
		state.GenesisValidatorList = append(
			state.GenesisValidatorList,
			sample.GenesisValidator(uint64(0), strconv.Itoa(i)),
		)
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.GenesisValidatorList
}

func TestShowGenesisValidator(t *testing.T) {
	net, objs := networkWithGenesisValidatorObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc      string
		idChainID string
		idAddress string

		args []string
		err  error
		obj  types.GenesisValidator
	}{
		{
			desc:      "found",
			idChainID: strconv.Itoa(int(objs[0].LaunchID)),
			idAddress: objs[0].Address,

			args: common,
			obj:  objs[0],
		},
		{
			desc:      "not found",
			idChainID: strconv.Itoa(100000),
			idAddress: strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				tc.idChainID,
				tc.idAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowGenesisValidator(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetGenesisValidatorResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.GenesisValidator)
				require.Equal(t, tc.obj, resp.GenesisValidator)
			}
		})
	}
}

func TestListGenesisValidator(t *testing.T) {
	net, objs := networkWithGenesisValidatorObjects(t, 5)

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
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisValidator(), args)
			require.NoError(t, err)
			var resp types.QueryAllGenesisValidatorResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t, objs, resp.GenesisValidator)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(chainID, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisValidator(), args)
			require.NoError(t, err)
			var resp types.QueryAllGenesisValidatorResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.GenesisValidator), step)
			require.Subset(t, objs, resp.GenesisValidator)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(chainID, nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListGenesisValidator(), args)
		require.NoError(t, err)
		var resp types.QueryAllGenesisValidatorResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t, objs, resp.GenesisValidator)
	})
}
