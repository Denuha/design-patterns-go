package square

// Квадратный колышек
type SquarePeg struct {
	width int
}

func NewSquarePeg(radius int) *SquarePeg {
	return &SquarePeg{
		width: radius,
	}
}

func (r *SquarePeg) GetWidth() int {
	return r.width
}
