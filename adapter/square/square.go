package square

// Квадратный колышек
type SquarePeg struct {
	width float64
}

type ISquarePeg interface {
	GetWidth() float64
}

func NewSquarePeg(radius float64) *SquarePeg {
	return &SquarePeg{
		width: radius,
	}
}

func (r *SquarePeg) GetWidth() float64 {
	return r.width
}
