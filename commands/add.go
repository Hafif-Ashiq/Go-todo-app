// # Command for adding todo
package commands

import (
	"fmt"
	"todo-app/todo"
)

func AddToList(list *todo.TodoList, args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: add <title> ")
		return
	}

	title := args[0]

	list.Add(title)

}
