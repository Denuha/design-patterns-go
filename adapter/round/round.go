package round

// Круглое отверстие
type RoundHole struct {
	radius int
}

// Круглый колышек
type RoundPeg struct {
	radius int
}

//
//
/* Functions for RoundHole*/
func NewRoundHole(radius int) *RoundHole {
	return &RoundHole{
		radius: radius,
	}
}

func (r *RoundHole) GetRadius() int {
	return r.radius
}

func (r *RoundHole) Fits(peg RoundPeg) bool {
	return r.GetRadius() >= peg.GetRadius()
}

//
//
/* Functions for RoundPeg*/
func NewRoundPeg(radius int) *RoundPeg {
	return &RoundPeg{
		radius: radius,
	}
}

func (r *RoundPeg) GetRadius() int {
	return r.radius
}
