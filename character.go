package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Character struct {
	Position
	Width, Height int
	Collider      ColliderComponent
}

func NewCharacter(x, y float64, width, height int) *Character {
	return &Character{
		Position: Position{x, y},
		Width:    width,
		Height:   height,
		Collider: *NewColliderComponent(int32(x), int32(y), int32(width), int32(height)),
	}
}

func (c *Character) Update() {
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_UP] == 1 {
		c.Y -= 1
	}
	if keys[sdl.SCANCODE_DOWN] == 1 {
		c.Y += 1
	}
	if keys[sdl.SCANCODE_LEFT] == 1 {
		c.X -= 1
	}
	if keys[sdl.SCANCODE_RIGHT] == 1 {
		c.X += 1
	}

	// Update the collider
	c.Collider.rect.X = int32(c.X)
	c.Collider.rect.Y = int32(c.Y)
}

func (c *Character) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.FillRect(&sdl.Rect{int32(c.X), int32(c.Y), int32(c.Width), int32(c.Height)})

	// Draw the collider
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawRect(c.Collider.GetRect())
}
