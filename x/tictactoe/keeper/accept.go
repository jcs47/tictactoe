package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/lisbon/jcs/x/tictactoe/types"
)

// CreateAccept creates a accept
func (k Keeper) CreateAccept(ctx sdk.Context, accept types.Accept) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.AcceptPrefix + accept.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(accept)
	store.Set(key, value)
}

// GetAccept returns the accept information
func (k Keeper) GetAccept(ctx sdk.Context, key string) (types.Accept, error) {
	store := ctx.KVStore(k.storeKey)
	var accept types.Accept
	byteKey := []byte(types.AcceptPrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &accept)
	if err != nil {
		return accept, err
	}
	return accept, nil
}

// SetAccept sets a accept
func (k Keeper) SetAccept(ctx sdk.Context, accept types.Accept) {
	acceptKey := accept.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(accept)
	key := []byte(types.AcceptPrefix + acceptKey)
	store.Set(key, bz)
}

// DeleteAccept deletes a accept
func (k Keeper) DeleteAccept(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.AcceptPrefix + key))
}

//
// Functions used by querier
//

// Get creator of the item
func (k Keeper) GetAcceptOwner(ctx sdk.Context, key string) sdk.AccAddress {
	accept, err := k.GetAccept(ctx, key)
	if err != nil {
		return nil
	}
	return accept.Creator
}

// Check if the key exists in the store
func (k Keeper) AcceptExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.AcceptPrefix + key))
}
