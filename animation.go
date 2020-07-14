package ske

import "reflect"

// TODO animations
// an animation has a series of textures
type Animation struct {
	// this should be []Drawable
	Textures []*Texture
	// the speed at which the animation plays
	Speed    float64
	// whether the animation should loop
	Looping  bool
}

type AnimationComponent struct {
	Component
	// the map of animation resource
	Animation map[string]*Animation
	// the active animation
	Active   *Animation
	// the mesh we are updating
	Mesh      *MeshComponent
}

func (a*AnimationComponent) OnLoad() {
	a.Mesh = a.Entity.GetComponent(reflect.TypeOf(&MeshComponent{})).(*MeshComponent)
}
func (a*AnimationComponent) Update() {}