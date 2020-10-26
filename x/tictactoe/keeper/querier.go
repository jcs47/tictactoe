package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/lisbon/jcs/x/tictactoe/types"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for tictactoe clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {

		case types.QueryListInvitation:
			return listInvitation(ctx, k)
		case types.QueryGetGame:
			return getTTTGame(ctx, path[1], k)
		case types.QueryListGame:
			return listTTTGame(ctx, k)

		case types.QueryParams:
			return queryParams(ctx, k)
			// TODO: Put the modules query routes

			// this line is used by starport scaffolding # 2

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown tictactoe query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// TODO: Add the modules query functions
// They will be similar to the above one: queryParams()
