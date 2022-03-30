package main

import (
	"kevinhiggins/tb2d"
	"fmt"

)

const ViewportWidth = 800
const ViewportHeight = 600

func play() {

	sideWidth := (ViewportWidth - ViewportHeight) / 2
	but1 := tb2d.NewButtonFromFile("reset.bmp", func() {println("Button 1 clicked")}, 0, 0)
	// This bounds thing is misnamed, as it's holding position... bounds sounds more passive
	bounds := but1.GetBounds()
	but1.SetBounds(tb2d.Rect{sideWidth - bounds.W, ViewportHeight - bounds.H, bounds.W, bounds.H})
	game := TicTacToeGame{}
	tg1 := tb2d.NewTileGridFromFiles([]string{"tile.bmp","x.bmp","o.bmp"}, func(gridX, gridY int) {
		err := game.claimSquare(gridX, gridY)
		if err != nil {
			fmt.Println("Error: %w", err)
		}
	}, 3, 3, 200, 0)


	game.tileGrid = tg1
	tb2d.SetUpWindow(ViewportWidth, ViewportHeight, false)

	
	tb2d.Start(func() {print("Y")})
	//but2 := tb2d.NewGraphicFromFile("quit.bmp")
	//tile := tb2d.NewGraphicFromFile("tile.bmp")

}
