package entities

import (
	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
)

func CreateRock(manager *ecs.Manager, x, y float64) *ecs.Entity {
	rock := manager.AddEntity()
	rock.AddComponent(&components.SpriteComponent{
		TexturePath: "assets/rock.png",
	})
	rock.AddComponent(&components.TransformComponent{X: x, Y: y, Width: DisplaySize, Height: DisplaySize})
	rock.AddComponent(&components.ColliderComponent{})
	return rock
}
