package ecs

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type Entity struct {
	manager    *Manager
	components map[string]Component
}

func NewEntity(manager *Manager) *Entity {
	return &Entity{
		manager:    manager,
		components: make(map[string]Component),
	}
}

func (e *Entity) Update() {
	for _, c := range e.components {
		c.Update()
	}
}

func (e *Entity) Draw(renderer *sdl.Renderer) {
	for _, c := range e.components {
		c.Draw(renderer)
	}
}

func (e *Entity) AddComponent(c Component) {
	typeName := reflect.TypeOf(c).String()
	fmt.Println("Adding component:", typeName)
	e.components[typeName] = c
	c.SetEntity(e)
	c.Init()
}

func (e *Entity) GetComponent(typeName string) Component {
	return e.components[typeName]
}

func (e *Entity) GetManager() *Manager {
	return e.manager
}

func (e *Entity) HasComponent(typeName string) bool {
	_, ok := e.components[typeName]
	return ok
}
