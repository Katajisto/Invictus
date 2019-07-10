package main

import(
	"fmt"
	"os"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

type Point struct {
	X int32
	Y int32 
}

func (p *Point) toSdl() *sdl.Point {
	if p == nil {
		return nil
	}
	return &sdl.Point{p.X,p.Y}
}

func (r *Rect)  toSdl() *sdl.Rect  {
	if r == nil  {
		return nil
	}
	return &sdl.Rect{r.X, r.Y, r.W, r.H}
}

func DRect(x,y,w,h int32, col Color) {
	Renderer.SetDrawColor(col.toi())
	Renderer.FillRect(&sdl.Rect{x,y,w,h})
}

func Clear(col Color) {
	Renderer.SetDrawColor(col.toi())
	Renderer.Clear()
}

func Show() {
	Renderer.Present()
}

//Give a hint to sdl about the render quality. "0","1","2"
func Qhint(q string) {
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, q)
}

func LoadTexture(file string) *sdl.Texture {
	surfaceImg, err := img.Load(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load image", err)
		os.Exit(4)
	}
	textureImg, err := Renderer.CreateTextureFromSurface(surfaceImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	surfaceImg.Free()
	return textureImg
}

func Dtex(tex *sdl.Texture, where *Rect, what *Rect, angle float64, center *Point, flip int32) {
	switch flip {
	case 1:
		Renderer.CopyEx(tex, what.toSdl(), where.toSdl(), angle, center.toSdl(), sdl.FLIP_HORIZONTAL)
	case 2:
		Renderer.CopyEx(tex, what.toSdl(), where.toSdl(), angle, center.toSdl(), sdl.FLIP_VERTICAL)
	default:
		Renderer.CopyEx(tex, what.toSdl(), where.toSdl(), angle, center.toSdl(), sdl.FLIP_NONE)		
	}
}

func Sdtex(tex *sdl.Texture, where *Rect) {
	Dtex(tex, where, nil, 0, nil, 0)
}
