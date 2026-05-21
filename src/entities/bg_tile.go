package entities

import (
	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
)

func CreateBGTile(manager *ecs.Manager, x, y float64) *ecs.Entity {
	tile := manager.AddEntity()
	tile.AddComponent(&components.SpriteComponent{
		TexturePath: "assets/grass.png",
		Layer:       components.ZBackground,
	})
	tile.AddComponent(&components.TransformComponent{X: x, Y: y, Width: DisplaySize, Height: DisplaySize})
	return tile
}
