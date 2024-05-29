package ecs

import "reflect"

// Returns the component name from a component type
// eg. ecs.GetComponentName((*ColliderComponent)(nil))
func GetComponentName(c Component) string {
	return reflect.TypeOf(c).String()
}
