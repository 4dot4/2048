package main

import rl "github.com/gen2brain/raylib-go/raylib"

func update() {
	switch key := rl.GetKeyPressed(); key {
	case rl.KeyA:
		for y := 0; y < len(MyGame.Board); y++ {
			for x := len(MyGame.Board[y]) - 1; x > -1; x-- {
				if x-1 >= 0 {
					if MyGame.Board[y][x-1] == 0 && MyGame.Board[y][x] != 0 {
						MyGame.Board[y][x-1] = MyGame.Board[y][x]
						MyGame.Board[y][x] = 0
						if x+2 < 4 {
							x++
						}

					}
				}

			}
		}
	case rl.KeyD:

	case rl.KeyW:

	case rl.KeyS:

	default:

	}
}