package main

import (
	"errors"
	"fmt"

	"github.com/Denuha/design-patterns-go/factory-method/models"
)

func main() {
	var dialog models.IDialog

	//myOS := "Web"
	myOS := "Windows"

	if myOS == "Windows" {
		dialog = models.NewWindowsDialog() // создаем диалог для Виндовс
	} else if myOS == "Web" {
		dialog = models.NewHTMLDialog() // Создаем диалог для Веб
	} else {
		fmt.Printf("%v", errors.New("Unknown OS\n"))
		return
	}

	dialog.Render() // Отрисовываем созданную кнопку

	// Так же можно к кнопке обращаться и проводить с ней манипуляции, не меняя диалог
	b := dialog.GetButton()
	fmt.Println(b.GetName())
	b.ChangeName("asd")
	fmt.Println(b.GetName())

	b.Hide()
}
