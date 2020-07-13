package ske

import (
	"reflect"
)

// this is the main controller for the Entity Component model.
// note this is not an ECS, it is an EC. the entity logic is
// wrapped inside the components.
type EntityManager struct {
	// reference to the runtime
	Runtime *Runtime
	// store the array of Entities
	Entities []*Entity
}

func (EntityManager *EntityManager) NewEntity(tag string) *Entity{
	Entity := &Entity{
		ID:  0,
		Tag: tag,
		Components: nil,
	}
	EntityManager.Entities = append(EntityManager.Entities, Entity)
	return Entity
}

func (EntityManager *EntityManager) Update(){
	for _, Entity := range EntityManager.Entities{
		Entity.Update()
	}
}

type IComponent interface {
	OnLoad()
	Update()
}

type Component struct {
	// each component has a pointer to the game object
	Entity *Entity
}

type Entity struct {
	// unique ID
	ID uint32
	// string tag
	Tag string
	// the components attached to the entity
	Components []IComponent
}

func (e *Entity) Update(){
	for _, component := range e.Components{
		component.Update()
	}
}

func (e *Entity) NewComponent() Component{
	return Component{
		Entity:  e,
	}
}

func (e *Entity) Attach(component IComponent){
	e.Components = append(e.Components, component)
}

// TODO
func (e *Entity) Detach(component IComponent){
}

func (e *Entity) GetComponent(t reflect.Type) IComponent{
	for _, c := range e.Components{
		if reflect.TypeOf(c) == t{
			return c
		}
	}
	return nil
}