package cli

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/lisbon/jcs/x/tictactoe/types"
)

func GetCmdMakeMove(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "make-move [gameId] [coordinates]",
		Short: "Makes next move",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsGameId := string(args[0])
			argsCoordinates := string(args[1])

			x, _ := strconv.Atoi(string(argsCoordinates[0]))
			y, _ := strconv.Atoi(string(argsCoordinates[1]))

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgMakeMove(cliCtx.GetFromAddress(), string(argsGameId), x, y)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
