package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type ColliderComponent struct {
	ecs.BaseComponent
	X, Y, Width, Height float64
}

func (c *ColliderComponent) Init() {
	fmt.Println("ColliderComponent initialized")
}

func (c *ColliderComponent) Update() {

}

func (c *ColliderComponent) Draw(renderer *sdl.Renderer) {

}

func (c *ColliderComponent) CheckCollision(other *ColliderComponent) bool {
	return c.X < other.X+other.Width &&
		c.X+c.Width > other.X &&
		c.Y < other.Y+other.Height &&
		c.Y+c.Height > other.Y
}

func (c *ColliderComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
