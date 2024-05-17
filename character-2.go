package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Character2 struct {
	Position
	Width, Height int
	Collider      ColliderComponent
}

func NewCharacter2(x, y float64, width, height int) *Character2 {
	return &Character2{
		Position: Position{x, y},
		Width:    width,
		Height:   height,
		Collider: *NewColliderComponent(int32(x), int32(y), int32(width), int32(height)),
	}
}

func (c *Character2) Update(character *Character) {
	if c.Collider.CheckCollision(&character.Collider) {
		fmt.Println("Collision!")
		fmt.Println("!")
	}

	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_W] == 1 {
		c.Y -= 1
	}
	if keys[sdl.SCANCODE_S] == 1 {
		c.Y += 1
	}
	if keys[sdl.SCANCODE_A] == 1 {
		c.X -= 1
	}
	if keys[sdl.SCANCODE_D] == 1 {
		c.X += 1
	}

	// Update the collider
	c.Collider.rect.X = int32(c.X)
	c.Collider.rect.Y = int32(c.Y)
}

func (c *Character2) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.FillRect(&sdl.Rect{int32(c.X), int32(c.Y), int32(c.Width), int32(c.Height)})

	// Draw the collider
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.DrawRect(c.Collider.GetRect())
}
