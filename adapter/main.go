package main

import (
	"fmt"

	squareAdapter "github.com/Denuha/design-patterns-go/adapter/adapter"
	"github.com/Denuha/design-patterns-go/adapter/round"
	"github.com/Denuha/design-patterns-go/adapter/square"
)

func main() {
	hole := round.NewRoundHole(5)
	rpeg1 := round.NewRoundPeg(6)
	rpeg2 := round.NewRoundPeg(4)

	fmt.Println("fit rpeg1:", hole.Fits(*rpeg1)) // false
	fmt.Println("fit rpeg2:", hole.Fits(*rpeg2)) // true

	smallSqpeg := square.NewSquarePeg(5)
	largeSqpeg := square.NewSquarePeg(10)

	// hole.Fits(*small_sqpeg) // Ошибка компиляции, несовместимые типы

	smallSqpegAdapter := squareAdapter.NewSquarePegAdapter(*smallSqpeg)
	largeSqpegAdapter := squareAdapter.NewSquarePegAdapter(*largeSqpeg)

	fmt.Println("fit small speg:", hole.Fits(smallSqpegAdapter.RoundPeg)) // true
	fmt.Println("fit large speg:", hole.Fits(largeSqpegAdapter.RoundPeg)) // false
}
