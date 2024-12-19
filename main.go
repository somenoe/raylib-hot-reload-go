package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"

	r "github.com/gen2brain/raylib-go/raylib"

	"github.com/somenoe/raylib-hot-reload-go/symbols"
)

type Game struct {
	interpreter *interp.Interpreter
}

const (
	screenWidth  = 800
	screenHeight = 450
	title        = "raylib hot reload (poc)"
	targerFPS    = 60
)

var draw func()

func reloadScripts() {
	i := interp.New(interp.Options{
		GoPath: os.Getenv("GOPATH"),
	})
	i.Use(stdlib.Symbols)
	i.Use(symbols.Symbols)

	bytes, err := os.ReadFile("src/game.go")
	if err != nil {
		log.Fatal(err)
		return
	}

	src := string(bytes)

	_, err = i.Eval(src)
	if err != nil {
		log.Fatal(err)
		return
	}

	v, err := i.Eval("src.Draw")
	if err != nil {
		log.Fatal(err)
	}

	draw = v.Interface().(func())
}

func watchFileChanges() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add("src/game.go")
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				log.Println("exiting")
				return
			}
			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("src/game.go changed, reloading")
				reloadScripts()
			}
		}
	}
}

func main() {

	reloadScripts()

	go watchFileChanges()

	r.InitWindow(screenWidth, screenHeight, title)
	defer r.CloseWindow()

	r.SetTargetFPS(targerFPS)

	for !r.WindowShouldClose() {
		draw()
	}
}
