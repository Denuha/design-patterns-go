package models

import "fmt"

type Editor struct {
	Text string
}

func (e *Editor) GetSelection() string {
	// Вернуть выбранный текст.
	fmt.Println("GetSelection()")
	return e.Text
}

func (e *Editor) DeleteSelection() {
	// Удалить выбранный текст.
	fmt.Println("DeleteSelection()")
	e.Text = ""
}

func (e *Editor) ReplaceSelection(text string) {
	// Вставить текст из буфера обмена в текущей позиции.
	fmt.Println("ReplaceSelection()")
	e.Text += text
}
