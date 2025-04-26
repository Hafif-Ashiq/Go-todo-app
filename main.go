package main

import (
	"fmt"
	"os"
	"todo-app/commands"
	"todo-app/todo"
)

func main() {

	list := todo.GetDataFromFile()

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo-app <command> [arguments]")
		return
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		commands.AddToList(list, args)
	case "remove":
		commands.DeleteFromList(list, args)
	case "list":
		commands.ListItems(list, args)
	case "mark-complete":
		commands.MarkAsComplete(list, args)
	default:
		fmt.Printf("Unknown command: %s\n", command)
	}

	todo.WriteDataToFile(list)

}
