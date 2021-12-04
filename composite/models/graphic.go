package models

type Graphic interface {
	Move(dx, dy int)
	Draw()
}
