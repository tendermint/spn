package keeper

import (
	"fmt"
	// this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/spn/x/chat/types"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		case types.QueryShowChannel:
			return showChannel(ctx, req, k, legacyQuerierCdc)
		case types.QueryListChannels:
			return listChannel(ctx, req, k, legacyQuerierCdc)
		case types.QueryListMessages:
			return listMessages(ctx, req, k, legacyQuerierCdc)
		case types.QuerySearchMessages:
			return searchMessages(ctx, req, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}

func showChannel(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryShowChannelRequest

	// Decode the request params
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get the channel
	channel, found := keeper.GetChannel(ctx, params.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "The channel doesn't exist")
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, channel)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func listChannel(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryListChannelsRequest

	// Decode the request params
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	channels := keeper.GetAllChannels(ctx)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, channels)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func listMessages(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QueryListMessagesRequest

	// Decode the request params
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	messages, found := keeper.GetAllMessagesFromChannel(ctx, params.ChannelId)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("channel %v doesn't exist", params.ChannelId))
	}

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, messages)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func searchMessages(ctx sdk.Context, req abci.RequestQuery, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	var params types.QuerySearchMessagesRequest
	var messages []types.Message

	// Decode the request params
	err := legacyQuerierCdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	// Get the tag references
	tagReferences := keeper.GetTagReferencesFromChannel(ctx, params.Tag, params.ChannelId)
	if len(tagReferences) == 0 {
		bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, messages)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
		return bz, nil
	}

	// Get the messages from the tag references
	messages = keeper.GetMessagesByIDs(ctx, tagReferences)

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, messages)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
