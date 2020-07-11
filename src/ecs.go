package src

// this is the main controller for the Entity Component model.
// note this is not an ECS, it is an EC. the entity logic is
// wrapped inside the components.
type GOManager struct {}


type IComponent interface {
	// this is called every in-game tick (20 times a second)
	Tick()
}

type Component struct {
	// each component has a pointer to the game object
	GameObject *GameObject
}

type IGameObject interface {

}

type GameObject struct {
	// unique ID
	ID uint32
	// list of components
	Components []IComponent
}

func (GameObject *GameObject) AddComponent(component IComponent){

}