package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers tictactoe-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding # 1
	r.HandleFunc("/tictactoe/move", makeMoveHandler(cliCtx)).Methods("POST")

	r.HandleFunc("/tictactoe/accept", acceptInvitationHandler(cliCtx)).Methods("POST")

	r.HandleFunc("/tictactoe/invitation", createInvitationHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/tictactoe/invitation", listInvitationHandler(cliCtx, "tictactoe")).Methods("GET")
	r.HandleFunc("/tictactoe/invitation", deleteInvitationHandler(cliCtx)).Methods("DELETE")

	r.HandleFunc("/tictactoe/game", listTTTGameHandler(cliCtx, "tictactoe")).Methods("GET")
	r.HandleFunc("/tictactoe/game/{key}", getTTTGameHandler(cliCtx, "tictactoe")).Methods("GET")

	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
