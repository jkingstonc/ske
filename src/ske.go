package src

// this is the main driver struct, it will be used to drive the entire engine
type Ske struct {
	Running bool
	GOManager *GOManager
}

func (Ske *Ske) Run(){
	// this is the main game loop
	for Ske.Running{
		// update the GOManager
	}
}