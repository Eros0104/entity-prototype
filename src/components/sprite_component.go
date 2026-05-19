package components

import (
	"fmt"
	"reflect"

	ecs "entity-prototype/src/entity_component_system"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteComponent struct {
	ecs.BaseComponent
	TexturePath string
	Texture     *sdl.Texture

	SpriteSize  *int32
	TotalFrames *int32
	CurrentStep *int32
}

func (c *SpriteComponent) isAnimated() bool {
	return c.SpriteSize != nil && c.TotalFrames != nil && c.CurrentStep != nil
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
	s := int32(1)
	c.CurrentStep = &s
}

func (c *SpriteComponent) Update() {
	if !c.isAnimated() {
		return
	}

	if *c.CurrentStep == *c.TotalFrames-1 {
		*c.CurrentStep = 1
	}

	*c.CurrentStep += 1
}

func (c *SpriteComponent) Draw(renderer *sdl.Renderer) {
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*TransformComponent)

	var currentFrame *sdl.Rect = nil

	if c.isAnimated() {
		currentFrame = &sdl.Rect{
			X: *c.SpriteSize * *c.CurrentStep,
			Y: 0,
			W: *c.SpriteSize,
			H: *c.SpriteSize,
		}
	}

	renderer.Copy(c.Texture, currentFrame, &sdl.Rect{int32(pos.X), int32(pos.Y), int32(pos.Width), int32(pos.Height)})
}

func (c *SpriteComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
