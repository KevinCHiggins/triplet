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
	aiPlayer map[int]bool // zero value is two human players
	turnsTaken int
	tileGrid *tb2d.TileGrid // would be nicer to separate out this detail
	message *tb2d.Button
}

// zero turns and tile states, but leave the loaded graphics & AI config
func (game *TicTacToeGame) resetPlay() {
	fmt.Println("Reset")
	game.turnsTaken = 0
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			game.tileGrid.SetTileAt(x, y, 0)
		}
	}
	if game.message != nil {
		tb2d.DeleteButton(game.message)
		game.message = nil
	}
	game.tileGrid.SetDirtyFlag()
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


func (game *TicTacToeGame) getMarkingsInTriplet(t TripletPosition) (xAmount, oAmount int) {
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

func (game *TicTacToeGame) makeMove(x, y int) {
	game.tileGrid.SetTileAt(x, y, game.getCurrentPlayer())
	lastTurnTimestamp = time.Now()
	game.tileGrid.SetDirtyFlag()
	game.turnsTaken++
	if game.evaluateState() == XWon {
		game.setMessageFromFile("xwon.bmp")
	} else if game.evaluateState() == OWon {
		game.setMessageFromFile("owon.bmp")
	} else if game.evaluateState() == Draw {
		game.setMessageFromFile("draw.bmp")
	}
}

func (game TicTacToeGame) setMessageFromFile(filename string) {
	msg := tb2d.NewButtonFromFile(filename, func() {}, 0, 0)
	msg.CenterInRect(game.tileGrid.GetBounds())
	game.message = msg
}

func (game *TicTacToeGame) evaluateState() State {
	winnableTripletForXFound := false // needed if we get to the 9th turn
	winnableTripletForOFound := false
	for _, t := range triplets {
		xAmount, oAmount := game.getMarkingsInTriplet(t)
		if xAmount == 3 {
			return XWon
		} else if oAmount == 3 {
			return OWon
		}
		if xAmount == 0 {
			winnableTripletForOFound = true
		}
		if oAmount == 0 {
			winnableTripletForXFound = true
		}
	}
	if !winnableTripletForOFound && !winnableTripletForXFound {
		return Draw
	} else if game.turnsTaken == 8 && !winnableTripletForXFound {
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
