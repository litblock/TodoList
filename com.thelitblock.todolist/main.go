package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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
			markTaskComplete(&loadList)
		case "delete":
			deleteTask(&loadList)
		case "save":
			saveTodoList(loadList)
		case "quit":
			os.Exit(0)
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
		"save",
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

func deleteTask(todoList *TodoList) {
	listTasks(*todoList)
	idStr := getUserInput("Enter the ID of the task to delete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}
	for i, task := range todoList.Tasks {
		if task.ID == id {
			todoList.Tasks = append(todoList.Tasks[:i], todoList.Tasks[i+1:]...)
			fmt.Println("Task deleted!")
			return
		}
	}
	fmt.Println("Task not found.")
}

func saveTodoList(todoList TodoList) {
	data, _ := json.MarshalIndent(todoList, "", "  ")
	err := os.WriteFile(file, data, 0644)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
	}
}

func markTaskComplete(todoList *TodoList) {
	listTasks(*todoList)
	idStr := getUserInput("Enter the ID of the task to mark as complete: ")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}
	for i, task := range todoList.Tasks {
		if task.ID == id {
			todoList.Tasks[i].Completed = true
			fmt.Println("Task marked as complete!")
			return
		}
	}
	fmt.Println("Task not found.")
}
