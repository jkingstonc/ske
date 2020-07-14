package ske

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	PRESSED  = 0x0
	RELEASED = 0x1
	HELD 	 = 0x2
	SCROLLED = 0x3
)

type InputManager struct {
	active *InputNode
}

func (i*InputManager) Update(){
	i.active = nil
}

func SDLKeyToString(keycode sdl.Keycode) string{
	switch int(keycode){
	case 27:
		return "escape"
	case 32:
		return "space"
	default:
		return string(rune(keycode))
	}
}

func SDLButtonToString(button uint8) string{
	return fmt.Sprintf("mouse%d",button)
}

func (i*InputManager) SetActive(button string, interaction uint8, data interface{}){
	switch interaction{
	case PRESSED:
		i.active = &InputNode{
			button:      button,
			pressed:  true,
		}
	case RELEASED:
		i.active = &InputNode{
			button:      button,
			released:  true,
		}
	case HELD:
		i.active = &InputNode{
			button:      button,
			held:  true,
		}
	case SCROLLED:
		i.active = &InputNode{
			button:       button,
			scrolled:  data.(float64),
		}
	}
}

// represents a 'thing' that can be pressed, clicked or held
type InputNode struct {
	button      string
	pressed  bool
	released bool
	held     bool
	scrolled float64
}

func (i InputManager) Button(button string) InputNode{
	if i.active != nil && i.active.button==button{
		return *i.active
	}
	return InputNode{}
}

func (i InputNode) Pressed() bool{return i.pressed}
func (i InputNode) Released() bool{return i.released}
func (i InputNode) Held() bool{return i.held}
func (i InputNode) Scrolled() float64{return i.scrolled}