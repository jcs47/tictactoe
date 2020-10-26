package tictactoe

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/lisbon/jcs/x/tictactoe/keeper"
	"github.com/lisbon/jcs/x/tictactoe/types"
)

func handleMsgMakeMove(ctx sdk.Context, k keeper.Keeper, msg types.MsgMakeMove) (*sdk.Result, error) {

	game, err := k.GetTTTGame(ctx, msg.GameId)

	if err != nil {

		return nil, err
	}

	err = game.MakeMove(msg.Creator, msg.X, msg.Y)

	if err != nil {

		return nil, err
	}

	logTTTGame(game, ctx.Logger())

	k.SetTTTGame(ctx, game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func logTTTGame(game types.TTTGame, log log.Logger) {

	status := game.GameFinished()

	log.Info(fmt.Sprintf(
		"\nPlayers:\n\nX: %s\nO: %s\n\nStatus: %s\n\n",
		game.Players[types.X], game.Players[types.O], printStatus(status)))

	if !status {

		log.Info(fmt.Sprintf(
			"\nNext move: %s\n\n", printRole(game.Next)))
	}

	log.Info("\n" + printTTTState(game) + "\n")

}

func printTTTState(game types.TTTGame) string {

	return fmt.Sprintf("\n"+
		" %s | %s | %s \n"+
		"-----------\n"+
		" %s | %s | %s \n"+
		"-----------\n"+
		" %s | %s | %s \n",
		printRole(game.State[0][2]), printRole(game.State[1][2]), printRole(game.State[2][2]),
		printRole(game.State[0][1]), printRole(game.State[1][1]), printRole(game.State[2][1]),
		printRole(game.State[0][0]), printRole(game.State[1][0]), printRole(game.State[2][0]))
}

func printRole(role types.Role) string {

	if role == types.Empty {
		return " "
	}

	if role == types.O {
		return "O"
	}

	return "X"
}

func printStatus(status bool) string {

	if status {
		return "Finished"
	}

	return "Ongoing"

}
