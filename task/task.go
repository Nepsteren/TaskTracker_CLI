package task

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

var id = 1

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func loadTasks() []Task {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks
}

func CreateTask(description string) {
	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		os.Create("tasks.json")
		id = 0
	}

	tasks := loadTasks()
	id = len(tasks) + 1

	task := Task{Id: id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   "",
	}

	tasks = append(tasks, task)

	_, err = json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}

	value, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("tasks.json", value, 0644)
	fmt.Println(tasks)

}
