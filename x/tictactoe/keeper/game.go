package keeper

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lisbon/jcs/x/tictactoe/types"
)

func (k Keeper) CreateTTTGame(ctx sdk.Context, game types.TTTGame) {

	store := ctx.KVStore(k.storeKey)
	key := []byte(types.TTTGamePrefix + game.ID)

	b, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
		return
	}

	value := k.cdc.MustMarshalBinaryLengthPrefixed(b)
	store.Set(key, value)
}

func (k Keeper) SetTTTGame(ctx sdk.Context, game types.TTTGame) {
	gameKey := game.ID
	store := ctx.KVStore(k.storeKey)

	b, err := json.Marshal(game)
	if err != nil {
		fmt.Println(err)
		return
	}

	bz := k.cdc.MustMarshalBinaryLengthPrefixed(b)
	key := []byte(types.TTTGamePrefix + gameKey)
	store.Set(key, bz)
}

func (k Keeper) GetTTTGame(ctx sdk.Context, key string) (types.TTTGame, error) {
	store := ctx.KVStore(k.storeKey)

	var b []byte
	var game types.TTTGame

	byteKey := []byte(types.TTTGamePrefix + key)

	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &b)
	if err != nil {
		return game, err
	}

	err = json.Unmarshal(b, &game)

	if err != nil {
		return game, err
	}

	return game, nil
}

func listTTTGame(ctx sdk.Context, k Keeper) ([]byte, error) {
	var gameIDsList []types.TTTListItem

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.TTTGamePrefix))
	for ; iterator.Valid(); iterator.Next() {

		var b []byte
		var game types.TTTGame

		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &b)

		err := json.Unmarshal(b, &game)

		if err != nil {
			return nil, err
		}

		item := types.TTTListItem{
			ID:       game.ID,
			Finished: game.GameFinished(),
		}

		gameIDsList = append(gameIDsList, item)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, gameIDsList)
	return res, nil
}

func getTTTGame(ctx sdk.Context, key string, k Keeper) ([]byte, error) {

	store := ctx.KVStore(k.storeKey)

	var b []byte

	byteKey := []byte(types.TTTGamePrefix + key)

	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &b)
	if err != nil {
		return []byte(err.Error()), nil
	}

	return b, nil

}
