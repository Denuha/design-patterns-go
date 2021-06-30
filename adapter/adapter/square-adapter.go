package adapter

import (
	"math"

	"github.com/Denuha/design-patterns-go/adapter/round"
	"github.com/Denuha/design-patterns-go/adapter/square"
)

type SquarePegAdapter struct {
	peg square.SquarePeg
}

type ISquarePegAdapter interface {
	round.IRoundPeg // IRoundPeg contains interface with GetRadius()
}

func NewSquarePegAdapter(peg square.SquarePeg) ISquarePegAdapter {
	return &SquarePegAdapter{
		peg: peg,
	}
}

// Здесь мы переопределяем функцию GetRadius() из IRoundPeg
func (s *SquarePegAdapter) GetRadius() float64 {
	// Pythagorean theorem
	return float64(s.peg.GetWidth()) * math.Sqrt(2) / 2
}
