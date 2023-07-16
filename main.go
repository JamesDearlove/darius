package main

import (
	"math"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

const screenDiameter = 480
const screenCenter = 240
const pi = false

func percentToVector(percent float64, length int32) rl.Vector2 {

	// Calculate angle, then adjust 90 degrees for top of screen.
	angle := 2 * math.Pi * percent - math.Pi / 2

	return rl.Vector2{
		X: float32(math.Cos(angle)) * float32(length) + screenCenter, 
		Y: float32(math.Sin(angle)) * float32(length) + screenCenter,
	}
}

func main() {
	rl.InitWindow(screenDiameter, screenDiameter, "Darius")
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	// Constants
	backgroundColour := rl.DarkPurple

	centerVec := rl.Vector2{X: 240, Y: 240}

	for !rl.WindowShouldClose() {

		currentTime := time.Now()

		hour := float64(currentTime.Hour() % 12)
		minute := float64(currentTime.Minute())
		second := float64(currentTime.Second())
		milli := float64(currentTime.Nanosecond() / 1e6)

		hourVector := percentToVector((hour * 60 + minute) / 720, 150)
		minuteVector := percentToVector((minute + second / 60) / 60, 200)
		secondVector := percentToVector((second * 1000 + milli) / 60000, 200)

		// DRAW
		rl.BeginDrawing()

		// Draw circle to simulate if not on Pi
		if pi {
			rl.ClearBackground(backgroundColour)
		} else {
			rl.ClearBackground(rl.Black)
			rl.DrawCircle(240, 240, 240, backgroundColour)
		}

		rl.DrawLineEx(centerVec, hourVector, 12, rl.White)
		rl.DrawLineEx(centerVec, minuteVector, 8, rl.White)
		rl.DrawLineEx(centerVec, secondVector, 4, rl.White)

		rl.DrawCircle(240, 240, 10, rl.White)

		rl.DrawFPS(200, 440)

		rl.EndDrawing()
	}
}
