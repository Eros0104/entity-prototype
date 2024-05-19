package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"

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
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	myPos := c.GetEntity().GetComponent(typeName).(*TransformComponent)
	otherPos := other.GetEntity().GetComponent(typeName).(*TransformComponent)

	return myPos.X < otherPos.X+other.Width &&
		myPos.X+c.Width > otherPos.X &&
		myPos.Y < otherPos.Y+other.Height &&
		myPos.Y+c.Height > otherPos.Y
}

func (c *ColliderComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
