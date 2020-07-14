package ske

// a mesh component must be added to any entity that is to be drawn
type MeshComponent struct {
	Component
	*Texture
}

func (*MeshComponent) OnLoad() {}
func (*MeshComponent) Update() {}