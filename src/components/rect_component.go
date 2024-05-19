package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type RectComponent struct {
	ecs.BaseComponent
	R, G, B, A uint8
}

func (c *RectComponent) Init() {
	fmt.Println("RectComponent initialized")
}

func (c *RectComponent) Update() {

}

func (c *RectComponent) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(c.R, c.G, c.B, c.A)
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*TransformComponent)
	renderer.DrawRect(&sdl.Rect{int32(pos.X), int32(pos.Y), int32(pos.Width), int32(pos.Height)})
}

func (c *RectComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
