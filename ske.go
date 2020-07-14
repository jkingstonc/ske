package ske

var (
	Engine *Ske
	Events *EventManager
	Inputs *InputManager
	// TODO link scenes with ECS correctly0
	ECS    *EntityManager
	Scenes *SceneManager
    Loader *FileManager
	Screen *SDLScreen
)

const (
	// where assets should be loaded from
	AssetsRoot = "F:\\OneDrive\\Programming\\GO\\src\\ske\\examples\\assets\\"
)


// this is the main driver struct, it will be used to drive the entire engine
type Ske struct {
	running   bool
	options   *SkeOptions
}

type SkeOptions struct {
	Title     string
	Width     int32
	Height    int32
	Resizable bool
}

func (s*Ske) ForceStop(){
	s.running = false
}

func (s*Ske) Options() *SkeOptions{
	return s.options
}

func NewSKE(options *SkeOptions) *Ske {

	Engine = &Ske{
		running:   false,
		options: options,
	}

	Screen = &SDLScreen{}
	Screen.Setup()

	Scenes = &SceneManager{}
	Events = &EventManager{Listeners: make(map[string][]func(event Event))}
	Inputs = &InputManager{}
	Loader = &FileManager{LoadedFiles: make(map[string]Resource)}

	return Engine
}

func (ske *Ske) Run(scene string){
	// go to the first scene
	Scenes.ToScene(scene, false)

	ske.running = true
	// this is the main game loop
	for ske.running{
		Inputs.Update()
		Screen.PollEvents()
		Screen.RendererPrepare()
		Scenes.Update()
		Screen.RendererFlush()
	}
	ske.running = false
	// close the screen
	Screen.Close()
}