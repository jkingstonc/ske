package ske

import (
	"encoding/json"
	"reflect"
)

// this is the main controller for the Entity Component model.
// note this is not an ECS, it is an EC. the entity logic is
// wrapped inside the components.
type EntityManager struct {
	// store the array of Entities
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

func (e *Entity) NewComponent() Component {
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

func (e *Entity) GetComponent(t reflect.Type) IComponent {
	for _, c := range e.Components{
		if reflect.TypeOf(c) == t{
			return c
		}
	}
	return nil
}