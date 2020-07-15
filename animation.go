package ske

import (
	"reflect"
	"time"
)

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
	previousTime time.Time
}

func (a*AnimationComponent) AddAnimations(animations... *Animation){
	for _, animation := range animations{
		a.Animations[animation.Tag] = animation
	}
	a.Active = animations[0]
}

func (a*AnimationComponent) OnLoad() {
	a.Mesh = a.Entity.GetComponent(reflect.TypeOf(&MeshComponent{})).(*MeshComponent)
	a.Mesh.Drawable = a.Active.Atlas
	a.previousTime = time.Now()
}
func (a*AnimationComponent) Update() {
	// we need to get the texture from the animation
	// set the atlas position to the correct index
	a.Active.Atlas.SetAtlasPosition(a.Active.Frames[a.Active.Index])

	// if its time to advance
	if time.Since(a.previousTime).Seconds() > a.Active.Speed {
		a.previousTime = time.Now()
		// advance
		a.Active.Index++

		// if we have finished
		if a.Active.Index >= len(a.Active.Frames){
			if a.Active.Looping{
				a.Active.Index = 0
			}else{
				a.Active.Index = len(a.Active.Frames)-1
			}
		}
	}
}