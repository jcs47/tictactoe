package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/lisbon/jcs/x/tictactoe/types"
	"github.com/spf13/cobra"
)

func GetCmdListInvitation(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-invitation",
		Short: "list all invitation",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListInvitation, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Invitation\n%s\n", err.Error())
				return nil
			}
			var out []types.Invitation
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
