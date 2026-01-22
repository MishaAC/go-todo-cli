package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/MishaAC/go-todo-cli/manager"
)

func main() {
	m := manager.New()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("======================")
		fmt.Println("       TODO CLI")
		fmt.Println("======================")
		fmt.Println("1) Add task")
		fmt.Println("2) List tasks")
		fmt.Println("3) Complete task")
		fmt.Println("4) Delete task")
		fmt.Println("5) Exit")
		fmt.Println()
		fmt.Print("Choose an option: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		input = strings.TrimSpace(input)
		userOption, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid option. Please enter a number between 1 and 5.")
			continue
		}

		switch userOption {
		case 1:
			fmt.Print("Enter task title: ")
			title, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading title:", err)
				continue
			}

			title = strings.TrimSpace(title)
			if err := m.AddTask(title); err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Task added successfully!")

		case 2:
			if m.IsEmpty() {
				fmt.Println("No tasks available")
				continue
			}

			fmt.Println("\nYour tasks:")
			for _, t := range m.ListTasks() {
				status := " "
				if t.Completed {
					status = "x"
				}
				fmt.Printf("[%s] %d - %s\n", status, t.ID, t.Title)
			}

		case 3:
			fmt.Print("Enter task ID to complete: ")
			idInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading ID:", err)
				continue
			}

			idInput = strings.TrimSpace(idInput)
			id, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}

			if err := m.CompleteTask(id); err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Task marked as completed!")

		case 4:
			fmt.Print("Enter task ID to delete: ")
			idInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading ID:", err)
				continue
			}

			idInput = strings.TrimSpace(idInput)
			id, err := strconv.Atoi(idInput)
			if err != nil {
				fmt.Println("Invalid ID. Please enter a number.")
				continue
			}

			if err := m.DeleteTask(id); err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("Task deleted successfully!")

		case 5:
			fmt.Println("Goodbye! 👋")
			return

		default:
			fmt.Println("Option must be between 1 and 5.")
		}
	}
}
