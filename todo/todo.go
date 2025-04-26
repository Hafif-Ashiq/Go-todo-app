package todo

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Todo struct {
	ID        int
	Title     string
	Completed bool
}

type TodoList struct {
	Items []Todo
}

func NewTodoList() *TodoList {
	return &TodoList{Items: []Todo{}}
}

func (t *TodoList) Add(title string) {
	// Append the todo to the list
	var len int = len(t.Items)
	var id int = 1
	if len > 0 {
		id = t.Items[len-1].ID + 1
	}

	t.Items = append(t.Items, Todo{ID: id, Title: title, Completed: false})
	fmt.Println("Added the task Successfully")
}

func (t *TodoList) Remove(id int) {
	// Remove the todo from the list
	var removed bool = false
	for i, item := range t.Items {
		if item.ID == id {
			t.Items = append(t.Items[:i], t.Items[i+1:]...)
			removed = true
			break
		}
	}

	if removed {
		fmt.Println("Task removed Successfully")
	} else {
		fmt.Println("Task not found")
	}
}

func (t *TodoList) List() {
	// List all the tasks

	fmt.Println("-------------------- TASKS --------------------")
	fmt.Println("ID \t Title \t\t\t Completed")
	for _, item := range t.Items {
		fmt.Printf("%v \t %s \t\t %v\n", item.ID, item.Title, item.Completed)
	}
	fmt.Println("-----------------------------------------------")

}

func (t *TodoList) MarkComplete(id int) {
	// Mark todo as completed
	var updated bool = false
	for i, item := range t.Items {

		if item.ID == id {
			t.Items[i].Completed = true
			updated = true
			break
		}

	}

	if updated {
		fmt.Println("Updated the task")
	} else {
		fmt.Println("Task not found")
	}

}

func GetDataFromFile() *TodoList {
	jsonFile, err := os.Open("tasks.json")

	if err != nil {
		fmt.Println(err)
		return NewTodoList()
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		fmt.Println(err)
		return NewTodoList()
	}

	var result TodoList
	err = json.Unmarshal(byteValue, &result)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("\nGot length : %v.\n", len(result.Items))

	return &result
}

func WriteDataToFile(list *TodoList) {
	jsonString, _ := json.Marshal(list)

	os.WriteFile("tasks.json", jsonString, os.ModePerm)

}
