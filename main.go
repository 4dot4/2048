package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	board [4][4]int
}

var myGame Game

func main() {
	rl.InitWindow(700, 700, "2048 game")
	for !rl.WindowShouldClose() {
		update(myGame)
		draw(myGame) //teste
	}
	rl.CloseWindow()
}