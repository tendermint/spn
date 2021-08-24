package profile

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/spn/x/profile/keeper"
	"github.com/tendermint/spn/x/profile/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	msgServer := keeper.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		var res proto.Message
		var err error
		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgUpdateValidatorDescription:
			res, err = msgServer.UpdateValidatorDescription(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgDeleteValidator:
			res, err = msgServer.DeleteValidator(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgCreateCoordinator:
			res, err = msgServer.CreateCoordinator(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgUpdateCoordinatorDescription:
			res, err = msgServer.UpdateCoordinatorDescription(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgUpdateCoordinatorAddress:
			res, err = msgServer.UpdateCoordinatorAddress(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgDeleteCoordinator:
			res, err = msgServer.DeleteCoordinator(sdk.WrapSDKContext(ctx), msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}

		return sdk.WrapServiceResult(ctx, res, err)
	}
}
