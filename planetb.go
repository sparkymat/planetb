package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	display := rl.GetCurrentMonitor()
	width := int32(rl.GetMonitorWidth(display))
	height := int32(rl.GetMonitorHeight(display))

	rl.InitWindow(width, height, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}
}
