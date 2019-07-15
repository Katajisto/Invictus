package invictus

import(
	"github.com/veandco/go-sdl2/sdl"
)

var animManager EntityHandler

type Animation struct {
	Frames []Rect
}

type EntityHandler struct {
	Tick           int32
	GlobalEntities []*Entity
}

type Entity struct {
	Texture    *sdl.Texture
	Animations []Animation
	Duration   int32
	AnimMode   int32
}

type EntityFrame struct {
	Texture *sdl.Texture
	What   *Rect
}

func animTick() {
	animManager.Tick += 1
	if animManager.Tick > 99 {
		animManager.Tick = 0
	}
}

func (e Entity) curDrawable() EntityFrame {
	var ef EntityFrame
	ef.Texture = e.Texture
	duration  := e.Duration
	animTick := animManager.Tick/duration
	if animTick >= int32(len(e.Animations[e.AnimMode].Frames)) {
		animTick %= int32(len(e.Animations[e.AnimMode].Frames))
	}
	WR := &e.Animations[e.AnimMode].Frames[animTick]
	ef.What = WR
	return ef
}
