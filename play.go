package main

import (
	"kevinhiggins/tb2d"
	"fmt"
	"time"
)

const ViewportWidth = 800
const ViewportHeight = 600

func play() {
	fillTriplets()
	tb2d.SetUpWindow(ViewportWidth, ViewportHeight, false)
	aiSeedRand(time.Now().UnixMilli())
	sideWidth := (ViewportWidth - ViewportHeight) / 2
	but1 := tb2d.NewButtonFromFile("reset.bmp", func() {println("Button 1 clicked")}, 0, 0)
	bounds := but1.GetBounds()
	// This bounds thing is misnamed, as here we're actually changing
	// the position... bounds sounds more passive
	but1.SetBounds(tb2d.Rect{sideWidth - bounds.W, ViewportHeight - bounds.H, bounds.W, bounds.H})
	game := TicTacToeGame{}
	//game.setPlayerOAi(true)
	game.setAiPlayer(O)
	tg1 := tb2d.NewTileGridFromFiles([]string{"tile.bmp","x.bmp","o.bmp"}, func(gridX, gridY int) {
		if game.isOngoing() && game.hasUnmarkedSquare(gridX, gridY) {
			game.makeMove(gridX, gridY)
		}
	}, 3, 3, 200, 0)


	game.tileGrid = tg1
	

	
	tb2d.Start(func() {
		if time.Since(lastTurnTimestamp).Milliseconds() > 750 && game.isAiTurn() {
			if game.isOngoing() { // really the UI etc. should check this earlier
				moveX, moveY := aiSelectedPlayableMove(game)
				fmt.Printf("Debug %v, %v\n", moveX, moveY)
				for _, t := range tripletsContaining(Square{moveX, moveY}) {
					fmt.Println(t)
				}
				game.makeMove(moveX, moveY)
			}
		}
	})
	//but2 := tb2d.NewGraphicFromFile("quit.bmp")
	//tile := tb2d.NewGraphicFromFile("tile.bmp")

}
