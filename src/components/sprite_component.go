package components

import (
	"fmt"
	"reflect"

	ecs "entity-prototype/src/entity_component_system"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const animFPS = 8.0

type SpriteComponent struct {
	ecs.BaseComponent
	TexturePath string
	Texture     *sdl.Texture

	Animations       []SpriteAnimator
	currentAnimation string
	accumulator      float64
	currentStep      int32
}

type SpriteAnimator struct {
	Name        string
	SpriteSize  *int32
	TotalFrames *int32
	CurrentStep *int32
	accumulator float64
}

func (c *SpriteComponent) isAnimated() bool {
	return len(c.Animations) > 0
}

func (c *SpriteComponent) getAnimationByName(name string) *SpriteAnimator {
	for i := 0; i < len(c.Animations); i++ {
		if c.Animations[i].Name == name {
			return &c.Animations[i]
		}
	}

	return nil
}

func (c *SpriteComponent) getCurrentAnimation() *SpriteAnimator {
	return c.getAnimationByName(c.currentAnimation)
}

func (c *SpriteComponent) Play(animName string) {
	if c.currentAnimation == animName {
		return
	}
	if c.getAnimationByName(animName) == nil {
		return
	}

	c.currentAnimation = animName
	c.currentStep = 1
	c.accumulator = 0
}

func (c *SpriteComponent) Init() {
	fmt.Println("RectComponent initialized")

	renderer := c.GetEntity().GetManager().Renderer

	texture, err := img.LoadTexture(renderer, c.TexturePath)
	if err != nil {
		fmt.Println("Failed to load texture:", err)
		return
	}
	c.Texture = texture
	c.currentStep = 1

	if c.isAnimated() {
		c.currentAnimation = c.Animations[0].Name
	}
}

func (c *SpriteComponent) Update(dt float64) {
	if !c.isAnimated() {
		return
	}

	curAnim := c.getCurrentAnimation()

	c.accumulator += dt
	frameDuration := 1.0 / animFPS
	for c.accumulator >= frameDuration {
		c.accumulator -= frameDuration
		c.currentStep += 1
		if c.currentStep >= *curAnim.TotalFrames {
			c.currentStep = 1
		}
	}
}

func (c *SpriteComponent) Draw(renderer *sdl.Renderer) {
	typeName := reflect.TypeOf((*TransformComponent)(nil)).String()
	pos := c.GetEntity().GetComponent(typeName).(*TransformComponent)

	var currentFrame *sdl.Rect = nil

	if c.isAnimated() {
		curAnim := c.getCurrentAnimation()
		currentFrame = &sdl.Rect{
			X: *curAnim.SpriteSize * (c.currentStep - 1),
			Y: 0,
			W: *curAnim.SpriteSize,
			H: *curAnim.SpriteSize,
		}
	}

	renderer.Copy(c.Texture, currentFrame, &sdl.Rect{int32(pos.X), int32(pos.Y), int32(pos.Width), int32(pos.Height)})
}

func (c *SpriteComponent) SetEntity(e *ecs.Entity) {
	c.BaseComponent.SetEntity(e)
}
