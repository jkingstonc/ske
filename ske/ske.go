package ske

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
	return &Ske{
		running:   false,
		runtime: &Runtime{
			EntityManager:   &EntityManager{},
			//screen:      nil,
			scenes:      make(map[string]Scene),
			activeScene: nil,
		},
	}
}

// register a scene to the game
func (ske *Ske) RegisterScene(scene Scene){
	ske.runtime.scenes[scene.Tag()] = scene
	if ske.runtime.activeScene == nil{
		ske.runtime.activeScene = scene
	}
}

func (Ske *Ske) Run(){

	// go to the first scene
	Ske.runtime.ToScene(Ske.runtime.activeScene.Tag())

	Ske.running = true
	// this is the main game loop
	for Ske.running{
		Ske.runtime.EntityManager.Update()
	}
	Ske.running = false
}