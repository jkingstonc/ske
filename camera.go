package ske

import "reflect"

// the camera's position is the transform of the entity it is attached to
type CameraComponent struct {
	Component
	t *TransformComponent
}

func (m*CameraComponent) OnLoad(){
	m.t = m.Entity.GetComponent(reflect.TypeOf(&TransformComponent{})).(*TransformComponent)
}
func (m*CameraComponent) Update(){}

func BasicCamera() *Entity{
	cam := ECS.NewEntity("camera")
	cam.Attach(&CameraComponent{
		Component: cam.NewComponent(),
	}, &TransformComponent{
		Component: cam.NewComponent(),
		Pos:   V2(1,1),
		Scale: V2(1,1),
	})
	return cam
}