package models

type Strategy interface {
	Execute(a, b int) int
}

type StrategyAdd struct {
	Strategy
}

func (s *StrategyAdd) Execute(a, b int) int {
	return a + b
}

func NewStrategyAdd() Strategy {
	return &StrategyAdd{}
}

/////////////////
type StrategySub struct {
	Strategy
}

func (s *StrategySub) Execute(a, b int) int {
	return a - b
}

func NewStrategySub() Strategy {
	return &StrategySub{}
}

////////////////
type StrategyMult struct {
	Strategy
}

func (s *StrategyMult) Execute(a, b int) int {
	return a * b
}

func NewStrategyMult() Strategy {
	return &StrategyMult{}
}
