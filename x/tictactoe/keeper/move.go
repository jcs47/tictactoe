package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lisbon/jcs/x/tictactoe/types"
)

// CreateMove creates a move
func (k Keeper) CreateMove(ctx sdk.Context, move types.Move) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.MovePrefix + move.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(move)
	store.Set(key, value)
}

// GetMove returns the move information
func (k Keeper) GetMove(ctx sdk.Context, key string) (types.Move, error) {
	store := ctx.KVStore(k.storeKey)
	var move types.Move
	byteKey := []byte(types.MovePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &move)
	if err != nil {
		return move, err
	}
	return move, nil
}

// SetMove sets a move
func (k Keeper) SetMove(ctx sdk.Context, move types.Move) {
	moveKey := move.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(move)
	key := []byte(types.MovePrefix + moveKey)
	store.Set(key, bz)
}

// DeleteMove deletes a move
func (k Keeper) DeleteMove(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.MovePrefix + key))
}

//
// Functions used by querier
//

// Get creator of the item
func (k Keeper) GetMoveOwner(ctx sdk.Context, key string) sdk.AccAddress {
	move, err := k.GetMove(ctx, key)
	if err != nil {
		return nil
	}
	return move.Creator
}

// Check if the key exists in the store
func (k Keeper) MoveExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.MovePrefix + key))
}
