package models

type Application struct {
	Clipboard string
	history   CommandHistory
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) ExecuteCommand(c ICommand) {
	if c.Execute() {
		a.history.Push(c)
	}
}

func (s *Application) Undo() {
	command := s.history.Pop()
	if command != nil {
		command.Undo()
	}
}
