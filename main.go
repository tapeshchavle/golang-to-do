package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"go_project1/todo"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	// Load existing tasks
	taskList, err := todo.Load()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		os.Exit(1)
	}

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task description")
			fmt.Println("Usage: go run . add \"Your task description\"")
			return
		}
		description := strings.Join(os.Args[2:], " ")
		task := taskList.Add(description)
		if err := todo.Save(taskList); err != nil {
			fmt.Println("Error saving task:", err)
			return
		}
		fmt.Printf("✅ Added task #%d: %s\n", task.ID, task.Description)

	case "list":
		if len(taskList.Tasks) == 0 {
			fmt.Println("📋 No tasks yet! Add one with: go run . add \"Your task\"")
			return
		}
		fmt.Println("\n📋 Your To-Do List:")
		fmt.Println(strings.Repeat("-", 50))
		for _, task := range taskList.Tasks {
			status := "[ ]"
			if task.Completed {
				status = "[✓]"
			}
			fmt.Printf("%s #%d: %s\n", status, task.ID, task.Description)
		}
		fmt.Println(strings.Repeat("-", 50))

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID")
			fmt.Println("Usage: go run . complete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}
		if taskList.Complete(id) {
			if err := todo.Save(taskList); err != nil {
				fmt.Println("Error saving:", err)
				return
			}
			fmt.Printf("🎉 Completed task #%d\n", id)
		} else {
			fmt.Printf("❌ Task #%d not found\n", id)
		}

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task ID")
			fmt.Println("Usage: go run . delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Error: Invalid task ID")
			return
		}
		if taskList.Delete(id) {
			if err := todo.Save(taskList); err != nil {
				fmt.Println("Error saving:", err)
				return
			}
			fmt.Printf("🗑️  Deleted task #%d\n", id)
		} else {
			fmt.Printf("❌ Task #%d not found\n", id)
		}

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println(`
📝 CLI To-Do List

Usage:
  go run . <command> [arguments]

Commands:
  add <task>      Add a new task
  list            Show all tasks
  complete <id>   Mark a task as completed
  delete <id>     Delete a task

Examples:
  go run . add "Learn Go basics"
  go run . list
  go run . complete 1
  go run . delete 2
`)
}
