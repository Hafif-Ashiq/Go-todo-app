// # Command for updating todo
package commands

import (
	"fmt"
	"strconv"
	"todo-app/todo"
)

func MarkAsComplete(list *todo.TodoList, args []string) {

	if len(args) < 1 {
		fmt.Println("Usage: mark-complete <id> ")
		return
	}

	ids := args[0]
	id, err := strconv.Atoi(ids)

	if err != nil {
		fmt.Println("Got invalid ID.\n\tTry Again....")
	}

	list.MarkComplete(id)

}
