package participation

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	fundraisingkeeper "github.com/tendermint/fundraising/x/fundraising/keeper"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/tendermint/spn/x/participation/client/cli"
	"github.com/tendermint/spn/x/participation/keeper"
	"github.com/tendermint/spn/x/participation/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface for the participation module.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the participation module's name.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers the participation module's interface types
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns the participation module's default genesis state.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis performs genesis state validation for the participation module.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterRESTRoutes registers the participation module's REST service handlers.
func (AppModuleBasic) RegisterRESTRoutes(clientCtx client.Context, rtr *mux.Router) {
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the participation module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)) // nolint
}

// GetTxCmd returns the participation module's root tx command.
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the participation module's root query command.
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface for the participation module.
type AppModule struct {
	AppModuleBasic
	keeper            keeper.Keeper
	accountKeeper     authkeeper.AccountKeeper
	bankKeeper        bankkeeper.Keeper
	fundraisingKeeper fundraisingkeeper.Keeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	bankKeeper bankkeeper.Keeper,
	fundraisingkeeper fundraisingkeeper.Keeper,
) AppModule {
	return AppModule{
		AppModuleBasic:    NewAppModuleBasic(cdc),
		keeper:            keeper,
		accountKeeper:     accountKeeper,
		bankKeeper:        bankKeeper,
		fundraisingKeeper: fundraisingkeeper,
	}
}

// Name returns the participation module's name.
func (am AppModule) Name() string {
	return am.AppModuleBasic.Name()
}

// Route returns the participation module's message routing key.
func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, NewHandler(am.keeper))
}

// QuerierRoute returns the participation module's query routing key.
func (AppModule) QuerierRoute() string { return types.QuerierRoute }

// LegacyQuerierHandler returns the participation module's Querier.
func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a GRPC query service to respond to the
// module-specific GRPC queries.
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the participation module's invariants.
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	keeper.RegisterInvariants(ir, am.keeper)
}

// InitGenesis performs the participation module's genesis initialization It returns
// no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the participation module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion implements ConsensusVersion.
func (AppModule) ConsensusVersion() uint64 { return 2 }

// BeginBlock executes all ABCI BeginBlock logic respective to the participation module.
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock executes all ABCI EndBlock logic respective to the participation module. It
// returns no validator updates.
func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}
