package ske

import "os"

var (
	Events *EventManager
	Inputs *InputManager
	// TODO link scenes with ECS correctly0
	ECS    *EntityManager
	Scenes *SceneManager
    Loader *FileManager
)

const (
	// where assets should be loaded from
	AssetsRoot = "F:\\OneDrive\\Programming\\GO\\src\\ske\\"
)


// this is the main driver struct, it will be used to drive the entire engine
type Ske struct {
	running   bool
	runtime   *Runtime
	options   *SkeOptions
}

type SkeOptions struct {
	Title  string
	Width  int
	Height int
}

func NewSKE(options *SkeOptions) *Ske{


	ske := &Ske{
		running:   false,
		runtime: &Runtime{
			scenes:      make(map[string]Scene),
			activeScene: nil,
		},
		options: options,
	}

	ECS = &EntityManager{
		Runtime:  ske.runtime,
		Entities: nil,
	}

	Scenes = &SceneManager{}

	Events = &EventManager{Listeners: make(map[string][]func(event Event))}

	Loader = &FileManager{LoadedFiles: make(map[string]*os.File)}

	return ske
}

// register a scene to the game
func (ske *Ske) RegisterScene(scene Scene){
	ske.runtime.RegisterScene(scene)
}

func (Ske *Ske) Run(){

	// go to the first scene
	Scenes.ToFirstScene()

	Ske.running = true
	// this is the main game loop
	for Ske.running{
		ECS.Update()
	}
	Ske.running = false
}