package main

import(
	"os"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

func Init(title string, w,h int32) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to init: %s\n", err)
		os.Exit(1)
	}
	windowAndRenderer(title, w, h)
}

var Window *sdl.Window
var Renderer *sdl.Renderer

func windowAndRenderer(title string, w,h int32) {
	window, err := sdl.CreateWindow(title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
	Window = window
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(2)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	Renderer = renderer
	
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
}

func SleepMs(amount int32) {
	time.Sleep(time.Millisecond * time.Duration(amount))
}
