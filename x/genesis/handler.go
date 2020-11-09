package genesis

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/genesis/keeper"
	"github.com/tendermint/spn/x/genesis/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types.MsgChainCreate:
			return handleMsgChainCreate(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgChainCreate(ctx sdk.Context, k keeper.Keeper, msg *types.MsgChainCreate) (*sdk.Result, error) {
	// Get the identity of the creator
	identity, err := k.IdentityKeeper.GetIdentifier(ctx, msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check the chain doesn't already exist
	_, found := k.GetChain(ctx, msg.ChainID)
	if found {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, fmt.Sprint("Chain with chain ID %v already exists", msg.ChainID))
	}

	// Create the chain
	chain, err := types.NewChain(
		msg.ChainID,
		identity,
		msg.SourceURL,
		msg.SourceHash,
		ctx.BlockTime(),
		msg.Genesis,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidChain, err.Error())
	}

	// Set the new chain in the store
	k.SetChain(ctx, *chain)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
