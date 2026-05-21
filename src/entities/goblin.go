package entities

import (
	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
)

func CreateGoblin(manager *ecs.Manager, target *ecs.Entity, x, y float64) *ecs.Entity {
	goblin := manager.AddEntity()
	goblin.AddComponent(&components.SpriteComponent{TexturePath: "assets/goblin.png"})
	goblin.AddComponent(&components.TransformComponent{X: x, Y: y, Width: DisplaySize, Height: DisplaySize})
	goblin.AddComponent(&components.ColliderComponent{})
	goblin.AddComponent(&components.RectComponent{R: 0, G: 0, B: 255, A: 255})
	goblin.AddComponent(&components.FollowComponent{Destination: target, Speed: 30})
	return goblin
}
