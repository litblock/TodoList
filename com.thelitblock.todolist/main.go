package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	fmt.Println(loadList)
	for {
		printCommands()
		choice := getUserInput("Enter your choice: ")
		switch choice {
		case "list":
			listTasks(loadList)
		case "add":
			addTask(&loadList)
		case "complete":
			fmt.Println("You chose complete")
		case "delete":
			fmt.Println("You chose delete")
		case "quit":
			fmt.Println("You chose quit")
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
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

func addTask(todoList *TodoList) {
	title := getUserInput("Enter task title: ")
	newTask := Task{
		ID:        len(todoList.Tasks) + 1,
		Title:     title,
		Completed: false,
	}
	todoList.Tasks = append(todoList.Tasks, newTask)
	fmt.Println("Task added successfully!")
}

func listTasks(todoList TodoList) {
	if len(todoList.Tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, task := range todoList.Tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", task.ID, status, task.Title)
	}
}
