package ske

import "reflect"

// this file contains all the code to handle the UI components

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
		}, &ConstraintComponent{Component: text.NewComponent()},
		   &AnchorComponent{Component: text.NewComponent(), Anchor: anchor},
	)
	return text
}

// ConstraintComponent handles UI constraints.
// it's job is to update the transform of the entity that it is attached to, given the parent constraints.
type ConstraintComponent struct {
	Component
	// used for caching the entity transform
	transform *TransformComponent
}

func (c*ConstraintComponent) OnLoad() {
	c.transform = c.Entity.GetComponent(reflect.TypeOf(&TransformComponent{})).(*TransformComponent)
}
func (c*ConstraintComponent) Update() {
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
	Anchor uint8
	// used for caching the entity transform
	transform *TransformComponent
}

func (a*AnchorComponent) OnLoad() {
	a.transform = a.Entity.GetComponent(reflect.TypeOf(&TransformComponent{})).(*TransformComponent)
}
func (a*AnchorComponent) Update() {
	anchor := a.getAnchorPos()
	a.transform.Pos = anchor
}

// getAnchorPos calculates the coordinate given the anchor mode
func (a*AnchorComponent) getAnchorPos() Vec{
	// x & y here should be the constraints
	x, y := a.transform.Pos.X, a.transform.Pos.Y

	w, h := Screen.Window.GetSize()

	if a.Anchor & CENTER == 1 {
		return V2(float64(w/2), float64(h/2))
	}

	if a.Anchor & TOP > 0 {
		y = 0
	}else if a.Anchor & BOTTOM > 0{
		y = float64(h)
	}

	if a.Anchor & LEFT > 0 {
		x=0
	}else if a.Anchor & RIGHT > 0{
		x=float64(w)
	}

	return V2(x,y)
}