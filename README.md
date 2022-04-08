### Triplet

An implementation of Tic Tac Toe written in Go, built on my Turn-Based 2D ("TB2D") game engine, which in turn uses SDL2.

The AI assigns a value to each unmarked square by summing weights for the state of each of the possible threes-in-a-row it could be part of, and chooses randomly from the squares with the highest value, so as to play the best possible game.

I have built this project on Linux (Xubuntu) only, so far.

Kevin Higgins, 08/04/22

![Screenshot of game showing Tic Tac Toe board with some Xs and Os](https://github.com/KevinCHiggins/triplet/raw/main/screenshot.png)
