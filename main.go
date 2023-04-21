package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Board [4][4]int
}

var MyGame Game

func main() {

	MyGame.Board[0][3] = 3
	MyGame.Board[0][2] = 2
	rl.InitWindow(700, 700, "2048 game")
	for !rl.WindowShouldClose() {
		update()

		fmt.Println(MyGame)
		draw() //teste
	}
	rl.CloseWindow()
}