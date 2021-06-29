package main

import (
	"fmt"

	"github.com/Denuha/design-patterns-go/adapter/round"
)

func main() {
	hole := round.NewRoundHole(5)
	rpeg := round.NewRoundPeg(5)

	fmt.Println(hole.GetRadius())
	fmt.Println(rpeg.GetRadius())
}
