package main

import (
	"kevinhiggins/tb2d"
	"fmt"
)

const (
	Unmarked = 0
	MarkedX = 1
	MarkedO = 2
)

type State int

const (
	XToMove State = iota + 1
	OToMove
	XWon
	OWon
	Draw
)

type TicTacToeGame struct {
	//board [3][3]SquareState
	turnsTaken int
	tileGrid *tb2d.TileGrid
}

func (game *TicTacToeGame) claimSquare(x, y int) error {
	if x < 0 || x > 2 || y < 0 || y > 2 {
		return fmt.Errorf("x %v, y %v out of range - must be in row 0-2, column 0-2", x, y)
	}
	if game.tileGrid.GetTileAt(x, y) == Unmarked {

		if game.evaluateState() == XToMove {
			game.tileGrid.SetTileAt(x, y, MarkedX)
		} else if game.evaluateState() == OToMove { 
			game.tileGrid.SetTileAt(x, y, MarkedO)
		} else {
			return fmt.Errorf("Game not in playable state")
		}
		game.tileGrid.SetDirtyFlag(true)
		game.turnsTaken++
		if game.evaluateState() == XWon {
			fmt.Println("X won")
		} else if game.evaluateState() == OWon {
			fmt.Println("O won")
		}
	} else {
		return fmt.Errorf("Row %v, column %v already marked", x, y)
	}
	return nil
}

// I'll leave out detecting inevitable draws...
func (game *TicTacToeGame) evaluateState() State {
	tg := game.tileGrid.GetGrid()
	switch {
	case
		(tg[0] == MarkedX && tg[1] == MarkedX && tg[2] == MarkedX) ||
		(tg[3] == MarkedX && tg[4] == MarkedX && tg[5] == MarkedX) ||
		(tg[6] == MarkedX && tg[7] == MarkedX && tg[8] == MarkedX) ||
		(tg[0] == MarkedX && tg[3] == MarkedX && tg[6] == MarkedX) ||
		(tg[1] == MarkedX && tg[4] == MarkedX && tg[7] == MarkedX) ||
		(tg[2] == MarkedX && tg[5] == MarkedX && tg[8] == MarkedX) ||
		(tg[0] == MarkedX && tg[4] == MarkedX && tg[8] == MarkedX) ||
		(tg[2] == MarkedX && tg[4] == MarkedX && tg[6] == MarkedX):
			return XWon
	case
		(tg[0] == MarkedO && tg[1] == MarkedO && tg[2] == MarkedO) ||
		(tg[3] == MarkedO && tg[4] == MarkedO && tg[5] == MarkedO) ||
		(tg[6] == MarkedO && tg[7] == MarkedO && tg[8] == MarkedO) ||
		(tg[0] == MarkedO && tg[3] == MarkedO && tg[6] == MarkedO) ||
		(tg[1] == MarkedO && tg[4] == MarkedO && tg[7] == MarkedO) ||
		(tg[2] == MarkedO && tg[5] == MarkedO && tg[8] == MarkedO) ||
		(tg[0] == MarkedO && tg[4] == MarkedO && tg[8] == MarkedO) ||
		(tg[2] == MarkedO && tg[4] == MarkedO && tg[6] == MarkedO):
			return OWon
	}
	if game.turnsTaken == 9 {
		return Draw
	} else if game.turnsTaken % 2 == 0 {
		return XToMove
	} else {
		return OToMove
	}
}