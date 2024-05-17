package main

import (
	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
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
	follower := manager.AddEntity()
	wall := manager.AddEntity()

	// Add ColliderComponent to player and wall
	playerCollider := &components.ColliderComponent{X: 100, Y: 100, Width: 50, Height: 50}
	wallCollider := &components.ColliderComponent{X: 30, Y: 30, Width: 50, Height: 50}
	followerCollider := &components.ColliderComponent{X: 200, Y: 200, Width: 50, Height: 50}

	// Add input handler component to player
	playerInputHandler := &components.InputHandlerComponent{Speed: 1}

	// Add rect component to player and wall
	playerRect := &components.RectComponent{R: 255, G: 0, B: 0, A: 255}
	wallRect := &components.RectComponent{R: 0, G: 255, B: 0, A: 255}
	followerRect := &components.RectComponent{R: 0, G: 0, B: 255, A: 255}

	// Add follow component to follower
	followerFollow := &components.FollowComponent{Destination: playerCollider, Speed: 0.1}

	player.AddComponent(playerCollider)
	player.AddComponent(playerInputHandler)
	player.AddComponent(playerRect)

	wall.AddComponent(wallCollider)
	wall.AddComponent(wallRect)

	follower.AddComponent(followerCollider)
	follower.AddComponent(followerRect)
	follower.AddComponent(followerFollow)

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

		// Check for collision
		if playerCollider.CheckCollision(followerCollider) {
			fmt.Println("Collision detected between player and follower")
		} else {
			fmt.Println("No collision detected")
		}

		renderer.Present()
	}
}
