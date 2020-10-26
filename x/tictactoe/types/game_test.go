package types

import (
	"math/rand"
	"strings"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
)

var chars = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	"0123456789")

func TestGame(t *testing.T) {

	user1 := sdk.AccAddress(randStr(32))
	user2 := sdk.AccAddress(randStr(32))
	user3 := sdk.AccAddress(randStr(32))

	game, err := NewTTTGame(randStr(16), user1, user1)

	assert.NotNil(t, err) // Same player

	game, err = NewTTTGame(randStr(16), user1, user2)

	assert.Nil(t, err)

	assert.Equal(t, game.Players[O], user1)

	err = game.MakeMove(user1, 0, 0)

	assert.NotNil(t, err) // user1 has role O and cannot make first move

	err = game.MakeMove(user3, 0, 0)

	assert.NotNil(t, err) // user3 his not part of the game

	err = game.MakeMove(user2, -1, 0)

	assert.NotNil(t, err) // invalid coordinates

	err = game.MakeMove(user2, 0, 0)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user1, 0, 0)

	assert.NotNil(t, err) // Position already filled

	err = game.MakeMove(user1, 1, 1)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user2, 0, 1)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user1, 0, 2)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user2, 2, 2)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user1, 1, 2)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user2, 1, 0)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user1, 2, 1)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user2, 2, 0)

	assert.Nil(t, err)
	assert.True(t, game.GameFinished()) //Player X wins

	game, err = NewTTTGame(randStr(16), user3, user1)

	assert.Nil(t, err)

	err = game.MakeMove(user1, 1, 1)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user3, 0, 0)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user1, 1, 0)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user3, 1, 2)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user1, 0, 2)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user3, 2, 0)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user1, 2, 1)

	assert.Nil(t, err)
	assert.False(t, game.GameFinished())

	err = game.MakeMove(user3, 0, 1)

	assert.Nil(t, err)
	assert.True(t, game.GameFinished()) // draw

	err = game.MakeMove(user1, 2, 2)

	assert.NotNil(t, err) // game finished, no more moves allowed
	assert.True(t, game.GameFinished())

}

func randStr(length int) string {
	rand.Seed(time.Now().UnixNano())

	var b strings.Builder
	for i := 0; i < length; i++ {

		pos := rand.Intn(len(chars))
		b.WriteRune(chars[pos])
	}
	return b.String()
}
