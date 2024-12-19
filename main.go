package main

import (
	r "github.com/gen2brain/raylib-go/raylib"
	"hot-reload-poc/src/game"
)

const (
	screenWidth  = 800
	screenHeight = 450
	title        = "raylib hot reload (poc)"
	targerFPS    = 60
)

func main() {
	r.InitWindow(screenWidth, screenHeight, title)
	defer r.CloseWindow()

	r.SetTargetFPS(targerFPS)

	for !r.WindowShouldClose() {
		game.Draw()
	}
}
