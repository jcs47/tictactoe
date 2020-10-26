package tictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lisbon/jcs/x/tictactoe/keeper"
	"github.com/lisbon/jcs/x/tictactoe/types"
)

func handleMsgCreateInvitation(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateInvitation) (*sdk.Result, error) {
	var invitation = types.Invitation{
		Creator: msg.Creator,
		ID:      msg.ID,
		Message: msg.Message,
	}
	k.CreateInvitation(ctx, invitation)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
