package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type FollowComponent struct {
	ecs.BaseComponent
	Destination *ecs.Entity
	Speed       float64
}

func (c *FollowComponent) Init() {
	fmt.Println("RectComponent initialized")
}

func (c *FollowComponent) Update() {
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*TransformComponent)

	if c.Destination != nil {
		destination := c.Destination.GetComponent(typeName).(*TransformComponent)

		// moves gradually towards the destination
		if pos.X < destination.X {
			pos.X += c.Speed
		}
		if pos.X > destination.X {
			pos.X -= c.Speed
		}
		if pos.Y < destination.Y {
			pos.Y += c.Speed
		}
		if pos.Y > destination.Y {
			pos.Y -= c.Speed
		}
	}
}

func (c *FollowComponent) Draw(renderer *sdl.Renderer) {

}

func (c *FollowComponent) Follow(other *ecs.Entity) {
	c.Destination = other
}

func (c *FollowComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
