package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type FollowComponent struct {
	ecs.BaseComponent
	Destination *ColliderComponent
	Speed       float64
}

func (c *FollowComponent) Init() {
	fmt.Println("RectComponent initialized")
}

func (c *FollowComponent) Update() {
	typeName := reflect.TypeOf((*ColliderComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*ColliderComponent)

	if c.Destination != nil {

		// moves gradually towards the destination
		if c.Destination.X > pos.X {
			pos.X += c.Speed
		}
		if c.Destination.X < pos.X {
			pos.X -= c.Speed
		}
		if c.Destination.Y > pos.Y {
			pos.Y += c.Speed
		}
		if c.Destination.Y < pos.Y {
			pos.Y -= c.Speed
		}
	}
}

func (c *FollowComponent) Draw(renderer *sdl.Renderer) {

}

func (c *FollowComponent) Follow(other *ColliderComponent) {
	c.Destination = other
}

func (c *FollowComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
