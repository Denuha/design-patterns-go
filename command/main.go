package main

import (
	"fmt"

	"github.com/Denuha/design-patterns-go/command/models"
)

// https://refactoring.guru/ru/design-patterns/command

func main() {
	app := models.NewApplication()
	app.Clipboard = "string"

	commandCopy := models.ICopyCommand(new(models.CopyCommand))
	app.ExecuteCommand(commandCopy) // Выполняем команду копирования
	fmt.Println("-------------")
	commandCut := models.ICutCommand(new(models.CutCommand))
	app.ExecuteCommand(commandCut) // Выполняем команду вырезания
	fmt.Println("-------------")
	commandPaste := models.IPasteCommand(new(models.PasteCommand))
	app.ExecuteCommand(commandPaste) // Выполняем команду вставки
	fmt.Println("-------------")
	commandUndo := models.IUndoCommand(new(models.UndoCommand))
	app.ExecuteCommand(commandUndo) // Выполняем команду отмены
}
