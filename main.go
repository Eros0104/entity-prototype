package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	// Initialize SDL
	err := sdl.Init(uint32(sdl.INIT_EVERYTHING))
	if err != nil {
		fmt.Println("Failed to initialize SDL:", err)
		return
	}
	defer sdl.Quit()

	// Create a window
	window, err := sdl.CreateWindow("SDL Window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println("Failed to create window:", err)
		return
	}
	defer window.Destroy()

	// Create a renderer
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Failed to create renderer:", err)
		return
	}
	defer renderer.Destroy()

	// Start player
	player := NewCharacter(100, 100, 50, 50)
	player2 := NewCharacter2(200, 200, 50, 50)

	// Main loop
	running := true
	for running {
		// Handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		// Clear the renderer
		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		// Draw your game objects here
		player.Draw(renderer)
		player2.Draw(renderer)

		// Update the screen
		player.Update()
		player2.Update(player)
		renderer.Present()

	}
}
