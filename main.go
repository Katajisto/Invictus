package main

/*

Invictus game engine. Written in go, using sdl2 bindings.

Tuomas Katajisto - t@ktj.st

*/

import (
	"github.com/veandco/go-sdl2/sdl"
	// "fmt"		
)

func main() {
	Init("Test!", 700, 500)

	tt := LoadTexture("t_assets/asd.png")
	
	var xpos int32 = 0
	var movement int32 = 1
	var dire int32 = 0
	
	for PollEvent() {
		Clear(BLACK)
		xpos += movement
		if(xpos == 0 || xpos == 500) {
			movement *= -1
			if dire == 1 {
				dire = 0
			} else {
				dire = 1
			}
		}
		Dtex(tt, &Rect{xpos,0,200,200}, nil, 0, nil, dire)
		Show()
		SleepMs(10)
	}
	
	Window.Destroy()

	sdl.Quit()
}
