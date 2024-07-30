package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TodoList struct {
	Tasks []Task `json:"tasks"`
}

const file = "tasks.json"

func main() {
	loadList := loadTodoList()
	fmt.Println("Welcome to the Todo list app!")

	printCommands()
	fmt.Println(loadList)
}

func printCommands() {
	commands := []string{
		"list",
		"add",
		"complete",
		"delete",
		"quit",
	}
	fmt.Println("here are the commands:", commands)
}

func loadTodoList() TodoList {
	var todoList TodoList
	data, err := os.ReadFile(file)
	if err != nil {
		return todoList
	}
	err = json.Unmarshal(data, &todoList)
	if err != nil {
		return TodoList{}
	}
	return todoList
}
