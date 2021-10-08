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
	"github.com/tendermint/spn/x/campaign/client/cli"
	"github.com/tendermint/spn/x/campaign/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func networkWithMainnetVestingAccountObjects(t *testing.T, n int) (*network.Network, []types.MainnetVestingAccount) {
	t.Helper()
	cfg := network.DefaultConfig()
	state := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &state))

	campaignID := uint64(5)
	for i := 0; i < n; i++ {
		state.MainnetVestingAccountList = append(state.MainnetVestingAccountList, sample.MainnetVestingAccount(
			campaignID,
			sample.Address(),
		))
	}
	buf, err := cfg.Codec.MarshalJSON(&state)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	return network.New(t, cfg), state.MainnetVestingAccountList
}

func TestShowMainnetVestingAccount(t *testing.T) {
	net, objs := networkWithMainnetVestingAccountObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc         string
		idCampaignID uint64
		idAddress    string

		args []string
		err  error
		obj  types.MainnetVestingAccount
	}{
		{
			desc:         "found",
			idCampaignID: objs[0].CampaignID,
			idAddress:    objs[0].Address,

			args: common,
			obj:  objs[0],
		},
		{
			desc:         "not found",
			idCampaignID: 100000,
			idAddress:    strconv.Itoa(100000),

			args: common,
			err:  status.Error(codes.InvalidArgument, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idCampaignID)),
				tc.idAddress,
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowMainnetVestingAccount(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryGetMainnetVestingAccountResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
				require.NotNil(t, resp.MainnetVestingAccount)
				require.Equal(t, tc.obj, resp.MainnetVestingAccount)
			}
		})
	}
}

func TestListMainnetVestingAccount(t *testing.T) {
	net, objs := networkWithMainnetVestingAccountObjects(t, 5)

	campaignID := objs[0].CampaignID
	ctx := net.Validators[0].ClientCtx
	request := func(campaignID uint64, next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			strconv.FormatUint(campaignID, 10),
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
			args := request(campaignID, nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMainnetVestingAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllMainnetVestingAccountResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.MainnetVestingAccount), step)
			require.Subset(t, objs, resp.MainnetVestingAccount)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(campaignID, next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMainnetVestingAccount(), args)
			require.NoError(t, err)
			var resp types.QueryAllMainnetVestingAccountResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.MainnetVestingAccount), step)
			require.Subset(t, objs, resp.MainnetVestingAccount)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(campaignID, nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListMainnetVestingAccount(), args)
		require.NoError(t, err)
		var resp types.QueryAllMainnetVestingAccountResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t, objs, resp.MainnetVestingAccount)
	})
}
