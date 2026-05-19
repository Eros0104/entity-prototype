package components

import (
	"fmt"
	"math"
	"reflect"

	ecs "entity-prototype/src/entity_component_system"

	"github.com/veandco/go-sdl2/sdl"
)

type InputHandlerComponent struct {
	ecs.BaseComponent
	Speed float64
}

func (c *InputHandlerComponent) Init() {
	fmt.Println("InputHandlerComponent initialized")
}

func (c *InputHandlerComponent) Update(dt float64) {
	x, y := 0.0, 0.0
	keys := sdl.GetKeyboardState()
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*TransformComponent)

	if keys[sdl.SCANCODE_W] == 1 {
		y -= 1
	}
	if keys[sdl.SCANCODE_S] == 1 {
		y += 1
	}
	if keys[sdl.SCANCODE_A] == 1 {
		x -= 1
	}
	if keys[sdl.SCANCODE_D] == 1 {
		x += 1
	}

	length := math.Sqrt(y*y + x*x)

	if length > 0 {
		x = x / length
		y = y / length
	}

	pos.Y += y * c.Speed * dt
	pos.X += x * c.Speed * dt
}

func (c *InputHandlerComponent) Draw(renderer *sdl.Renderer) {
}

func (c *InputHandlerComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
