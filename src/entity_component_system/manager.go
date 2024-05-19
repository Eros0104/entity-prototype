package ecs

import "github.com/veandco/go-sdl2/sdl"

type Manager struct {
	entities []*Entity
	Renderer *sdl.Renderer
}

func (m *Manager) Update() {
	for _, e := range m.entities {
		e.Update()
	}
}

func (m *Manager) Draw() {
	for _, e := range m.entities {
		e.Draw(m.Renderer)
	}
}

func (m *Manager) AddEntity() *Entity {
	entity := NewEntity(m)
	m.entities = append(m.entities, entity)
	return entity
}
