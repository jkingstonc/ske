package src

// this is the main controller for the Entity Component model.
// note this is not an ECS, it is an EC. the entity logic is
// wrapped inside the components.
type GOManager struct {
	// store the array of GameObjects
	GameObjects []*GameObject
}

func (goManager *GOManager) NewGameObject() *GameObject{
	gameObject := &GameObject{
		ID:  0,
		FSM: &GameObjectFSM{
			States: make(map[string]*GameObjectState),
		},
	}
	goManager.GameObjects = append(goManager.GameObjects, gameObject)
	return gameObject
}

func (GOManager *GOManager) Process(){
	for _, gameObject := range GOManager.GameObjects{
		gameObject.FSM.Process()
	}
}

type IComponent interface {
	Tick()
}

type Component struct {
	// each component has a pointer to the game object
	GameObject *GameObject
}

type GameObject struct {
	// unique ID
	ID uint32
	// the FSM that controlls the gameObjects state
	FSM *GameObjectFSM
}

// add a state configuration for the GameObject FSM
func (gameObject *GameObject) NewState(id string) *GameObjectState{
	return gameObject.FSM.NewState(id)
}

// switch to the state given the state ID
func (gameObject *GameObject) ToState(id string) *GameObjectState{
	return nil
}