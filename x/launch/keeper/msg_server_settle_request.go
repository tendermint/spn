package keeper

import (
	"context"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	spnerrors "github.com/tendermint/spn/pkg/errors"
	"github.com/tendermint/spn/x/launch/types"
)

func (k msgServer) SettleRequest(
	goCtx context.Context,
	msg *types.MsgSettleRequest,
) (*types.MsgSettleRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chain, found := k.GetChain(ctx, msg.ChainID)
	if !found {
		return nil, sdkerrors.Wrap(types.ErrChainNotFound, msg.ChainID)
	}

	if chain.LaunchTriggered {
		return nil, sdkerrors.Wrap(types.ErrTriggeredLaunch, msg.ChainID)
	}

	coordAddress := k.profileKeeper.GetCoordinatorAddressFromID(ctx, chain.CoordinatorID)
	if msg.Coordinator != coordAddress {
		return nil, sdkerrors.Wrap(types.ErrNoAddressPermission, msg.Coordinator)
	}

	// first check if the request exist
	request, found := k.GetRequest(ctx, msg.ChainID, msg.RequestID)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrRequestNotFound,
			"request %d for chain %s not found",
			msg.RequestID,
			msg.ChainID,
		)
	}

	// perform request action
	k.RemoveRequest(ctx, msg.ChainID, request.RequestID)
	if msg.Approve {
		err := applyRequest(ctx, k.Keeper, msg, request)
		if err != nil {
			return nil, err
		}
	}

	return &types.MsgSettleRequestResponse{}, nil
}

// applyRequest approve the request and perform
// the launch information changes
func applyRequest(
	ctx sdk.Context,
	k Keeper,
	msg *types.MsgSettleRequest,
	request types.Request,
) error {
	cdc := codectypes.NewInterfaceRegistry()
	types.RegisterInterfaces(cdc)

	var content types.RequestContent
	if err := cdc.UnpackAny(request.Content, &content); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidRequestContent, err.Error())
	}

	switch c := content.(type) {
	case *types.GenesisAccount:
		_, foundGenesis := k.GetGenesisAccount(ctx, msg.ChainID, c.Address)
		_, foundVested := k.GetVestedAccount(ctx, msg.ChainID, c.Address)
		if foundVested || foundGenesis {
			return spnerrors.Critical(fmt.Sprintf(
				"account %s for chain %s already exist",
				c.Address, msg.ChainID),
			)
		}
		k.SetGenesisAccount(ctx, *c)
	case *types.VestedAccount:
		_, foundGenesis := k.GetGenesisAccount(ctx, msg.ChainID, c.Address)
		_, foundVested := k.GetVestedAccount(ctx, msg.ChainID, c.Address)
		if foundVested || foundGenesis {
			return spnerrors.Critical(fmt.Sprintf(
				"account %s for chain %s already exist",
				c.Address, msg.ChainID),
			)
		}
		k.SetVestedAccount(ctx, *c)
	case *types.AccountRemoval:
		_, foundGenesis := k.GetGenesisAccount(ctx, msg.ChainID, c.Address)
		_, foundVested := k.GetVestedAccount(ctx, msg.ChainID, c.Address)
		if foundGenesis {
			k.RemoveGenesisAccount(ctx, msg.ChainID, c.Address)
		}
		if foundVested {
			k.RemoveVestedAccount(ctx, msg.ChainID, c.Address)
		}
		if !foundGenesis && !foundVested {
			return spnerrors.Critical(fmt.Sprintf(
				"account %s for chain %s not found",
				c.Address, msg.ChainID),
			)
		}
	case *types.GenesisValidator:
		if _, found := k.GetGenesisValidator(ctx, msg.ChainID, c.Address); found {
			return spnerrors.Critical(fmt.Sprintf(
				"genesis validator %s for chain %s already exist",
				c.Address, msg.ChainID),
			)
		}
		k.SetGenesisValidator(ctx, *c)
	case *types.ValidatorRemoval:
		if _, found := k.GetGenesisValidator(ctx, msg.ChainID, c.ValAddress); !found {
			return spnerrors.Critical(fmt.Sprintf(
				"genesis validator %s for chain %s not found",
				c.ValAddress, msg.ChainID),
			)
		}
		k.RemoveGenesisValidator(ctx, msg.ChainID, c.ValAddress)
	default:
		return spnerrors.Critical("unknown request content type")
	}
	return nil
}
