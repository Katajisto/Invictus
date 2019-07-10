package main

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

//Returns wheter the program should keep running
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
		}
	}
	return true
}
