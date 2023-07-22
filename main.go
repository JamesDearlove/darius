package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"net/http"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
)

const screenDiameter = 480
const screenCenter = 240
const pi = false

const maxGestureStrings = 20

func percentToVector(percent float64, length int32) rl.Vector2 {

	// Calculate angle, then adjust 90 degrees for top of screen.
	angle := 2*math.Pi*percent - math.Pi/2

	return rl.Vector2{
		X: float32(math.Cos(angle))*float32(length) + screenCenter,
		Y: float32(math.Sin(angle))*float32(length) + screenCenter,
	}
}

func downloadImage(url string) (image.Image, error) {
	if url == "" {
		return nil, nil
	}

	r, err := http.Get(url)
	if err != nil {
		// HTTP GET error
		fmt.Printf("[IMG] Failed to fetch image: %s\n", err)
		return nil, err
	}

	img, _, err := image.Decode(r.Body)
	if err != nil {
		// Decoding error
		fmt.Printf("[IMG] Failed to decode image: %s\n", err)
		return nil, err
	}

	return img, nil
}

func main() {
	// Enable 4x MSAA
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	rl.InitWindow(screenDiameter, screenDiameter, "Darius")
	// TODO: This should drop down when not active, for CPU cycle savings.
	rl.SetTargetFPS(60)
	defer rl.CloseWindow()

	// Constants
	backgroundColour := rl.DarkPurple
	centerVec := rl.Vector2{X: 240, Y: 240}

	img, err := downloadImage("https://source.unsplash.com/480x480")

	if err != nil {
		return
	}

	rlImg := rl.NewImageFromImage(img)
	rlText := rl.LoadTextureFromImage(rlImg)
	rl.UnloadImage(rlImg)

	// Gestures

	touchPosition := rl.NewVector2(0, 0)
	touchArea := rl.NewRectangle(0, 0, screenDiameter, screenDiameter)

	gestureStrings := make([]string, 0)

	currentGesture := rl.GestureNone
	lastGesture := rl.GestureNone

	for !rl.WindowShouldClose() {

		lastGesture = currentGesture
		currentGesture = rl.GetGestureDetected()
		touchPosition = rl.GetTouchPosition(0)

		if rl.CheckCollisionPointRec(touchPosition, touchArea) && currentGesture != rl.GestureNone {
			if currentGesture != lastGesture {
				switch currentGesture {
				case rl.GestureTap:
					gestureStrings = append(gestureStrings, "GESTURE TAP")
				case rl.GestureDoubletap:
					gestureStrings = append(gestureStrings, "GESTURE DOUBLETAP")
				case rl.GestureHold:
					gestureStrings = append(gestureStrings, "GESTURE HOLD")
				case rl.GestureDrag:
					gestureStrings = append(gestureStrings, "GESTURE DRAG")
				case rl.GestureSwipeRight:
					gestureStrings = append(gestureStrings, "GESTURE SWIPE RIGHT")
				case rl.GestureSwipeLeft:
					gestureStrings = append(gestureStrings, "GESTURE SWIPE LEFT")
				case rl.GestureSwipeUp:
					gestureStrings = append(gestureStrings, "GESTURE SWIPE UP")
				case rl.GestureSwipeDown:
					gestureStrings = append(gestureStrings, "GESTURE SWIPE DOWN")
				case rl.GesturePinchIn:
					gestureStrings = append(gestureStrings, "GESTURE PINCH IN")
				case rl.GesturePinchOut:
					gestureStrings = append(gestureStrings, "GESTURE PINCH OUT")
				}

				if len(gestureStrings) >= maxGestureStrings {
					gestureStrings = make([]string, 0)
				}
			}
		}

		// fmt.Println(gestureStrings)

		currentTime := time.Now()

		hour := float64(currentTime.Hour() % 12)
		minute := float64(currentTime.Minute())
		second := float64(currentTime.Second())
		milli := float64(currentTime.Nanosecond() / 1e6)

		hourVector := percentToVector((hour*60+minute)/720, 150)
		minuteVector := percentToVector((minute+second/60)/60, 200)
		secondVector := percentToVector((second*1000+milli)/60000, 200)

		// DRAW
		rl.BeginDrawing()

		// Draw circle to simulate if not on Pi
		if pi {
			rl.ClearBackground(backgroundColour)
		} else {
			rl.ClearBackground(rl.Black)
			rl.DrawCircle(240, 240, 240, backgroundColour)
		}
		
		rl.DrawTexture(rlText, 0, 0, rl.White)

		rl.DrawLineEx(centerVec, hourVector, 12, rl.White)
		rl.DrawLineEx(centerVec, minuteVector, 8, rl.White)
		rl.DrawLineEx(centerVec, secondVector, 4, rl.White)

		rl.DrawCircle(240, 240, 10, rl.Black)

		rl.DrawFPS(200, 440)

		rl.EndDrawing()
	}
}
