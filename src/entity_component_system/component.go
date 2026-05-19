package ecs

import "github.com/veandco/go-sdl2/sdl"

type Component interface {
	Init()
	Update(dt float64)
	Draw(renderer *sdl.Renderer)
	SetEntity(*Entity)
}

type BaseComponent struct {
	entity *Entity
}

func (b *BaseComponent) SetEntity(e *Entity) {
	b.entity = e
}

func (b *BaseComponent) GetEntity() *Entity {
	return b.entity
}

func (b *BaseComponent) Init()                         {}
func (b *BaseComponent) Update(dt float64)             {}
func (b *BaseComponent) Draw(renderer *sdl.Renderer)   {}
