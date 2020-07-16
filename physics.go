package ske

import "reflect"

type CollisionEvent struct {
	// the entity that caused the collision
	Collider *Entity
	// the entity that received the collision
	Collidee *Entity
	ColliderTransform TransformComponent
	CollideeTransform TransformComponent
}

// add to an entity to trigger collision events
type BoxColliderComponent struct{
	Component
	// the collider boundaries are indipendent of the scale
	Boundaries Vec
	// cause solid collisions
	Solid bool
}

func (c*BoxColliderComponent) OnLoad(){}
func (c*BoxColliderComponent) Update(){
	// get our transform
	t := c.Component.Entity.GetComponent(reflect.TypeOf(&TransformComponent{})).(*TransformComponent)
	// check other entity
	for _, entity := range ECS.Entities{
		if entity.Active && entity != c.Entity{
			// check if the
			otherC := entity.GetComponent(reflect.TypeOf(&BoxColliderComponent{}))
			otherT := entity.GetComponent(reflect.TypeOf(&TransformComponent{}))


			if otherC!=nil && otherT!=nil {
				otherC := otherC.(*BoxColliderComponent)
				otherT := otherT.(*TransformComponent)

				// the transform scale is the diamater!
				ourLeft := t.Pos.X - c.Boundaries.X/2
				ourRight := t.Pos.X + c.Boundaries.X/2
				ourTop := t.Pos.Y + c.Boundaries.Y/2
				ourBottom := t.Pos.Y - c.Boundaries.Y/2

				theirLeft := otherT.Pos.X - otherC.Boundaries.X/2
				theirRight := otherT.Pos.X + otherC.Boundaries.X/2
				theirTop := otherT.Pos.Y + otherC.Boundaries.Y/2
				theirBottom := otherT.Pos.Y - otherC.Boundaries.Y/2

				if colliding(ourLeft, ourRight, ourTop, ourBottom, theirLeft, theirRight, theirTop, theirBottom) {
					// trigger a collision event
					Events.Send(CollisionEvent{
						Collider: c.Entity,
						Collidee: entity,
						ColliderTransform: *t,
						CollideeTransform: *otherT,
					})
				}
			}
		}
	}
}

func colliding(aLeft, aRight, aTop, aBottom, bLeft, bRight, bTop, bBottom float64) bool{
	return aLeft < bRight && aRight > bLeft && aTop > bBottom && aBottom < bTop
}
