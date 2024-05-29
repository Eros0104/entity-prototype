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

func (m *Manager) GetComponentGroup(typeName string) []*Entity {
	var group []*Entity
	for _, e := range m.entities {
		if e.HasComponent(typeName) {
			group = append(group, e)
		}
	}
	return group
}
