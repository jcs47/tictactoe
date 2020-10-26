package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/lisbon/jcs/x/tictactoe/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createMoveRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	GameId      string       `json:"gameId"`
	Coordinates string       `json:"coordinates"`
}

func makeMoveHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createMoveRequest
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}
		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}
		creator, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedGameId := req.GameId

		parsedCoordinates := req.Coordinates

		x, err := strconv.Atoi(string(parsedCoordinates[0]))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		y, err := strconv.Atoi(string(parsedCoordinates[1]))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgMakeMove(
			creator,
			parsedGameId,
			x, y,
		)

		err = msg.ValidateBasic()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}

type setMoveRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	ID          string       `json:"id"`
	Creator     string       `json:"creator"`
	GameId      string       `json:"gameId"`
	Coordinates string       `json:"coordinates"`
}

type deleteMoveRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
	ID      string       `json:"id"`
}
