package main

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var( 
	BLUE  = Color{0,0,255,255}
	RED   = Color{255,0,0,255}
	GREEN = Color{0,255,0,255}
	BLACK = Color{0,0,0,255}
	WHITE = Color{255,255,255,255}
)

func (col Color) toi() (uint8, uint8, uint8, uint8) {
	return col.R, col.G, col.B, col.A
}
