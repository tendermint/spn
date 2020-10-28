package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/identity/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		case types.QueryUsername:
			return getUsername(ctx, req, k, legacyQuerierCdc)
		case types.QueryUsernameFromAddress:
			return getUsernameFromAddress(ctx, req, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

func getUsername(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryUsernameRequest

	// Decode the request params
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	username, err := keeper.GetUsername(ctx, params.Identifier)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, username)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func getUsernameFromAddress(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryUsernameFromAddressRequest

	// Decode the request params
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	addr, err := sdk.AccAddressFromBech32(params.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	username, err := keeper.GetUsernameFromAddress(ctx, addr)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, username)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}
