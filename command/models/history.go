package models

import "fmt"

type CommandHistory struct {
	history []ICommand
}

func (c *CommandHistory) Push(com ICommand) {
	// Добавить команду в конец массива-истории.
	c.history = append(c.history, com)
	fmt.Println("Команда добавлена в историю")

}
func (c *CommandHistory) Pop() ICommand {
	// Достать последнюю команду из массива-истории.
	if len(c.history) == 0 {
		return nil
	}

	lastCommand := c.history[len(c.history)-1]
	c.history = c.history[:len(c.history)-1]

	fmt.Println("Последняя команда из истории взята", lastCommand)
	return lastCommand
}
