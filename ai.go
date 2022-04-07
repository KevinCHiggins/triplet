package main

import (
	"math/rand"
)

func aiSeedRand(seed int64) {
	rand.Seed(seed)
}

func aiSelectedPlayableMove(game TicTacToeGame) (int, int) {
	var bestSquaresToPlay []Square
	var bestValFound int
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if game.tileGrid.GetTileAt(x, y) == Unmarked {
				continue
			} else {
				s := Square{x, y}
				tripletsContainingSquare := tripletsContaining(s)
				for _, t := range tripletsContainingSquare {
					xAmount, yAmount := game.getMarkingsInTriplet(t)
					val := aiValueTripletMarkingsForPlayer(xAmount, yAmount, game.getCurrentPlayer())
					if bestSquaresToPlay == nil || val > bestValFound {
						bestSquaresToPlay = make([]Square, 1)
						bestSquaresToPlay[0] = s
					} else if val == bestValFound {
						// note you could be adding a zero-value square to a nil slice
						bestSquaresToPlay = append(bestSquaresToPlay, s)
					}
				}
			}
		}
	}
	// SHOULD BE RANDOMLY SELECTED, HOWEVER FOR NOW WE TAKE THE FIRST
	return bestSquaresToPlay[0].X, bestSquaresToPlay[0].Y
}

func aiValueTripletMarkingsForPlayer(xAmount, yAmount, player int) int {
	var p, e int // player, enemy
	if player == X {
		p = xAmount
		e = yAmount
	} else {
		p = yAmount
		e = xAmount
	}
	if e == 1 && p == 0 {
		return 2
	} else if e == 1 && p == 1 {
		return 0
	} else if e == 2 {
		return 4
	} else if e == 0 && p == 1 {
		return 2
	} else if e == 0 && p == 2 {
		return 4
	} else if e == 0 && p == 0 {
		return 1
	}
	return 0 // shouldn't happen as we have covered all possibilities apart
	// from three squares being marked which should never be queried
}

// function assumes that there is a playable move
func aiSelectedRandomPlayableMove(game TicTacToeGame) (int, int) {	

	for { // infinite loop as we're sure there is a free square
		rX := rand.Intn(3)
		rY := rand.Intn(3)
		if game.hasUnmarkedSquare(rX, rY) {
			return rX, rY
		}
	}
}