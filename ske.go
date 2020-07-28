package ske

import "time"

var (
	Engine 		 *Ske
	Events 		 *EventManager
	Inputs 		 *InputManager
	// TODO link scenes with ECS correctly (perhaps ECS should hotswap scenes?)
	ECS    		 *EntityManager
	Scenes 		 *SceneManager
	Kernel 		 *FileManager
	Screen 	     *SDLScreen

	DT           float64
	previousTime time.Time
)

// this is the main driver struct, it will be used to drive the entire engine
type Ske struct {
	running   bool
	options   *SkeOptions
}

type SkeOptions struct {
	Title      string
	Width      int32
	Height     int32
	Resizable  bool
	AssetsRoot string
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
	Kernel = &FileManager{LoadedFiles: make(map[string]Resource)}
	Kernel.Setup()

	return Engine
}

func (ske *Ske) Run(scene string){
	// go to the first scene
	Scenes.ToScene(scene, false)
	ske.running = true
	// this is the main game loop
	for ske.running{
		dt := time.Since(previousTime).Seconds()
		previousTime = time.Now()
		if dt > 0 {
			DT = dt
		}
		Inputs.Update()
		Screen.PollEvents()
		Screen.RendererPrepare()
		Scenes.Update()
		Screen.FetchMeshComponents()
		Screen.RendererFlush()
	}
	ske.running = false
	// close the screen
	Screen.Close()
}