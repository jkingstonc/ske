package ske

// a transform is used on any entity that has a spacial position
type TransformComponent struct {
	Component
	Pos   Vec
	Rot   Vec
	Scale Vec
}

func (*TransformComponent) OnLoad() {}
func (*TransformComponent) Update() {}

func (t*TransformComponent) Translate(other Vec){
	t.Pos = t.Pos.Add(other)
}