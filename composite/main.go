package main

import (
	"fmt"

	"github.com/Denuha/design-patterns-go/composite/models"
)

func main() {
	// Инициализация полотна
	canvas := models.NewCompoundGraphic()

	// Инициализация группу компонентов
	selection := models.NewCompoundGraphic()
	dot1 := models.NewDot(100, 100)
	dot2 := models.NewDot(111, 111)
	selection.Add(dot1)
	selection.Add(dot2)

	// ДОбавляем все компоненты на полотно
	circle1 := models.NewCircle(10, 10, 1)
	dot3 := models.NewDot(13, 13)
	canvas.Add(circle1)
	canvas.Add(dot3)
	canvas.Add(selection)

	canvas.Draw() // Выводим полотно
	fmt.Println("=======================")

	selection.Move(15, 15) // Двигаем все компоненты в группе
	canvas.Draw()

	fmt.Println("=======================")
	canvas.Move(-1000, -1000) // Двигаем все компоненты на полотне
	canvas.Draw()
}
