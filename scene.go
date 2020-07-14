package ske

// a scene represents a state of the game at any one time
type Scene interface {
	Tag() string
	Setup()
	OnLoad()
}

// represents the runtime for a scene (a scene needs an ECS)
type SceneRuntime struct {
	Scene Scene
	ECS   *EntityManager
	// true if we previously exited this scene, and wanted to save the state
	Saved bool
}

type SceneManager struct {
	registered []*SceneRuntime
	active     *SceneRuntime
}

func (s *SceneManager) Update(){
	s.active.ECS.Update()
}

// switch to another scene
func (s *SceneManager) Register(scenes... Scene){
	for _, scene := range scenes {
		runtime := &SceneRuntime{
			Scene: scene,
			ECS:   &EntityManager{},
		}
		s.registered = append(s.registered, runtime)
		// call setup on each scene that is registered
		scene.Setup()
	}
}

// switch to another scene
func (s *SceneManager) ToScene(tag string, save bool){
	// if we don't want to save the scene, then we clear the entities
	if s.active != nil{
		s.active.Saved = save
		if !save{
			s.active.ECS.Entities = nil
		}
	}
	found := false
	// then set the new scene
	for _, sceneRuntime := range s.registered{
		if sceneRuntime.Scene.Tag() == tag{
			// set the active scene to the requested scene
			s.active = sceneRuntime
			// set the ECS to the scene's runtime EntityManager
			ECS = sceneRuntime.ECS
			// then call OnLoad on the scene if the scene wasn't saved, aswel as the components in the scene
			if !sceneRuntime.Saved {
				sceneRuntime.Scene.OnLoad()
				for _, entity := range sceneRuntime.ECS.Entities{
					for _, component := range entity.Components{
						component.OnLoad()
					}
				}
			}
			found = true
		}
	}
	Assert(found, "cannot find scene")
}

// go to the first scene (called when the engine runs)
func (s *SceneManager) ToFirstScene(){
	s.ToScene(s.active.Scene.Tag(), false)
}