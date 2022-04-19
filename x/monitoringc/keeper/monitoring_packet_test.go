package keeper_test

import (
	"testing"

	channeltypes "github.com/cosmos/ibc-go/v2/modules/core/04-channel/types"
	"github.com/stretchr/testify/require"

	spntypes "github.com/tendermint/spn/pkg/types"
	tc "github.com/tendermint/spn/testutil/constructor"
	testkeeper "github.com/tendermint/spn/testutil/keeper"
	"github.com/tendermint/spn/testutil/sample"
	"github.com/tendermint/spn/x/monitoringc/types"
	profiletypes "github.com/tendermint/spn/x/profile/types"
	rewardtypes "github.com/tendermint/spn/x/reward/types"
)

func Test_OnRecvMonitoringPacket(t *testing.T) {
	var (
		ctx, tk, _     = testkeeper.NewTestSetup(t)
		invalidChannel = "invalidchannel"
		validChannel   = "monitoringtest"
		chain          = sample.Chain(r, 0, 0)
		valFoo         = sample.Address(r)
		valBar         = sample.Address(r)
		valOpAddrFoo   = sample.Address(r)
		valOpAddrBar   = sample.Address(r)
		coins          = sample.Coins(r)
	)

	tk.MonitoringConsumerKeeper.SetLaunchIDFromChannelID(ctx, types.LaunchIDFromChannelID{
		ChannelID: invalidChannel,
		LaunchID:  10000,
	})
	chain.LaunchID = tk.LaunchKeeper.AppendChain(ctx, chain)
	tk.MonitoringConsumerKeeper.SetLaunchIDFromChannelID(ctx, types.LaunchIDFromChannelID{
		ChannelID: validChannel,
		LaunchID:  chain.LaunchID,
	})

	tk.RewardKeeper.SetRewardPool(ctx, rewardtypes.RewardPool{
		LaunchID:         chain.LaunchID,
		Provider:         sample.Address(r),
		InitialCoins:     coins,
		RemainingCoins:   coins,
		LastRewardHeight: 1,
		Closed:           false,
	})
	err := tk.BankKeeper.MintCoins(ctx, rewardtypes.ModuleName, coins)
	require.NoError(t, err)

	// set validator profiles
	tk.ProfileKeeper.SetValidator(ctx, profiletypes.Validator{
		Address:           valFoo,
		OperatorAddresses: []string{valOpAddrFoo},
	})
	tk.ProfileKeeper.SetValidatorByOperatorAddress(ctx, profiletypes.ValidatorByOperatorAddress{

		ValidatorAddress: valFoo,
		OperatorAddress:  valOpAddrFoo,
	})
	tk.ProfileKeeper.SetValidator(ctx, profiletypes.Validator{
		Address:           valBar,
		OperatorAddresses: []string{valOpAddrBar},
	})
	tk.ProfileKeeper.SetValidatorByOperatorAddress(ctx, profiletypes.ValidatorByOperatorAddress{
		ValidatorAddress: valBar,
		OperatorAddress:  valOpAddrBar,
	})

	tests := []struct {
		name   string
		packet channeltypes.Packet
		data   spntypes.MonitoringPacket
		valid  bool
	}{
		{
			name:   "invalid data",
			packet: channeltypes.Packet{},
			data: spntypes.MonitoringPacket{
				BlockHeight: 0,
				SignatureCounts: spntypes.SignatureCounts{
					BlockCount: 1,
				},
			},
			valid: false,
		},
		{
			name: "no launch ID associated to channel ID",
			packet: channeltypes.Packet{
				DestinationChannel: "invalid",
			},
			data: spntypes.MonitoringPacket{
				BlockHeight: 1,
				SignatureCounts: spntypes.SignatureCounts{
					BlockCount: 1,
				},
			},
			valid: false,
		},
		{
			name: "should fail distribute rewards",
			packet: channeltypes.Packet{
				DestinationChannel: invalidChannel,
			},
			data: spntypes.MonitoringPacket{
				BlockHeight: 1,
				SignatureCounts: spntypes.SignatureCounts{
					BlockCount: 1,
				},
			},
			valid: false,
		},
		{
			name: "should successfully distribute rewards",
			packet: channeltypes.Packet{
				DestinationChannel: validChannel,
			},
			data: spntypes.MonitoringPacket{
				BlockHeight: 10,
				SignatureCounts: tc.SignatureCounts(10,
					tc.SignatureCount(t, valOpAddrFoo, "0.5"),
					tc.SignatureCount(t, valOpAddrBar, "0.5"),
				),
			},
			valid: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tk.MonitoringConsumerKeeper.OnRecvMonitoringPacket(ctx, tt.packet, tt.data)
			if !tt.valid {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}
