package models

import "fmt"

type Dot interface {
	Graphic
}

type dot struct {
	X, Y int
}

func (d *dot) Move(dx, dy int) {
	d.X += dx
	d.Y += dy
}

func (d *dot) Draw() {
	fmt.Printf("dot is drew x=%d y=%d\n", d.X, d.Y)
}

func NewDot(x, y int) Dot {
	return &dot{
		X: x,
		Y: y,
	}
}
