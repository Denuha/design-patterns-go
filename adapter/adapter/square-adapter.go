package adapter

import (
	"math"

	"github.com/Denuha/design-patterns-go/adapter/round"
	"github.com/Denuha/design-patterns-go/adapter/square"
)

type SquarePegAdapter struct {
	RoundPeg round.RoundPeg
	peg      square.SquarePeg
}

func NewSquarePegAdapter(peg square.SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{
		peg:      peg,
		RoundPeg: *round.NewRoundPeg(float64(peg.GetWidth()) * math.Sqrt(2) / 2),
	}
}

// Хотелось реализовать с помощью данного метода (как бы реализуя наследование, но не получилось)
func (s *SquarePegAdapter) GetRadius() float64 {
	// Pythagorean theorem
	return 0 //float64(s.peg.GetWidth()) * math.Sqrt(2) / 2
}
