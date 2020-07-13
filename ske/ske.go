package ske



var (
	Events *EventManager
	Inputs *InputManager
	ECS    *EntityManager
    Loader *FileManager
)

const (
	// where assets should be loaded from
	AssetsRoot = "./"
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

	Events = &EventManager{Listeners: make(map[string][]func(event Event))}

	Loader = &FileManager{}

	return ske
}

// register a scene to the game
func (ske *Ske) RegisterScene(scene Scene){
	ske.runtime.RegisterScene(scene)
}

func (Ske *Ske) Run(){

	// go to the first scene
	Ske.runtime.ToScene(Ske.runtime.activeScene.Tag(), false)

	Ske.running = true
	// this is the main game loop
	for Ske.running{
		ECS.Update()
	}
	Ske.running = false
}