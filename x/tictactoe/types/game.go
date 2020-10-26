package types

import (
	"bytes"
	"fmt"
	"math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Role represents a player
type Role byte

const (
	Empty Role = iota
	O
	X
)

const (
	Size      = 3 // Size of the matrix
	firstMove = X // Player that can have first move
)

//TTTGame represents an instance of tic tac toe
type TTTGame struct {
	ID      string                  `json:"id" yaml:"id"`
	State   [][]Role                `json:"state" yaml:"state"`
	Players map[Role]sdk.AccAddress `json:"players" yaml:"players"`
	Next    Role                    `json:"next" yaml:"next"`
}

//TTTListItem is used to list created games
type TTTListItem struct {
	ID       string `json:"id" yaml:"id"`
	Finished bool   `json:"finished" yaml:"finished"`
}

//NewTTTGame creates a new game
func NewTTTGame(id string, o, x sdk.AccAddress) (TTTGame, error) {

	if o == nil || x == nil {

		return TTTGame{}, fmt.Errorf("O and X cannot be nil")
	}

	if id == "" {

		return TTTGame{}, fmt.Errorf("ID cannot be an empty string")

	}

	if bytes.Equal(o.Bytes(), x.Bytes()) {

		return TTTGame{}, fmt.Errorf("O and X cannot be the same")
	}

	players := make(map[Role]sdk.AccAddress)
	state := make([][]Role, Size)

	players[O] = o
	players[X] = x

	for i := 0; i < Size; i++ {

		state[i] = make([]Role, Size)

		for j := 0; j < Size; j++ {

			state[i][j] = Empty

		}
	}

	return TTTGame{
		ID:      id,
		Next:    firstMove,
		Players: players,
		State:   state,
	}, nil
}

//GetRole returns the account address associated to the given role or
//'Empty' if the address is not from one of the players
func (ttt *TTTGame) GetRole(player sdk.AccAddress) Role {

	if bytes.Equal(ttt.Players[X], player) {
		return X
	}

	if bytes.Equal(ttt.Players[O], player) {
		return O
	}

	return Empty
}

//MakeMove submits a player's move given its account address and coordinates
func (ttt *TTTGame) MakeMove(player sdk.AccAddress, x, y int) error {

	if ttt.GameFinished() {

		return fmt.Errorf("Game is finished, no more moves allowed")
	}

	role := ttt.GetRole(player)

	if role == Empty {

		return fmt.Errorf("Player %x is not part of game %s", player, ttt.ID)
	}

	if role != ttt.Next {

		return fmt.Errorf("Player %x does not have the next move", player)

	}

	if x < 0 || x >= Size || y < 0 || y >= Size {

		return fmt.Errorf("Invalid coordinates")
	}

	if ttt.State[x][y] != Empty {

		return fmt.Errorf("Position (%d,%d) already filled", x, y)
	}

	ttt.State[x][y] = role

	if role == X {

		ttt.Next = O

	} else {

		ttt.Next = X

	}

	return nil
}

//GameFinished indicates if a game is finished
func (ttt *TTTGame) GameFinished() bool {

	return ttt.GameWinner() != Empty || ttt.GameDraw()
}

//GameWinner returns the game winner or 'Empty' if the game
//is still either ongoing or finished in a draw
func (ttt *TTTGame) GameWinner() Role {

	winner := ttt.diagonalWin()

	if winner != Empty {
		return winner
	}

	winner = ttt.counterDiagonalWin()

	if winner != Empty {
		return winner
	}

	winner = ttt.rowWin()

	if winner != Empty {
		return winner
	}

	winner = ttt.collumnWin()

	if winner != Empty {
		return winner
	}

	return Empty

}

//GameDraw indicates if the game ended in a tie
func (ttt *TTTGame) GameDraw() bool {

	maxMoves := int(math.Pow(Size, 2))
	totalMoves := ttt.countMoves()
	x := -1
	y := -1

	// account for the case when there is only one move
	// left to do, but that still results in a draw
	if totalMoves+1 == maxMoves {

		x, y = ttt.firstEmpty()

		ttt.State[x][y] = ttt.Next

	}

	res := ttt.diagonalDraw() &&
		ttt.counterDiagonalDraw() &&
		ttt.rowDraw() && ttt.collumnDraw()

	// make sure the state is not altered by this function
	if x != -1 {
		ttt.State[x][y] = Empty
	}

	return res
}

// if there is a collumn win, return the role of the winner
func (ttt *TTTGame) collumnWin() Role {

	for collumn := 0; collumn < Size; collumn++ {

		win := ttt.State[collumn][0]

		if win == Empty {
			continue
		}

		for y := 1; y < Size; y++ {

			if ttt.State[collumn][y-1] != ttt.State[collumn][y] {

				win = Empty
				break
			}

		}

		if win == Empty {

			continue
		}

		return win
	}

	return Empty

}

// if there is a row win, return the role of the winner
func (ttt *TTTGame) rowWin() Role {

	for row := 0; row < Size; row++ {

		win := ttt.State[0][row]

		if win == Empty {
			continue
		}

		for x := 1; x < Size; x++ {

			if ttt.State[x-1][row] != ttt.State[x][row] {

				win = Empty
				break
			}

		}

		if win == Empty {

			continue
		}

		return win
	}

	return Empty

}

// if there is a diagonal win, return the role of the winner
func (ttt *TTTGame) diagonalWin() Role {

	if ttt.State[0][0] == Empty {
		return Empty
	}

	x := 1
	y := 1
	for x < Size {

		if ttt.State[x-1][y-1] != ttt.State[x][y] {

			return Empty
		}

		x++
		y++

	}

	return ttt.State[0][0]

}

// if there is a counter diagonal win, return the role of the winner
func (ttt *TTTGame) counterDiagonalWin() Role {

	if ttt.State[0][Size-1] == Empty {
		return Empty
	}

	x := 1
	y := Size - 2
	for x < Size {

		if ttt.State[x-1][y+1] != ttt.State[x][y] {

			return Empty
		}

		x++
		y--
	}

	return ttt.State[0][Size-1]
}

// counts the total number of moves made
func (ttt *TTTGame) countMoves() int {

	count := 0

	for i := 0; i < Size; i++ {

		for j := 0; j < Size; j++ {

			if ttt.State[i][j] != Empty {

				count++
			}
		}
	}

	return count
}

// returns the first empty position that it can find
func (ttt *TTTGame) firstEmpty() (int, int) {

	for i := 0; i < Size; i++ {

		for j := 0; j < Size; j++ {

			if ttt.State[i][j] == Empty {

				return i, j
			}
		}
	}

	return -1, -1
}

// returns true if it is no longer possible to get a win by filling a whole collumn
func (ttt *TTTGame) collumnDraw() bool {

	for collumn := 0; collumn < Size; collumn++ {

		o := false
		x := false

		for i := 0; i < Size; i++ {

			if ttt.State[collumn][i] == X {

				x = true
			}

			if ttt.State[collumn][i] == O {

				o = true
			}

			if x && o {

				break
			}

		}

		if !(x && o) {

			return false
		}

	}

	return true

}

// returns true if it is no longer possible to get a win by filling a whole row
func (ttt *TTTGame) rowDraw() bool {

	for row := 0; row < Size; row++ {

		o := false
		x := false

		for i := 0; i < Size; i++ {

			if ttt.State[i][row] == X {

				x = true
			}

			if ttt.State[i][row] == O {

				o = true
			}

			if x && o {

				break
			}
		}

		if !(x && o) {

			return false
		}
	}

	return true

}

// returns true if it is no longer possible to get a win by filling the diagonal
func (ttt *TTTGame) diagonalDraw() bool {

	o := false
	x := false

	for i := 0; i < Size; i++ {

		if ttt.State[i][i] == X {

			x = true
		}

		if ttt.State[i][i] == O {

			o = true
		}

		if x && o {
			return true
		}

	}

	return false

}

// returns true if it is no longer possible to get a win by filling the counter diagonal
func (ttt *TTTGame) counterDiagonalDraw() bool {

	o := false
	x := false

	i := 0
	j := Size - 1

	for i < Size {

		if ttt.State[i][j] == X {

			x = true
		}

		if ttt.State[i][j] == O {

			o = true
		}

		if x && o {
			return true
		}

		i++
		j--
	}

	return false

}
