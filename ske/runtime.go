package ske


// the runtime is an API into the runtime of the engine.
// it is passed to scenes so programmers can setup their game
type Runtime struct {
	EntityManager   *EntityManager
	//screen      *Screen
	scenes      map[string]Scene
	activeScene Scene
}

func (runtime *Runtime) ToScene(tag string){
	runtime.activeScene = runtime.scenes[tag]
	runtime.activeScene.Setup(runtime)
}

func (runtime *Runtime) GetEntity(tag string) *Entity {
	return nil
}

func (runtime *Runtime) NewEntity(tag string) *Entity{
	return runtime.EntityManager.NewEntity(tag)
}

func (runtime *Runtime) MakePrefab(tag string, Entity *Entity){}

func (runtime *Runtime) NewPrefab(tag string){}