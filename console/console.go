package console

import (
	"bufio"
	"fmt"
	"log"
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

func swithCommand(command string) {
	commands := strings.Fields(command)
	command = commands[0]
	description := ""
	if len(commands) >= 2 {
		description = commands[1]
	}
	if len(commands) >= 3 {
		description = strings.Join(commands[2:], " ")
	}

	if commands[0] == "list" && len(commands) >= 2 {
		status := commands[1]
		if status == "done" || status == "todo" || status == "in-progress" {
			task.ListByStatus(status)
			return
		}
	}

	switch command {
	case "help":
		help()
	case "list":
		err := task.ListTask()
		if err != nil {
			log.Fatal(err)
		}
	case "add":
		err := task.AddTask(description)
		if err != nil {
			log.Fatal(err)
		}
	case "delete":
		if len(commands) > 2 {
			log.Fatal(fmt.Errorf("failed too much argumentsS"))
		}
		id, err := strconv.Atoi(commands[1])
		if err != nil {
			log.Fatal(fmt.Errorf("failed incorrect id"))
		}
		err = task.DeleteTask(id)
		if err != nil {
			log.Fatal(err)
		}
	case "exit":
		exit()
	case "update":
		id, err := strconv.Atoi(commands[1])
		if err != nil {
			log.Fatal(fmt.Errorf("failed incorrect input - %w", err))
		}
		task.UpdateTask(id, description)
	case "mark-in-progress":
		id, err := strconv.Atoi(commands[1])
		if err != nil {
			log.Fatal(fmt.Errorf("failed incorrect id"))
		}
		err = task.MarkProgressTask(id)
		if err != nil {
			log.Fatal(fmt.Errorf("failed makr done - %w", err))
		}
	case "mark-done":
		id, err := strconv.Atoi(commands[1])
		if err != nil {
			log.Fatal(fmt.Errorf("failed incorrect id"))
		}
		err = task.MarkDoneTask(id)
		if err != nil {
			log.Fatal(fmt.Errorf("failed mark done - %w", err))
		}
	// case "list done":
	// 	err := task.ListByStatus(description)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// case "list todo":
	// 	err := task.ListByStatus(description)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// case "list in-progress":
	// 	err := task.ListByStatus(description)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	default:
		fmt.Println("wrong input, try command \"help\"")
		fmt.Println()
	}

}

func usetInput() {
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		swithCommand(scanner.Text())
	}
}

func Start() {
	meeting()
	for {
		usetInput()
	}
}
