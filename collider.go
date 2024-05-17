package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ColliderComponent struct {
	rect sdl.Rect
}

func NewColliderComponent(x, y, width, height int32) *ColliderComponent {
	return &ColliderComponent{
		rect: sdl.Rect{x, y, width, height},
	}
}

func (c *ColliderComponent) CheckCollision(other *ColliderComponent) bool {
	return c.rect.HasIntersection(&other.rect)
}

func (c *ColliderComponent) GetRect() *sdl.Rect {
	return &c.rect
}
