package components

import (
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteComponent struct {
	ecs.BaseComponent
	TexturePath string
	Texture     *sdl.Texture
}

func (c *SpriteComponent) Init() {
	fmt.Println("RectComponent initialized")

	// get the renderer from the manager
	renderer := c.GetEntity().GetManager().Renderer

	// loads the texture
	texture, err := img.LoadTexture(renderer, c.TexturePath)
	if err != nil {
		fmt.Println("Failed to load texture:", err)
		return
	}
	c.Texture = texture
}

func (c *SpriteComponent) Update() {

}

func (c *SpriteComponent) Draw(renderer *sdl.Renderer) {
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*TransformComponent)

	renderer.Copy(c.Texture, nil, &sdl.Rect{int32(pos.X), int32(pos.Y), int32(pos.Width), int32(pos.Height)})
}

func (c *SpriteComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
