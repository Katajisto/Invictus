package invictus

//Contains the functions related to event handling (including input)

import (
	"github.com/veandco/go-sdl2/sdl"
)

var keyStates map[sdl.Keycode]bool

func init() {
	keyStates = make(map[sdl.Keycode]bool)
}

func KeyPressed(kc sdl.Keycode) bool {
	return keyStates[kc]
}

//Returns wheter the program should keep running, required to run on start of every frame.
func PollEvent() bool {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			return false
		case *sdl.KeyboardEvent:
			if t.Type == sdl.KEYDOWN {
				keyStates[t.Keysym.Sym] = true
			}
			if t.Type == sdl.KEYUP   {
				keyStates[t.Keysym.Sym] = false
			}
		case *sdl.MouseButtonEvent :
			if t.Type == sdl.MOUSEBUTTONDOWN {
				switch t.Button {
				case sdl.BUTTON_LEFT:
					MB1 = true
				case sdl.BUTTON_RIGHT:
					MB2 = true
				case sdl.BUTTON_MIDDLE:
					MB3 = true
				}
			}
			if t.Type == sdl.MOUSEBUTTONUP {
				switch t.Button {
				case sdl.BUTTON_LEFT:
					MB1 = false
				case sdl.BUTTON_RIGHT:
					MB2 = false
				case sdl.BUTTON_MIDDLE:
					MB3 = false
				}
			}
		case *sdl.MouseMotionEvent:
			MOUSEX = t.X
			MOUSEY = t.Y
		}
		
	}
	return true
}
