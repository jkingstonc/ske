package ske

import (
	"reflect"
)

// this file contains all the code to handle the UI components

// NewContainerEntity creates an entity that acts as a container for other UI elements.
// the job of the container is to order and best-fit the UI elements that are children of this container.
func NewContainerEntity(){}

// NewTextEntity creates an entity that can be attached to a game object, it will be rendered to the screen
func NewTextEntity(s string, anchor uint8) *Entity{
	text := ECS.NewEntity("text")
	text.Attach(
		&MeshComponent{
			Component: text.NewComponent(),
			Target:    SCREEN_TARGET,
			Order:     UI_ORDER,
		}, &TextComponent{
			Component: text.NewComponent(),
			Text: s,
			Font: "anon.ttf",
			Size: 20,
		}, &TransformComponent{
			Component: text.NewComponent(),
		},
		   &AnchorComponent{Component: text.NewComponent(), Anchor: anchor},
	)
	return text
}

const (
	// different anchor positions
	CENTER = 0x1 << 0
	LEFT   = 0x1 << 1
	RIGHT  = 0x1 << 2
	TOP    = 0x1 << 3
	BOTTOM = 0x1 << 4
)

// AnchorComponent handles UI anchors.
// it's job is to update the transform of the entity that it is attached to, given the parent constraints.
type AnchorComponent struct {
	Component
	// the anchor mode (CENTER, LEFT etc)
	Anchor     uint8
	// the boundaries that the anchor will place the UI element within.
	// note, by default, this is the size of the screen.
	// also note, X&Y is the top left, Z&W is the bottom right.
	Boundaries Vec
	// used for caching the entity transform
	transform *TransformComponent
}

func (a*AnchorComponent) OnLoad() {
	a.transform = a.Entity.GetComponent(reflect.TypeOf(&TransformComponent{})).(*TransformComponent)
	a.Boundaries = V4(0,0, float64(Engine.options.Width), float64(Engine.options.Height))
}
func (a*AnchorComponent) Update() {
	anchor := a.getAnchorPos()
	a.transform.Pos = anchor
}

// getAnchorPos calculates the coordinate given the anchor mode
func (a*AnchorComponent) getAnchorPos() Vec{

	// x & y is the center of the anchor boundaries
	x := (a.Boundaries.Z - a.Boundaries.X) / 2
	y := (a.Boundaries.W - a.Boundaries.Y) / 2

	if a.Anchor & CENTER == 1 {
		return V2(x, y)
	}

	if a.Anchor & TOP > 0 {
		y = a.Boundaries.Y
	}else if a.Anchor & BOTTOM > 0{
		y = a.Boundaries.W
	}

	if a.Anchor & LEFT > 0 {
		x=a.Boundaries.X
	}else if a.Anchor & RIGHT > 0{
		x=a.Boundaries.Z
	}

	return V2(x,y)
}