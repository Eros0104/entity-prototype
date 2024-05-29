package main

import (
	"entity-prototype/src/components"
	ecs "entity-prototype/src/entity_component_system"
	"fmt"
	"reflect"
)

// adjusts the position of c2 to push it away from c1 based on their AABB collision.
func aabb(c1 *components.TransformComponent, c2 *components.TransformComponent) {
	// Calculate the half-widths and half-heights
	halfWidthC1 := c1.Width / 2
	halfHeightC1 := c1.Height / 2
	halfWidthC2 := c2.Width / 2
	halfHeightC2 := c2.Height / 2

	// Calculate the centers of c1 and c2
	centerC1X := c1.X + halfWidthC1
	centerC1Y := c1.Y + halfHeightC1
	centerC2X := c2.X + halfWidthC2
	centerC2Y := c2.Y + halfHeightC2

	// Calculate the differences in the centers
	dx := centerC2X - centerC1X
	dy := centerC2Y - centerC1Y

	// Calculate the overlap on the x and y axes
	overlapX := (halfWidthC1 + halfWidthC2) - abs(dx)
	overlapY := (halfHeightC1 + halfHeightC2) - abs(dy)

	// If there is no overlap, return early
	if overlapX <= 0 || overlapY <= 0 {
		return
	}

	// Determine the direction to push c2
	if overlapX < overlapY {
		// Push in x direction
		if dx > 0 {
			c2.X += overlapX // Push c2 to the right
		} else {
			c2.X -= overlapX // Push c2 to the left
		}
	} else {
		// Push in y direction
		if dy > 0 {
			c2.Y += overlapY // Push c2 down
		} else {
			c2.Y -= overlapY // Push c2 up
		}
	}
}

func ManageCollisions(collidersGroup []*ecs.Entity) {
	colliderType := ecs.GetComponentName((*components.ColliderComponent)(nil))
	for i := 0; i < len(collidersGroup); i++ {
		for j := i + 1; j < len(collidersGroup); j++ {
			if collidersGroup[i].GetComponent(colliderType).(*components.ColliderComponent).CheckCollision(collidersGroup[j].GetComponent(colliderType).(*components.ColliderComponent)) {
				fmt.Println("Collision detected")
				aabb(collidersGroup[i].GetComponent(colliderType).(*components.ColliderComponent).GetEntity().GetComponent(reflect.TypeOf((*components.TransformComponent)(nil)).String()).(*components.TransformComponent), collidersGroup[j].GetComponent(colliderType).(*components.ColliderComponent).GetEntity().GetComponent(reflect.TypeOf((*components.TransformComponent)(nil)).String()).(*components.TransformComponent))
			} else {
				fmt.Println("No collision")
			}
		}
	}
}
