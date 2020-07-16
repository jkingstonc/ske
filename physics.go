package ske

type CollisionEvent struct {
	// the entity that caused the collision
	Collider *Entity
	// the entity that received the collision
	Collidee *Entity
}

// add to an entity to trigger collision events
type BoxColliderComponent struct{

}
