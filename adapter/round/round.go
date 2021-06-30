package round

// Круглое отверстие
type RoundHole struct {
	radius float64
}

type IRoundHole interface {
	GetRadius() float64
	Fits(peg RoundPeg) bool
}

// Круглый колышек
type RoundPeg struct {
	radius float64
}

type IRoundPeg interface {
	GetRadius() float64
}

//
//
/* Functions for RoundHole*/
func NewRoundHole(radius float64) *RoundHole {
	return &RoundHole{
		radius: radius,
	}
}

func (r *RoundHole) GetRadius() float64 {
	return r.radius
}

func (r *RoundHole) Fits(peg IRoundPeg) bool {
	return r.GetRadius() >= peg.GetRadius()
}

//
//
/* Functions for RoundPeg*/
func NewRoundPeg(radius float64) *RoundPeg {
	return &RoundPeg{
		radius: radius,
	}
}

func (r *RoundPeg) GetRadius() float64 {
	return r.radius
}
