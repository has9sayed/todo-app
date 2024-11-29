package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"github.com/has9sayed/todo-app/todo"
)

func main() {
	taskList := todo.NewTaskList()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			taskList.AddTask(strings.TrimSpace(title))
			fmt.Println("Task added!")

		case "2":
			taskList.ListTasks()

		case "3":
			fmt.Print("Please enter task ID to mark as done: ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))
			if err != nil {
				fmt.Println("Invalid ID. Try again.")
				continue
			}
			taskList.MarkDone(id)

		case "4":
			fmt.Print("Enter task ID to delete: ")
			idInput, _ := reader.ReadString('\n')
			id, err := strconv.Atoi(strings.TrimSpace(idInput))
			if err != nil {
				fmt.Println("Invalid ID. Try again.")
				continue
			}
			taskList.DeleteTask(id)

		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}