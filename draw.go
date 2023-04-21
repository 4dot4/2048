package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Slice struct {
	Pos    rl.Vector2
	Color  rl.Color
	Side   int
	Number int
}

var TableSlice = Slice{
	Color: rl.DarkBrown,
	Side:  100,
}

func draw() {
	rl.BeginDrawing()
	for y := range MyGame.Board {
		for x := range MyGame.Board[y] {
			rl.DrawRectangleLinesEx(rl.Rectangle{
				X:      float32(50 + (x * TableSlice.Side) + 100),
				Y:      float32(50 + (y * TableSlice.Side) + 100),
				Width:  float32(TableSlice.Side),
				Height: float32(TableSlice.Side),
			}, 10, TableSlice.Color)
		}
	}
	rl.BeginDrawing()
	rl.ClearBackground(rl.White)
	rl.EndDrawing()
}