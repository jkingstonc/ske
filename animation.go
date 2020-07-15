package ske

import "reflect"

// TODO animations
// an animation has a series of textures
type Animation struct {
	Tag      string
	// the source textures that the animation uses
	Atlas    *Atlas
	// indexes into the atlas that the animation uses
	Frames   []uint
	// the speed at which the animation plays
	Speed    float64
	// whether the animation should loop
	Looping  bool
	// the active animation index
	Index    int
}

func NewAnimation(tag string, atlas *Atlas, frames []uint, speed float64, looping bool) *Animation{
	return &Animation{
		Tag:     tag,
		Atlas:   atlas,
		Frames:  frames,
		Speed:   speed,
		Looping: looping,
		Index:   0,
	}
}

type AnimationComponent struct {
	Component
	// the map of animation resource
	Animations map[string]*Animation
	// the active animation
	Active   *Animation
	// the mesh we are updating
	Mesh      *MeshComponent
}

func (a*AnimationComponent) AddAnimations(animations... *Animation){
	for _, animation := range animations{
		a.Animations[animation.Tag] = animation
	}
	a.Active = animations[0]
}

func (a*AnimationComponent) OnLoad() {
	a.Mesh = a.Entity.GetComponent(reflect.TypeOf(&MeshComponent{})).(*MeshComponent)
}
func (a*AnimationComponent) Update() {
	// we need to get the texture from the animation
	// set the atlas position to the correct index
	a.Active.Atlas.SetAtlasPosition(a.Active.Frames[a.Active.Index])

	// loop to the next animation
	if a.Active.Index<len(a.Active.Frames)-1 {
		a.Active.Index++
	} else {
		// we have finished the animtion
		if a.Active.Looping {
			a.Active.Index = 0
		}
	}
}