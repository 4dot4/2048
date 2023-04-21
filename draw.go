package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.DrawText("hell", 40, 40, 20, rl.Black)
	rl.EndDrawing()
}