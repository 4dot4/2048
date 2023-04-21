package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Board [4][4]int
}

var MyGame Game

func main() {

	MyGame.Board[0][3] = 3
	MyGame.Board[0][2] = 2
	MyGame.Board[0][1] = 1
	rl.InitWindow(1000, 1000, "2048 game")
	for !rl.WindowShouldClose() {
		update()

		draw() //teste
	}
	rl.CloseWindow()
}