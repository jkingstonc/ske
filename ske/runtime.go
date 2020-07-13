package ske


// the runtime is an API into the runtime of the engine.
// it is passed to scenes so programmers can setup their game
type Runtime struct {
	EntityManager   *EntityManager
	//screen      *Screen
	scenes      map[string]Scene
	activeScene Scene
}

func (runtime *Runtime) RegisterScene(scene Scene){
	runtime.scenes[scene.Tag()] = scene
	if runtime.activeScene == nil{
		runtime.activeScene = scene
	}
	scene.Setup()
}

// switch to a different scene, and optionally save the scene state
// TODO should there be an entity manager for each scene?
func (runtime *Runtime) ToScene(tag string, save bool){
	// clear the entity map
	ECS.Entities = nil
	// then load the scene
	runtime.activeScene = runtime.scenes[tag]



	// call on load on the scene
	runtime.activeScene.OnLoad()
	// call on load on the entities in the scene
	for _, e := range ECS.Entities{
		for _, c := range e.Components{
			c.OnLoad()
		}
	}
}

func (runtime *Runtime) GetEntity(tag string) *Entity {
	for _, entity := range ECS.Entities{
		if entity.Tag == tag{
			return entity
		}
	}
	return nil
}

func (runtime *Runtime) NewEntity(tag string) *Entity{
	return runtime.EntityManager.NewEntity(tag)
}

func (runtime *Runtime) MakePrefab(tag string, Entity *Entity){}

func (runtime *Runtime) NewPrefab(tag string){}