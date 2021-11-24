package main

import (
	"fmt"

	"github.com/Denuha/design-patterns-go/strategy/models"
)

var (
	actAdd  = "add"
	actSub  = "sub"
	actMult = "mult"
)

func main() {
	context := models.NewContext()
	a := 1
	b := 2
	action := actSub

	switch action {
	case actAdd:
		context.SetStrategy(models.NewStrategyAdd())
	case actSub:
		context.SetStrategy(models.NewStrategySub())
	case actMult:
		context.SetStrategy(models.NewStrategyMult())
	}

	result := context.ExecuteStrategy(a, b)
	fmt.Println(result)
}
