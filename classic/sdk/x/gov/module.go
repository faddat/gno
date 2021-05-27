package gov

import (
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/classic/abci/types"
	"github.com/tendermint/go-amino-x"

	"github.com/tendermint/classic/sdk/client/context"
	sdk "github.com/tendermint/classic/sdk/types"
	"github.com/tendermint/classic/sdk/types/module"
	"github.com/tendermint/classic/sdk/x/gov/client"
	"github.com/tendermint/classic/sdk/x/gov/client/cli"
	"github.com/tendermint/classic/sdk/x/gov/client/rest"
	"github.com/tendermint/classic/sdk/x/gov/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// app module basics object
type AppModuleBasic struct {
	proposalHandlers []client.ProposalHandler // proposal handlers which live in governance cli and rest
}

// NewAppModuleBasic creates a new AppModuleBasic object
func NewAppModuleBasic(proposalHandlers ...client.ProposalHandler) AppModuleBasic {
	return AppModuleBasic{
		proposalHandlers: proposalHandlers,
	}
}

var _ module.AppModuleBasic = AppModuleBasic{}

// module name
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// default genesis state
func (AppModuleBasic) DefaultGenesis() json.RawMessage {
	return amino.MustMarshalJSON(DefaultGenesisState())
}

// module validate genesis
func (AppModuleBasic) ValidateGenesis(bz json.RawMessage) error {
	var data GenesisState
	err := amino.UnmarshalJSON(bz, &data)
	if err != nil {
		return err
	}
	return ValidateGenesis(data)
}

// register rest routes
func (a AppModuleBasic) RegisterRESTRoutes(ctx context.CLIContext, rtr *mux.Router) {
	var proposalRESTHandlers []rest.ProposalRESTHandler
	for _, proposalHandler := range a.proposalHandlers {
		proposalRESTHandlers = append(proposalRESTHandlers, proposalHandler.RESTHandler(ctx))
	}

	rest.RegisterRoutes(ctx, rtr, proposalRESTHandlers)
}

// get the root tx command of this module
func (a AppModuleBasic) GetTxCmd() *cobra.Command {

	var proposalCLIHandlers []*cobra.Command
	for _, proposalHandler := range a.proposalHandlers {
		proposalCLIHandlers = append(proposalCLIHandlers, proposalHandler.CLIHandler())
	}

	return cli.GetTxCmd(StoreKey, proposalCLIHandlers)
}

// get the root query command of this module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(StoreKey)
}

//___________________________
// app module
type AppModule struct {
	AppModuleBasic
	keeper       Keeper
	supplyKeeper SupplyKeeper
}

// NewAppModule creates a new AppModule object
func NewAppModule(keeper Keeper, supplyKeeper SupplyKeeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,
		supplyKeeper:   supplyKeeper,
	}
}

// module name
func (AppModule) Name() string {
	return types.ModuleName
}

// register invariants
func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {
	RegisterInvariants(ir, am.keeper)
}

// module message route name
func (AppModule) Route() string {
	return RouterKey
}

// module handler
func (am AppModule) NewHandler() sdk.Handler {
	return NewHandler(am.keeper)
}

// module querier route name
func (AppModule) QuerierRoute() string {
	return QuerierRoute
}

// module querier
func (am AppModule) NewQuerierHandler() sdk.Querier {
	return NewQuerier(am.keeper)
}

// module init-genesis
func (am AppModule) InitGenesis(ctx sdk.Context, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState GenesisState
	amino.MustUnmarshalJSON(data, &genesisState)
	InitGenesis(ctx, am.keeper, am.supplyKeeper, genesisState)
	return []abci.ValidatorUpdate{}
}

// module export genesis
func (am AppModule) ExportGenesis(ctx sdk.Context) json.RawMessage {
	gs := ExportGenesis(ctx, am.keeper)
	return amino.MustMarshalJSON(gs)
}

// module begin-block
func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// module end-block
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	EndBlocker(ctx, am.keeper)
	return []abci.ValidatorUpdate{}
}
