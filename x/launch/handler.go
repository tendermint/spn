package launch

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/spn/x/launch/keeper"
	"github.com/tendermint/spn/x/launch/types"
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
		case *types.MsgCreateChain:
			res, err = msgServer.CreateChain(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgEditChain:
			res, err = msgServer.EditChain(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgRequestAddAccount:
			res, err = msgServer.RequestAddAccount(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgRequestAddVestingAccount:
			res, err = msgServer.RequestAddVestingAccount(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgRequestRemoveAccount:
			res, err = msgServer.RequestRemoveAccount(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgRequestAddValidator:
			res, err = msgServer.RequestAddValidator(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgRequestRemoveValidator:
			res, err = msgServer.RequestRemoveValidator(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgSettleRequest:
			res, err = msgServer.SettleRequest(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgTriggerLaunch:
			res, err = msgServer.TriggerLaunch(sdk.WrapSDKContext(ctx), msg)
		case *types.MsgRevertLaunch:
			res, err = msgServer.RevertLaunch(sdk.WrapSDKContext(ctx), msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			err = sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}

		return sdk.WrapServiceResult(ctx, res, err)
	}
}
