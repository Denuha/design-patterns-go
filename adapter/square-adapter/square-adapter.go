package square_adapter

import(
	"github.com/Denuha/design-patterns-go/adapter/round"
	"github.com/Denuha/design-patterns-go/adapter/square"
)
type SquarePegAdapter stuct {
	round.RoundPeg
	peg square.SquarePeg
}

func NewSquarePegAdapter(peg square.SquarePeg) *SquarePegAdapter {
	return &SquarePegAdapter{
		peg = peg,
	}
}