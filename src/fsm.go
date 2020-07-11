package src

// this file handles all Finite State Machine behaviour

// represents a Finite State Machine
type IFSM interface {
	Process()
}

type IState interface {
	Process()
}

// implements IFSM
type GameObjectFSM struct {
	GameObject *GameObject
	CurrentState IState
	States map[string]*GameObjectState
}

func (gameObjectFSM *GameObjectFSM) NewState(tag string) *GameObjectState{
	state := &GameObjectState{
		GameObjectFSM: gameObjectFSM,
	}
	gameObjectFSM.States[tag] = state
	if gameObjectFSM.CurrentState == nil{
		gameObjectFSM.CurrentState = state
	}
	return state
}

func (gameObjectFSM *GameObjectFSM) Process(){
	// call process on the current state
	gameObjectFSM.CurrentState.Process()
}

// implements IState
type GameObjectState struct {
	GameObjectFSM *GameObjectFSM
	// the components attached to this current state
	Components []IComponent
}

func (gameObjectState *GameObjectState) Process(){
	// call tick on each component
	for _, component := range gameObjectState.Components{
		// this should use timing so it only ticks 20 times a second
		component.Tick()
	}
}

func (gameObjectState *GameObjectState) With(component IComponent) *GameObjectState{
	// use reflection to set the GameObject field of the component
	Field("GameObject", FieldIndex(0, ValuePtr(component))).Set(Value(gameObjectState.GameObjectFSM.GameObject))
	gameObjectState.Components = append(gameObjectState.Components, component)
	return gameObjectState
}