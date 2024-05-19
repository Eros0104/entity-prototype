package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type InputHandlerComponent struct {
	ecs.BaseComponent
	Speed float64
}

func (c *InputHandlerComponent) Init() {
	fmt.Println("InputHandlerComponent initialized")
}

func (c *InputHandlerComponent) Update() {
	keys := sdl.GetKeyboardState()
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*TransformComponent)

	if keys[sdl.SCANCODE_W] == 1 {
		pos.Y -= c.Speed
	}
	if keys[sdl.SCANCODE_S] == 1 {
		pos.Y += c.Speed
	}
	if keys[sdl.SCANCODE_A] == 1 {
		pos.X -= c.Speed
	}
	if keys[sdl.SCANCODE_D] == 1 {
		pos.X += c.Speed
	}
}

func (c *InputHandlerComponent) Draw(renderer *sdl.Renderer) {
	// Drawing logic if needed
}

func (c *InputHandlerComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
