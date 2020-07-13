package common

import "ske/ske"

// a transform is used on any entity that has a spacial position
type TransformComponent struct {
	ske.Component
}

func (*TransformComponent) OnLoad(){}
func (*TransformComponent) Update(){}