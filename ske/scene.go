package ske

// a scene represents a state of the game at any one time
type Scene interface {
	Tag() string
	Setup(runtime *Runtime)
}