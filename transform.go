package ske

// a transform is used on any entity that has a spacial position
type TransformComponent struct {
	Component
}

func (*TransformComponent) OnLoad() {}
func (*TransformComponent) Update() {}