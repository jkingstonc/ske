package ske

import (
	"encoding/json"
	"reflect"
)

// this is the main controller for the Entity Component model.
// note this is not an ECS, it is an EC. the entity logic is
// wrapped inside the components.
type EntityManager struct {
	// store the array of active Entities
	Entities []*Entity
}

func Serialize(entity *Entity) []byte{
	bytes, err := json.Marshal(entity)
	Assert(err==nil, "couldn't marshal entity to JSON")
	return bytes
}

func Deserialize(path string) *Entity{
	bytes := ReadRaw(path)
	var entity *Entity
	err := json.Unmarshal(bytes, &entity)
	Log(err)
	Assert(err==nil, "couldn't unmarshal entity JSON")
	return entity
}

func (entityManager *EntityManager) EntityWithTag(tag string) *Entity{
	for _, entity := range entityManager.Entities{
		if entity.Tag == tag{
			return entity
		}
	}
	return nil
}

// instantiate a prefab
func (EntityManager *EntityManager) Instantiate(tag string) *Entity {
	// deserialize an entity from a file path
	entity := Deserialize("prefabs/"+tag+".prefab")
	return entity
}

// turn an entity into a prefab
func (EntityManager *EntityManager) MakePrefab(entity *Entity) {
	// serialize the prefab and store it in the assets package
	bytes := Serialize(entity)
	WriteRaw("prefabs/"+entity.Tag+".prefab", bytes)
}

func (EntityManager *EntityManager) NewEntity(tag string) *Entity {
	Entity := &Entity{
		ID:  0,
		Active: true,
		Tag: tag,
		Components: nil,
	}
	EntityManager.Entities = append(EntityManager.Entities, Entity)
	return Entity
}

func (EntityManager *EntityManager) Update(){
	for _, entity := range EntityManager.Entities{
		if entity.Active {
			entity.Update()
		}
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
	// should the entity be updated
	Active bool
	// the components attached to the entity
	Components []IComponent
	// the children attached to the entity
	Children   []*Entity
	// the parent of this entity, can be nil
	Parent     *Entity
}

func (e *Entity) Update(){
	for _, component := range e.Components{
		component.Update()
	}
}

func (e *Entity) NewComponent() Component {
	return Component{
		Entity:  e,
	}
}

func (e *Entity) Attach(components... IComponent){
	for _, component := range components {
		e.Components = append(e.Components, component)
	}
}

// TODO
func (e *Entity) Detach(component IComponent){
}

func (e *Entity) GetComponent(t reflect.Type) IComponent {
	for _, c := range e.Components{
		if reflect.TypeOf(c) == t{
			return c
		}
	}
	return nil
}

func (e *Entity) AddChild(entities... *Entity) {
	for _, entity := range entities {
		e.Children = append(e.Children, entity)
	}
}

func (e *Entity) GetChildren() []*Entity {
	return e.Children
}

// enable or disable an entity.
// this will enable or disable all the children
func (e *Entity) SetActive(active bool){
	e.Active = active
	for _, child := range e.Children{
		child.SetActive(active)
	}
}