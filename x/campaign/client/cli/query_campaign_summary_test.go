package cli_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/tendermint/spn/testutil/network"
	"github.com/tendermint/spn/testutil/nullify"
	"github.com/tendermint/spn/x/campaign/client/cli"
	"github.com/tendermint/spn/x/campaign/types"
	launchtypes "github.com/tendermint/spn/x/launch/types"
)

func networkWithCampaignSummariesObjects(t *testing.T, n int) (*network.Network, []types.CampaignSummary) {
	t.Helper()
	cfg := network.DefaultConfig()
	campaignState := types.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[types.ModuleName], &campaignState))
	chainState := launchtypes.GenesisState{}
	require.NoError(t, cfg.Codec.UnmarshalJSON(cfg.GenesisState[launchtypes.ModuleName], &chainState))
	objs := make([]types.CampaignSummary, 0)

	for i := 0; i < n; i++ {
		campaign := types.Campaign{
			CampaignID:         uint64(i),
			TotalSupply:        sdk.NewCoins(),
			AllocatedShares:    types.Shares(sdk.NewCoins()),
			SpecialAllocations: types.EmptySpecialAllocations(),
		}
		campaignState.CampaignChainsList = append(campaignState.CampaignChainsList, types.CampaignChains{
			CampaignID: uint64(i),
			Chains:     []uint64{uint64(i)},
		})
		campaignState.CampaignList = append(campaignState.CampaignList, campaign)
		chainState.ChainList = append(chainState.ChainList, launchtypes.Chain{
			LaunchID:    uint64(i),
			HasCampaign: true,
			CampaignID:  uint64(i),
		})
		chainState.ChainCounter += 1

		objs = append(objs, types.CampaignSummary{
			Campaign:           campaign,
			HasMostRecentChain: true,
			MostRecentChain: types.MostRecentChain{
				LaunchID: uint64(i),
			},
			Rewards:         sdk.NewCoins(),
			PreviousRewards: sdk.NewCoins(),
		})
	}
	buf, err := cfg.Codec.MarshalJSON(&campaignState)
	require.NoError(t, err)
	cfg.GenesisState[types.ModuleName] = buf
	buf, err = cfg.Codec.MarshalJSON(&chainState)
	require.NoError(t, err)
	cfg.GenesisState[launchtypes.ModuleName] = buf
	return network.New(t, cfg), objs
}

func TestListCampaignSummary(t *testing.T) {
	net, objs := networkWithCampaignSummariesObjects(t, 1)

	ctx := net.Validators[0].ClientCtx
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
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
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaignSummary(), args)
			require.NoError(t, err)
			var resp types.QueryCampaignSummariesResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.CampaignSummaries), step)
			require.Subset(t, nullify.Fill(objs), nullify.Fill(resp.CampaignSummaries))
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaignSummary(), args)
			require.NoError(t, err)
			var resp types.QueryCampaignSummariesResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.CampaignSummaries), step)
			require.Subset(t, nullify.Fill(objs), nullify.Fill(resp.CampaignSummaries))
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListCampaignSummary(), args)
		require.NoError(t, err)
		var resp types.QueryCampaignSummariesResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t, nullify.Fill(objs), nullify.Fill(resp.CampaignSummaries))
	})
}

func TestShowCampaignSummary(t *testing.T) {
	net, objs := networkWithCampaignSummariesObjects(t, 2)

	ctx := net.Validators[0].ClientCtx
	common := []string{
		fmt.Sprintf("--%s=json", tmcli.OutputFlag),
	}
	for _, tc := range []struct {
		desc         string
		idCampaignID uint64

		args []string
		err  error
		obj  types.CampaignSummary
	}{
		{
			desc:         "found",
			idCampaignID: objs[0].Campaign.CampaignID,

			args: common,
			obj:  objs[0],
		},
		{
			desc:         "not found",
			idCampaignID: 100000,

			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			args := []string{
				strconv.Itoa(int(tc.idCampaignID)),
			}
			args = append(args, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowCampaignSummary(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
			} else {
				require.NoError(t, err)
				var resp types.QueryCampaignSummaryResponse
				require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))

			}
		})
	}
}
