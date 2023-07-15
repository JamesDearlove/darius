package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const screenWidth = 480
const screenHeight = 480

func main() {
	rl.InitWindow(screenWidth, screenHeight, "Darius")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {

		// DRAW
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkPurple)

		rl.DrawText("Hello World!", 100, 200, 40, rl.White)

		rl.EndDrawing()
	}
}
