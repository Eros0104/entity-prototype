package entities

import (
	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
)

func CreatePlayer(manager *ecs.Manager, x, y float64) *ecs.Entity {
	spriteSize := int32(16)
	idleFrames := int32(1)
	walkFrames := int32(8)

	player := manager.AddEntity()
	player.AddComponent(&components.SpriteComponent{
		TexturePath: "assets/orangutan-sprite.png",
		Animations: []components.SpriteAnimator{
			{Name: "idle", SpriteSize: &spriteSize, TotalFrames: &idleFrames},
			{Name: "walking", SpriteSize: &spriteSize, TotalFrames: &walkFrames},
		},
	})
	player.AddComponent(&components.TransformComponent{X: x, Y: y, Width: DisplaySize, Height: DisplaySize})
	player.AddComponent(&components.ColliderComponent{})
	player.AddComponent(&components.InputHandlerComponent{Speed: 300})
	player.AddComponent(&components.RectComponent{R: 255, G: 0, B: 0, A: 255})
	return player
}
