package models

import (
	"fmt"
)

type Command struct {
	app    Application
	editor Editor
	backup string
}

type ICommand interface {
	SaveBackup()
	Undo()
	Execute() bool
}

//--------------------------

type CopyCommand struct {
	Command
}

type ICopyCommand interface {
	ICommand
}

func (c *CopyCommand) Execute() bool {
	c.app.Clipboard = c.editor.GetSelection()
	fmt.Println("Выполнена команда копирования")
	return false
}

//--------------------------
type CutCommand struct {
	Command
}
type ICutCommand interface {
	ICommand
}

func (c *CutCommand) Execute() bool {
	c.SaveBackup()

	c.app.Clipboard = c.editor.GetSelection()
	c.editor.DeleteSelection()
	fmt.Println("Выполнена команда вырезать")
	return true
}

//--------------------------
type PasteCommand struct {
	Command
}
type IPasteCommand interface {
	ICommand
}

func (c *PasteCommand) Execute() bool {
	c.SaveBackup()

	c.editor.ReplaceSelection(c.app.Clipboard)
	fmt.Println("Выполнена команда вставить")
	return true
}

//--------------------------
type UndoCommand struct {
	Command
}
type IUndoCommand interface {
	ICommand
}

func (c *UndoCommand) Execute() bool {
	c.app.Undo()
	fmt.Println("Выполнена команда отменить")
	return false
}

//--------------------------
//
//
/* Command methods */
func NewCommand( /*app models.Application, editor models.Editor*/ ) ICommand {
	return &Command{
		//app:    app,
		//editor: editor,
	}
}

func (c *Command) SaveBackup() {
	fmt.Println("Выполнена комманда 'SaveBackup'")
	c.backup = c.editor.Text
}

func (c *Command) Undo() {
	fmt.Println("Выполнена комманда 'Undo'")
	c.editor.Text = c.backup
}

func (c *Command) Execute() bool {
	fmt.Println("Абстрактная комманда")
	return false
}
