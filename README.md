# Raylib Hot Reload in Go

A proof of concept of using hot reload with the raylib library in Go.

## Overview

This project demonstrates a hot-reloading mechanism for game development in Go, leveraging [Raylib](https://github.com/gen2brain/raylib-go) for graphics and [Yaegi](https://github.com/traefik/yaegi) for dynamic code interpretation. This setup allows for rapid iteration by applying code changes without recompiling or restarting the application.

## Prerequisites

Before running this project, ensure you have the following installed:

- **Go:** Make sure you have Go installed and configured correctly. This includes having a working Go environment with either `GOPATH` set up or Go modules enabled. (otherwise, `yaegi extract` will fail)
- **Yaegi:** Install Yaegi using the following command:

    ```bash
    go install github.com/traefik/yaegi/cmd/yaegi@latest
    ```

## How to Extract Raylib Symbols

To enable Yaegi to interact with Raylib, you need to extract the necessary symbols. Follow these steps:

1. Navigate to the `symbols` directory:

    ```bash
    cd symbols
    ```

2. Run the Yaegi extraction command:

    ```bash
    yaegi extract github.com/gen2brain/raylib-go/raylib
    ```

## How to Run the Project

1. Clone the repository.
2. Navigate to the project directory.
3. Run the main application:

    ```bash
    go run .
    ```

4. Modify the `src/game.go` file while the application is running. Changes will be applied dynamically.

## Project Structure

- `main.go`: The main application file that sets up the Raylib window, the Yaegi interpreter, and the file watcher.
- `src/game.go`: Contains the game logic that is dynamically reloaded.
- `symbols/`: Contains the extracted Raylib symbols.

## Inspiration

This project was inspired by [Gandalf-Le-Dev/ebitengine-yaegi-hotreload](https://github.com/Gandalf-Le-Dev/ebitengine-yaegi-hotreload), which demonstrates a similar hot-reloading approach for the Ebitengine game library.
