package ske

import (
	"github.com/veandco/go-sdl2/mix"
	"math"
)

type Audio struct {
	Music   *mix.Music
	Looping bool
	Volume  float64
}

func (a*Audio) Play(){
	if !a.Looping {
		a.Music.Play(0)
	}else{
		a.Music.Play(math.MaxInt32)
	}
}

func (a*Audio) SetVolume(vol float64){
	mix.VolumeMusic(int(vol*100))
}

func (a*Audio) Pause(){}

func (*Audio) Type() uint8 {
	return AUDIO
}


type AudioComponent struct {
	Component
	Audio *Audio
}

func (*AudioComponent) OnLoad(){}
func (*AudioComponent) Update(){}