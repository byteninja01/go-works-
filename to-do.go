package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const taskFile = "tasks.txt"

func loadTasks() []string {
	file, err := os.Open(taskFile)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}
	return tasks
}

func saveTasks(tasks []string) {
	file, err := os.Create(taskFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	for _, task := range tasks {
		file.WriteString(task + "\n")
	}
}

func main() {
	tasks := loadTasks()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTo-Do List App")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")

		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)

		switch choiceInput {
		case "1":
			fmt.Print("Enter task: ")
			task, _ := reader.ReadString('\n')
			tasks = append(tasks, strings.TrimSpace(task))
			saveTasks(tasks)
			fmt.Println("Task added!")

		case "2":
			fmt.Println("Your Tasks:")
			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task)
			}

		case "3":
			fmt.Print("Enter task number to delete: ")
			var index int
			fmt.Scanln(&index)
			if index > 0 && index <= len(tasks) {
				tasks = append(tasks[:index-1], tasks[index:]...)
				saveTasks(tasks)
				fmt.Println("Task deleted!")
			} else {
				fmt.Println("Invalid task number!")
			}

		case "4":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}
