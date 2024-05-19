package main

import (
	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	const FPS uint32 = 60
	const frameDelay uint32 = 1000 / FPS

	var frameStart, frameTime uint32

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

	manager := ecs.Manager{Renderer: renderer}

	// Create two entities
	player := manager.AddEntity()
	follower := manager.AddEntity()
	wall := manager.AddEntity()

	// Add SpriteComponent to player
	playerSprite := &components.SpriteComponent{TexturePath: "assets/player.png"}

	// Add TransformComponent to player and wall
	playerTransform := &components.TransformComponent{X: 100, Y: 100, Width: 50, Height: 50}
	wallTransform := &components.TransformComponent{X: 30, Y: 30, Width: 50, Height: 50}
	followerTransform := &components.TransformComponent{X: 200, Y: 200, Width: 50, Height: 50}

	// Add ColliderComponent to player and wall
	playerCollider := &components.ColliderComponent{}
	wallCollider := &components.ColliderComponent{}
	followerCollider := &components.ColliderComponent{}

	// Add input handler component to player
	playerInputHandler := &components.InputHandlerComponent{Speed: 5}

	// Add rect component to player and wall
	playerRect := &components.RectComponent{R: 255, G: 0, B: 0, A: 255}
	wallRect := &components.RectComponent{R: 0, G: 255, B: 0, A: 255}
	followerRect := &components.RectComponent{R: 0, G: 0, B: 255, A: 255}

	// Add follow component to follower
	followerFollow := &components.FollowComponent{Destination: player, Speed: 0.5}

	// Add components to entities
	player.AddComponent(playerSprite)

	player.AddComponent(playerTransform)
	wall.AddComponent(wallTransform)
	follower.AddComponent(followerTransform)

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
		frameStart = sdl.GetTicks()

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
		manager.Draw()

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

		// Calculate frame time
		frameTime = sdl.GetTicks() - frameStart

		// Delay the frame if needed
		if frameDelay > frameTime {
			sdl.Delay(frameDelay - frameTime)
		}
	}
}
