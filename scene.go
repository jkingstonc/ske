package ske

// a scene represents a state of the game at any one time
type Scene interface {
	Tag() string
	Setup()
	OnLoad()
}

type SceneManager struct {
	registered []Scene
	active     Scene
}

// switch to another scene
func (s *SceneManager) Register(scene Scene){
	s.registered = append(s.registered, scene)
	if s.active == nil{
		s.active = scene
	}
	// call setup on each scene that is registered
	scene.Setup()
}

// switch to another scene
func (s *SceneManager) ToScene(tag string, save bool){
	// clear the ECS
	ECS.Entities = nil
	found := false
	// then set the new scene
	for _, scene := range s.registered{
		if scene.Tag() == tag{
			s.active = scene
			// then call OnLoad on the scene
			scene.OnLoad()
			found = true
		}
	}
	Assert(found, "cannot find scene")
}

// go to the first scene (called when the engine runs)
func (s *SceneManager) ToFirstScene(){
	s.ToScene(s.active.Tag(), false)
}