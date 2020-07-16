package ske

import "reflect"

// a transform is used on any entity that has a spacial position
type TransformComponent struct {
	Component
	Pos   Vec
	Rot   float64
	Scale Vec
}

func (*TransformComponent) OnLoad() {}
func (t*TransformComponent) Update() {
	for _, child := range t.Entity.Children{
		if childTransform:=child.GetComponent(reflect.TypeOf(&TransformComponent{})).(*TransformComponent); childTransform != nil{
			*childTransform = *t
		}
	}
}

func (t*TransformComponent) Translate(other Vec){
	t.Pos = t.Pos.Add(other)
}