package main

import (
	"fmt"
	"reflect"

	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
	"entity-prototype/src/entities"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	const FPS uint32 = 60
	const frameDelay uint32 = 1000 / FPS

	var frameStart, frameTime uint32
	var lastTime uint32

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

	player := entities.CreatePlayer(&manager, 100, 100)
	entities.CreateGoblin(&manager, player, 200, 200)

	wall := manager.AddEntity()
	wall.AddComponent(&components.TransformComponent{X: 30, Y: 30, Width: entities.DisplaySize, Height: entities.DisplaySize})
	wall.AddComponent(&components.ColliderComponent{})
	wall.AddComponent(&components.RectComponent{R: 0, G: 255, B: 0, A: 255})

	for i := 0.0; i < 26; i++ {
		for j := 0.0; j < 18; j++ {
			entities.CreateBGTile(&manager, entities.TileStep*i, entities.TileStep*j)
		}
	}

	colliderType := reflect.TypeOf((*components.ColliderComponent)(nil)).String()
	collidersGroup := manager.GetComponentGroup(colliderType)

	// Main loop
	lastTime = sdl.GetTicks()
	running := true
	for running {
		frameStart = sdl.GetTicks()
		dt := float64(frameStart-lastTime) / 1000.0
		lastTime = frameStart

		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}

		renderer.SetDrawColor(0, 0, 0, 255)
		renderer.Clear()

		manager.Update(dt)
		manager.Draw()

		ManageCollisions(collidersGroup)

		renderer.Present()

		frameTime = sdl.GetTicks() - frameStart
		if frameDelay > frameTime {
			sdl.Delay(frameDelay - frameTime)
		}
	}
}
