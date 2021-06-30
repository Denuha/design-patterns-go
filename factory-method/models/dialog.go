package models

type Dialog struct {
	IDialog
	Button
}

type IDialog interface {
	Render()
	createButton() IButton
	GetButton() Button
}

func (d *Dialog) newDialog() IDialog {
	return &Dialog{}
}

// abstract method
func (d *Dialog) Render() {
}

// abstract method
func (d *Dialog) createButton() IButton {
	return nil
}

// abstract method
func (d *Dialog) GetButton() Button {
	return d.Button
}

/********************************************/
type WindowsDialog struct {
	Dialog
}

type IWindowsDialog interface {
	Render()
	createButton() IButton
	GetButton() Button
}

func NewWindowsDialog() IWindowsDialog {
	return &WindowsDialog{}
}

func (d *WindowsDialog) createButton() IButton {
	var button = newWindowsButton()
	d.Button = Button{name: button.name}
	return button
}

// надо убрать этот метод от сюда, его нужно перенести в Dialog
func (d *WindowsDialog) Render() {
	button := d.createButton()
	button.render()
}

// надо убрать этот метод от сюда, его нужно перенести в Dialog
func (d *WindowsDialog) GetButton() Button {
	return d.Button
}

/********************************************/
type HTMLDialog struct {
	Dialog
}

type IHTMLDialog interface {
	Render()
	createButton() IButton
	GetButton() Button
}

func NewHTMLDialog() IHTMLDialog {
	return &HTMLDialog{}
}

func (d *HTMLDialog) createButton() IButton {
	var button = newHTMLButton()
	d.Button = Button{name: button.name}
	return button
}

// надо убрать этот метод от сюда, его нужно перенести в Dialog
func (d *HTMLDialog) Render() {
	button := d.createButton()
	button.render()
}

// надо убрать этот метод от сюда, его нужно перенести в Dialog
func (d *HTMLDialog) GetButton() Button {
	return d.Button
}
