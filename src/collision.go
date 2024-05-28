package main

import (
	"entity-prototype/src/components"
)

// push_collider adjusts the position of c2 to push it away from c1 based on their AABB collision.
func push_collider(c1 *components.TransformComponent, c2 *components.TransformComponent) {
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
