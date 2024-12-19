package main

import r "github.com/gen2brain/raylib-go/raylib"

func main() {
	r.InitWindow(800, 450, "raylib hot reload (poc)")
	defer r.CloseWindow()

	r.SetTargetFPS(60)

	for !r.WindowShouldClose() {
		r.BeginDrawing()

		r.ClearBackground(r.Black)
		r.DrawText("Congrats! You created your first window!", 190, 200, 20, r.Green)

		r.EndDrawing()
	}
}
