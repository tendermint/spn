package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	connectiontypes "github.com/cosmos/ibc-go/modules/core/03-connection/types"
	"github.com/cosmos/ibc-go/modules/core/exported"
	launchtypes "github.com/tendermint/spn/x/launch/types"
)

type LaunchKeeper interface {
	GetChain(ctx sdk.Context, launchID uint64) (val launchtypes.Chain, found bool)
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
