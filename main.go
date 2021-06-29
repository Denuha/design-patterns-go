package main

import (
	"fmt"

	"github.com/Denuha/design-patterns-go/adapter/round"
	"github.com/Denuha/design-patterns-go/adapter/square"
	adapter "github.com/Denuha/design-patterns-go/adapter/square_adapter"
)

func main() {
	hole := round.NewRoundHole(5)
	rpeg1 := round.NewRoundPeg(6)
	rpeg2 := round.NewRoundPeg(4)

	fmt.Println("fit rpeg1", hole.Fits(*rpeg1)) // false
	fmt.Println("fit rpeg2", hole.Fits(*rpeg2)) // true

	small_sqpeg := square.NewSquarePeg(5)
	large_sqpeg := square.NewSquarePeg(10)

	// hole.Fits(*small_sqpeg) // Ошибка компиляции, несовместимые типы

	small_sqpeg_adapter := adapter.NewSquarePegAdapter(*small_sqpeg)
	large_sqpeg_adapter := adapter.NewSquarePegAdapter(*large_sqpeg)
	fmt.Println(small_sqpeg_adapter, large_sqpeg_adapter)
}
