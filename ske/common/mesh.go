package common

import "ske/ske"

// a mesh component must be added to any entity that is to be drawn
type MeshComponent struct {
	ske.Component
}

func (*MeshComponent) OnLoad(){}
func (*MeshComponent) Update(){}