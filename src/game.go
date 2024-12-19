package src

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

func Draw() {
	r.BeginDrawing()

	r.ClearBackground(r.Black)
	r.DrawText("Congratulations! Your code is running now.", 190, 200, 20, r.Green)

	r.EndDrawing()
}
