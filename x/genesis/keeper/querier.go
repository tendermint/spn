package keeper

import (
	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/genesis/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		case types.QueryListChains:
			return listChains(ctx, req, k, legacyQuerierCdc)
		case types.QueryShowChain:
			return showChain(ctx, req, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

func listChains(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryListChainsParams
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get the chains from the keeper
	chains := keeper.GetAllChains(ctx)

	// Read chainID
	var chainIDs []string
	for _, chain := range chains {
		chainIDs = append(chainIDs, chain.ChainID)
	}

	// Paginate
	start, end := client.Paginate(len(chainIDs), params.Page, params.Limit, 0)
	if start < 0 || end < 0 {
		chainIDs = []string{}
	} else {
		chainIDs = chainIDs[start:end]
	}

	res, err := codec.MarshalJSONIndent(legacyQuerierCdc, chainIDs)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func showChain(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryShowChainRequest
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get the chain
	chain, found := keeper.GetChain(ctx, params.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The chain doesn't exist")
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, chain)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
