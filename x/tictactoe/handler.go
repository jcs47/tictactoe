package tictactoe

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/x/auth"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/lisbon/jcs/x/tictactoe/keeper"
	"github.com/lisbon/jcs/x/tictactoe/types"
)

// NewHandler creates an sdk.Handler for all the tictactoe type messages
func NewHandler(k keeper.Keeper, a auth.AccountKeeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {

		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case types.MsgMakeMove:
			return handleMsgMakeMove(ctx, k, msg)
		case types.MsgAcceptInvitation:
			return handleMsgAcceptInvitation(ctx, k, a, msg)
		case types.MsgCreateInvitation:
			return handleMsgCreateInvitation(ctx, k, msg)
		case types.MsgDeleteInvitation:
			return handleMsgDeleteInvitation(ctx, k, msg)

		// TODO: Define your msg cases
		//
		//Example:
		// case Msg<Action>:
		// 	return handleMsg<Action>(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handle<Action> does x
/*
func handleMsg<Action>(ctx sdk.Context, k keeper.Keeper, msg Msg<Action>) (*sdk.Result, error) {
	err := k.<Action>(ctx, msg.ValidatorAddr)
	if err != nil {
		return nil, err
	}

	// TODO: Define your msg events
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.ValidatorAddr.String()),
		),
	)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
*/
