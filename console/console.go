package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func meeting() {
	fmt.Println("Добро пожаловать в TaskTracker!")
	fmt.Println("Если вам нужна помощь введите команду \"help\"")
}

func help() {
	fmt.Println("command 'add' \"something\" - add task")
	fmt.Println("command 'update' \"1 something and anything\" - update task")
	fmt.Println("command 'delete' \"1\" - delete task")
	fmt.Println("command 'mark-in-progress' \"1\" - mark task in progress")
	fmt.Println("command 'mark-done' \"1\" - mark task done")
	fmt.Println("command 'list done' - list done task")
	fmt.Println("command 'list todo' - list todo task")
	fmt.Println("command 'list in-progress' - list task in progress")
	fmt.Println()
}

func swithCommand(command string) {
	commands := strings.Fields(command)
	command = commands[0]
	switch command {
	case "help":
		help()
	case "add":
		fmt.Println("some add")
		fmt.Println()
	case "delete":
		fmt.Println("some delete")
		fmt.Println()
	case "update":
		fmt.Println("update task")
		fmt.Println()
	case "list done":
		fmt.Println("some done list")
		fmt.Println()
	case "list not done":
		fmt.Println("some not done list")
		fmt.Println()
	case "list in progress":
		fmt.Println("some progress list")
		fmt.Println()
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
