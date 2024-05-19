package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type TransformComponent struct {
	ecs.BaseComponent
	X, Y          float64
	Width, Height float64
}

func (c *TransformComponent) Init() {
	fmt.Println("Position component initialized")
}

func (c *TransformComponent) Update() {

}

func (c *TransformComponent) Draw(renderer *sdl.Renderer) {

}

func (c *TransformComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
