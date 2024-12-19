# Raylib Hot Reload in Go

This project demonstrates a proof of concept for implementing hot reload in Go using [Raylib](https://github.com/gen2brain/raylib-go) and [Yaegi](https://github.com/traefik/yaegi). It showcases how to dynamically reload code changes without restarting the application.

## How to extract raylib symbols

1. Install [Yaegi](https://github.com/traefik/yaegi)

	```bash
	go install github.com/traefik/yaegi/cmd/yaegi@latest
	```

2. Run the following command to extract the symbols of the `raylib-go/raylib` package.

	```bash
	cd symbols
	yaegi extract github.com/gen2brain/raylib-go/raylib
	```

## Inspired by

[Gandalf-Le-Dev/ebitengine-yaegi-hotreload](https://github.com/Gandalf-Le-Dev/ebitengine-yaegi-hotreload)
