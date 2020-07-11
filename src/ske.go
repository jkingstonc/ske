package src

// this is the main driver struct, it will be used to drive the entire engine
type Ske struct {
	running bool
	goManager *GOManager
}

func NewSKE() *Ske{
	return &Ske{
		running:   false,
		goManager: &GOManager{},
	}
}

func (Ske *Ske) NewGameObject() *GameObject{
	return Ske.goManager.NewGameObject()
}

func (Ske *Ske) MakePrefab(tag string, gameObject *GameObject){}

func (Ske *Ske) NewPrefab(tag string){}

func (Ske *Ske) Run(){
	Ske.running = true
	// this is the main game loop
	for Ske.running{
		Ske.goManager.Process()
	}
	Ske.running = false
}