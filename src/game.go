package src

import r "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 450
	title        = "raylib hot reload (poc)"
	targerFPS    = 60
)

func draw() {
	r.BeginDrawing()

	r.ClearBackground(r.Black)
	r.DrawText("Congratulations! Your code is running now.", 190, 200, 20, r.Green)

	r.EndDrawing()
}

func Run() {
	r.InitWindow(screenWidth, screenHeight, title)
	defer r.CloseWindow()

	r.SetTargetFPS(targerFPS)

	for !r.WindowShouldClose() {
		draw()
	}
}
