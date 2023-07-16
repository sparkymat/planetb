package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	display := rl.GetCurrentMonitor()

	width := int32(rl.GetMonitorWidth(display))
	height := int32(rl.GetMonitorHeight(display))

	rl.SetConfigFlags(rl.FlagVsyncHint)

	rl.InitWindow(width, height, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))

	camera := rl.NewCamera3D(
		rl.NewVector3(0.0, -20.0, -20.0),
		rl.NewVector3(0.0, 0.0, 0.0),
		rl.NewVector3(0.0, 1.0, 0.0),
		45.0,
		rl.CameraPerspective,
	)

	for !rl.WindowShouldClose() {
		camera.Position.X = 500.0 - float64(time.Now().UnixMilli()%1000)
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawCube(rl.NewVector3(0, 0, 0), 10.0, 10.0, 10.0, rl.Green)

		rl.EndMode3D()
		rl.DrawFPS(10, 10)

		rl.EndDrawing()
	}
}
