package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	tmtypes "github.com/tendermint/tendermint/types"

	spntypes "github.com/tendermint/spn/pkg/types"
	launchtypes "github.com/tendermint/spn/x/launch/types"
)

type LaunchKeeper interface {
	GetChain(ctx sdk.Context, launchID uint64) (launchtypes.Chain, bool)
	EnableMonitoringConnection(ctx sdk.Context, launchID uint64) error
	CheckValidatorSet(
		ctx sdk.Context,
		launchID uint64,
		chainID string,
		validatorSet tmtypes.ValidatorSet,
	) error
}

type RewardKeeper interface {
	DistributeRewards(
		ctx sdk.Context,
		launchID uint64,
		signatureCounts spntypes.SignatureCounts,
		lastBlockHeight int64,
		closeRewardPool bool,
	) error
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
}

// ClientKeeper is imported to add the ability to create IBC Client from the module
type ClientKeeper interface {
	CreateClient(
		ctx sdk.Context, clientState exported.ClientState, consensusState exported.ConsensusState,
	) (string, error)
}

// ConnectionKeeper is imported to check client ID during IBC handshake
type ConnectionKeeper interface {
	GetConnection(ctx sdk.Context, connectionID string) (connectiontypes.ConnectionEnd, bool)
}

// ChannelKeeper defines the expected IBC channel keeper
type ChannelKeeper interface {
	GetChannel(ctx sdk.Context, srcPort, srcChan string) (channeltypes.Channel, bool)
	GetNextSequenceSend(ctx sdk.Context, portID, channelID string) (uint64, bool)
	SendPacket(
		ctx sdk.Context,
		channelCap *capabilitytypes.Capability,
		sourcePort string,
		sourceChannel string,
		timeoutHeight clienttypes.Height,
		timeoutTimestamp uint64,
		data []byte,
	) (uint64, error)
	WriteAcknowledgement(ctx sdk.Context, chanCap *capabilitytypes.Capability, packet exported.PacketI, acknowledgement exported.Acknowledgement) error
	ChanCloseInit(ctx sdk.Context, portID, channelID string, chanCap *capabilitytypes.Capability) error
}

// PortKeeper defines the expected IBC port keeper
type PortKeeper interface {
	BindPort(ctx sdk.Context, portID string) *capabilitytypes.Capability
}
