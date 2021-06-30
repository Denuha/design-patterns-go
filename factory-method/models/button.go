package models

import "fmt"

type Button struct {
	name string
	IButton
}

// Общий интерфейс кнопок
type IButton interface {
	render()
	onClick()
	//Hide()
}

func newButton() *Button {
	return &Button{
		name: "Just button",
	}
}

func (b *Button) Hide() {
	fmt.Println("I`m hided")
}

func (b *Button) ChangeName(newName string) {
	b.name = newName
}
func (b *Button) GetName() string {
	return b.name
}

/*************************************************/
type WindowsButton struct {
	Button
}

func (w *WindowsButton) render() {
	fmt.Println("Отрисовка кнопки windows")
	fmt.Printf("hello, I`m %s\n", w.name)
}

func newWindowsButton() *WindowsButton {
	b := WindowsButton{}
	b.name = "Windows button"

	return &b
}

/**********************************************/
type HTMLButtom struct {
	Button
}

func (h *HTMLButtom) render() {
	fmt.Println("Отрисовка кнопки HTML")
	fmt.Printf("hello, I`m %s\n", h.name)
}

func newHTMLButton() *HTMLButtom {
	b := HTMLButtom{}
	b.name = "HTML button"

	return &b
}
