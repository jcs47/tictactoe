package cli

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lisbon/jcs/x/tictactoe/types"

	"github.com/spf13/cobra"
)

func GetCmdListGame(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-game",
		Short: "list all created games",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListGame, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list created games \n%s\n", err.Error())
				return nil
			}
			var out []types.TTTListItem

			//
			cdc.MustUnmarshalJSON(res, &out)
			//fmt.Printf(string(res))

			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetGame(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-game [id]",
		Short: "Obtain game state by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			id := args[0]

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetGame, id), nil)
			if err != nil {
				fmt.Printf("could not resolve game %s \n%s\n", id, err.Error())

				return nil
			}

			return printTTTGame(res)

		},
	}
}

func printTTTGame(data []byte) error {

	var game types.TTTGame
	err := json.Unmarshal(data, &game)

	if err != nil {
		return err
	}

	status := game.GameFinished()

	fmt.Printf(
		"\nPlayers:\n\nX: %s\nO: %s\n\nStatus: %s\n\n",
		game.Players[types.X], game.Players[types.O], printStatus(status))

	if !status {

		fmt.Printf(
			"\nNext move: %s\n\n", printRole(game.Next))
	}

	fmt.Printf("\n" + printTTTState(game) + "\n")

	return printJSON(data)

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

func printJSON(data []byte) error {

	fmt.Printf("\nJSON:\n\n")

	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, data, "", "\t")
	if err != nil {

		return err
	}

	fmt.Printf(string(prettyJSON.Bytes()) + "\n\n")

	return nil
}
