package invictus

import(
	"github.com/veandco/go-sdl2/sdl"
)


//Mouse variables to be used locally within the engine.
var MOUSEX,MOUSEY int32
var MB1, MB2, MB3 bool

//Gets the relative mouse state.
func RelativeMouseState() (x,y,s int32) {
	x,y,ss := sdl.GetRelativeMouseState()
	s = int32(ss)
	return
}

//Returns mouse position relative to when last asked.
func RelMPos() (x,y int32) {
	x,y,_ = RelativeMouseState()
	return
}


//Returns the current mouse position, relative to the window.
func MPos() (int32, int32) {
	return MOUSEX,MOUSEY
}


func LogMouse() {
	Log(MOUSEX,MOUSEY,MB1,MB2,MB3)
}

//Function wich returns if MouseButton is Pressed. Argument specifies the button 1-3
//On unknown button number, returns false.
func MBP(w int32) bool {
	switch w {
	case 1:
		return MB1
	case 2:
		return MB2
	case 3:
		return MB3
	default:
		return false
	}
}

