package ske

// a scene represents a state of the game at any one time
type Scene interface {
	Tag() string
	Setup()
	OnLoad()
}


type SceneManager struct {}

// scene manager
var Scenes *SceneManager

// switch to another scene
func (s *SceneManager) ToScene(tag string, save bool){}