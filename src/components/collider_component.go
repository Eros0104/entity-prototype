package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type ColliderComponent struct {
	ecs.BaseComponent
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

	return myPos.X < otherPos.X+otherPos.Width &&
		myPos.X+myPos.Width > otherPos.X &&
		myPos.Y < otherPos.Y+otherPos.Height &&
		myPos.Y+myPos.Height > otherPos.Y
}

func (c *ColliderComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
