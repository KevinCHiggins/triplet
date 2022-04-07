package main

import (
	"kevinhiggins/tb2d"
	"fmt"
	"time"
)

const (
	Unmarked = 0
	X = 1
	O = 2
)

type State int

const (
	XToMove State = iota + 1
	OToMove
	XWon
	OWon
	Draw
)

var lastTurnTimestamp time.Time

type TicTacToeGame struct {
	//board [3][3]SquareState
	aiPlayer map[int]bool // zero value is two human players
	turnsTaken int
	tileGrid *tb2d.TileGrid
}

func (game *TicTacToeGame) isAiTurn() bool {
	return game.aiPlayer[game.getCurrentPlayer()]
}

func (game *TicTacToeGame) setAiPlayer(player int) {
	if game.aiPlayer == nil {
		game.aiPlayer = make(map[int]bool)
	}
	game.aiPlayer[player] = true
}

func (game *TicTacToeGame) hasUnmarkedSquare(x, y int) bool {
	/*
	if x < 0 || x > 2 || y < 0 || y > 2 {
		return false // fmt.Errorf("x %v, y %v out of range - must be in row 0-2, column 0-2", x, y)
	}
	*/
	if game.tileGrid.GetTileAt(x, y) == Unmarked {
		return true

	} else {
		return false // fmt.Errorf("Row %v, column %v already marked", x, y)
	}

}


func (game *TicTacToeGame) getMarkingsInTriplet(t Triplet) (xAmount, oAmount int) {
	for _, square := range t {
		switch game.tileGrid.GetTileAt(square.X, square.Y) {
		case X:
			xAmount++
		case O:
			oAmount++
		}
	}
	return xAmount, oAmount
}


/*
func (game *TicTacToeGame) claimSquare(x, y int) error {
	if !game.hasUnmarkedSquare(x, y) {
		return fmt.Errorf("Game has no unmarked square at %v, %v", x, y)
	}

	switch game.evaluateState() {
	case XToMove:
		game.tileGrid.SetTileAt(x, y, X)
	case OToMove:
		game.tileGrid.SetTileAt(x, y, O)
	default:
		return fmt.Errorf("Game is over")
	}
	lastTurnTimestamp = time.Now()
	game.tileGrid.SetDirtyFlag(true)
	game.turnsTaken++
	if game.evaluateState() == XWon {
		fmt.Println("X won")
	} else if game.evaluateState() == OWon {
		fmt.Println("O won")
	}

	return nil
}
*/

func (game *TicTacToeGame) makeMove(x, y int) {
	game.tileGrid.SetTileAt(x, y, game.getCurrentPlayer())
	lastTurnTimestamp = time.Now()
	game.tileGrid.SetDirtyFlag(true)
	game.turnsTaken++
	if game.evaluateState() == XWon {
		fmt.Println("X won")
	} else if game.evaluateState() == OWon {
		fmt.Println("O won")
	}
}
// I'll leave out detecting upcoming draws...
func (game *TicTacToeGame) evaluateState() State {
	tg := game.tileGrid.GetGrid()
	// one could do this with functions, but the baked data is more concise
	switch {
	case
		(tg[0] == X && tg[1] == X && tg[2] == X) ||
		(tg[3] == X && tg[4] == X && tg[5] == X) ||
		(tg[6] == X && tg[7] == X && tg[8] == X) ||
		(tg[0] == X && tg[3] == X && tg[6] == X) ||
		(tg[1] == X && tg[4] == X && tg[7] == X) ||
		(tg[2] == X && tg[5] == X && tg[8] == X) ||
		(tg[0] == X && tg[4] == X && tg[8] == X) ||
		(tg[2] == X && tg[4] == X && tg[6] == X):
			return XWon
	case
		(tg[0] == O && tg[1] == O && tg[2] == O) ||
		(tg[3] == O && tg[4] == O && tg[5] == O) ||
		(tg[6] == O && tg[7] == O && tg[8] == O) ||
		(tg[0] == O && tg[3] == O && tg[6] == O) ||
		(tg[1] == O && tg[4] == O && tg[7] == O) ||
		(tg[2] == O && tg[5] == O && tg[8] == O) ||
		(tg[0] == O && tg[4] == O && tg[8] == O) ||
		(tg[2] == O && tg[4] == O && tg[6] == O):
			return OWon
	}
	if game.turnsTaken == 9 {
		return Draw
	} else {
		p := game.getCurrentPlayer()
		// HORRID double conversion
		if p == X {
			return XToMove
		} else {
			return OToMove
		}
	}
}

func (game *TicTacToeGame) isOngoing() bool {
	state := game.evaluateState()
	return state == XToMove || state == OToMove
}
func (game *TicTacToeGame) getCurrentPlayer() int {
	if game.turnsTaken % 2 == 0 {
		return X
	} else {
		return O
	}
}
