package ske

const (
	WORLD_TARGET  = 0x0
	SCREEN_TARGET = 0x1

	WORLD_ORDER = 0x0
	UI_ORDER    = 0x10
)

// a mesh component must be added to any entity that is to be drawn
type MeshComponent struct {
	Component
	Target   uint8
	Drawable Drawable
	Order    uint32
}

func (*MeshComponent) OnLoad() {}
func (*MeshComponent) Update() {}