package console

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	task "taskTracker/task"
)

func meeting() {
	fmt.Println("Добро пожаловать в TaskTracker!")
	fmt.Println("Если вам нужна помощь введите команду \"help\"")
}

func help() {
	fmt.Println("command 'exit' - exit programm")
	fmt.Println("command 'add' \"something\" - add task")
	fmt.Println("command 'update' \"1 something and anything\" - update task")
	fmt.Println("command 'delete' \"1\" - delete task")
	fmt.Println("command 'mark-in-progress' \"1\" - mark task in progress")
	fmt.Println("command 'mark-done' \"1\" - mark task done")
	fmt.Println("command 'list' - list all tasks")
	fmt.Println("command 'list done' - list done tasks")
	fmt.Println("command 'list todo' - list todo tasks")
	fmt.Println("command 'list in-progress' - list tasks in progress")
	fmt.Println()
}

func exit() {
	os.Exit(0)
}

func parseCommand(input string) (string, []string) {
	commands := strings.Fields(input)
	if len(commands) == 0 {
		return "", nil
	}
	return commands[0], commands[1:]
}

func validateArgsCount(args []string, expected int) error {
	if len(args) != expected {
		return fmt.Errorf("expected %d arguments, got %d", expected, len(args))
	}
	return nil
}

func validateArgsMin(args []string, min int) error {
	if len(args) < min {
		return fmt.Errorf("expected at least %d arguments, got %d", min, len(args))
	}
	return nil
}

func parseID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid task ID: %s", idStr)
	}
	if id <= 0 {
		return 0, fmt.Errorf("task ID must be positive: %d", id)
	}
	return id, nil
}

func switchCommand(command string) error {
	cmd, args := parseCommand(command)

	switch cmd {
	case "help":
		help()
	case "list":
		if len(args) == 0 {
			err := task.ListTask()
			if err != nil {
				return fmt.Errorf("failed to list tasks: %w", err)
			}
		} else {
			status := args[0]
			if status == "done" || status == "todo" || status == "in-progress" {
				task.ListByStatus(status)
			} else {
				return fmt.Errorf("unknown status for list command: %s. Use 'done', 'todo' or 'in-progress'", status)
			}
		}
	case "add":
		if err := validateArgsMin(args, 1); err != nil {
			return fmt.Errorf("add command requires task description: %w", err)
		}
		description := strings.Join(args, " ")
		err := task.AddTask(description)
		if err != nil {
			return fmt.Errorf("failed add task - %w", err)
		}
		fmt.Printf("Task add successfully\n")
	case "delete":
		if err := validateArgsCount(args, 1); err != nil {
			return fmt.Errorf("delete command requires exactly one argument (task ID): %w", err)
		}
		id, err := parseID(args[0])
		if err != nil {
			return fmt.Errorf("delete command: %w", err)
		}
		err = task.DeleteTask(id)
		if err != nil {
			return err
		}
		fmt.Printf("Task %d deleted successfully\n", id)
	case "exit":
		exit()
	case "update":
		if err := validateArgsMin(args, 2); err != nil {
			return fmt.Errorf("update command requires at least 2 arguments (ID and description): %w", err)
		}
		id, err := parseID(args[0])
		if err != nil {
			return fmt.Errorf("update command: %w", err)
		}

		description := strings.Join(args[1:], " ")

		err = task.UpdateTask(id, description)
		if err != nil {
			return fmt.Errorf("failed to update task - %w", err)
		}
		fmt.Printf("Task %d updated successfully\n", id)
	case "mark-in-progress":
		if err := validateArgsCount(args, 1); err != nil {
			return fmt.Errorf("mark-in-progress command requires exactly one argument (task ID): %w", err)
		}
		id, err := parseID(args[0])
		if err != nil {
			return fmt.Errorf("mark-in-progress command: %w", err)
		}

		err = task.MarkTask(id, "in-progress")
		if err != nil {
			return fmt.Errorf("failed mark in-progress - %w", err)
		}
		fmt.Printf("Task %d marked as in-progress\n", id)
	case "mark-done":
		if err := validateArgsCount(args, 1); err != nil {
			return fmt.Errorf("mark-done command requires exactly one argument (task ID): %w", err)
		}
		id, err := parseID(args[0])
		if err != nil {
			return fmt.Errorf("mark-done command: %w", err)
		}
		err = task.MarkTask(id, "done")
		if err != nil {
			return fmt.Errorf("failed mark done - %w", err)
		}
		fmt.Printf("Task %d marked as done\n", id)
	default:
		fmt.Println("wrong input, try command \"help\"")
		fmt.Println()
	}
	return nil
}

func userInput() error {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		err := switchCommand(scanner.Text())
		if err != nil {
			return fmt.Errorf("failed switch command - %w", err)
		}
	}
	return nil
}

func Start() {
	meeting()
	for {
		err := userInput()
		if err != nil {
			fmt.Println("failed  to user input - %w", err)
		}
	}
}
