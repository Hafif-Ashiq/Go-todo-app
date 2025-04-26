// # Command for listing all TODOs
package commands

import (
	"fmt"
	"todo-app/todo"
)

func ListItems(list *todo.TodoList, args []string) {

	if len(args) > 0 {
		fmt.Println("Usage: list ")
		return
	}

	list.List()
}
