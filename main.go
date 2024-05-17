package main

import (
	"entity-prototype/components"
	ecs "entity-prototype/entity_component_system"
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

	manager := ecs.Manager{}

	// Create two entities
	player := manager.AddEntity()
	wall := manager.AddEntity()

	// Add ColliderComponent to player and wall
	playerCollider := &components.ColliderComponent{X: 0, Y: 0, Width: 50, Height: 50}
	wallCollider := &components.ColliderComponent{X: 30, Y: 30, Width: 50, Height: 50}

	player.AddComponent(playerCollider)
	wall.AddComponent(wallCollider)

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

		// Update the screen
		manager.Update()

		// Draw your game objects here
		manager.Draw(renderer)

		// Check for collision
		if playerCollider.CheckCollision(wallCollider) {
			fmt.Println("Collision detected between player and wall")
		} else {
			fmt.Println("No collision detected")
		}

		renderer.Present()
	}
}
