package models

import "fmt"

type Circle interface {
	Dot
}

type circle struct {
	dot
	Radius int
}

func (c *circle) Draw() {
	fmt.Printf("Circle is drew x=%d y=%d, r=%d\n", c.X, c.Y, c.Radius)
}

func NewCircle(x, y, radius int) Circle {
	dot := dot{
		X: x,
		Y: y,
	}
	return &circle{
		dot:    dot,
		Radius: radius,
	}
}
