package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var id = 1

type Tasks struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func getNextId(tasks []Tasks) int {
	maxId := 0
	for _, task := range tasks {
		if task.Id > maxId {
			maxId = task.Id
		}
	}
	return maxId + 1
}

func loadTasks() ([]Tasks, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasks []Tasks
	err = json.Unmarshal(data, &tasks)
	return tasks, nil
}

func AddTask(description string) error {
	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		os.Create("tasks.json")
		id = 0
	}

	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks %w", err)
	}

	task := Tasks{Id: getNextId(tasks),
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:   "",
	}

	tasks = append(tasks, task)

	err = MarshalJson(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal - %w", err)
	}
	fmt.Println()
	return nil
}

func ListTask() error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks %w", err)
	}
	for _, task := range tasks {
		fmt.Println(task)
	}
	fmt.Println()
	return nil
}

func DeleteTask(idT int) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks %w", err)
	}

	for i := range tasks {
		if tasks[i].Id == idT {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	err = MarshalJson(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal - %w", err)
	}

	fmt.Println()
	return nil
}

func UpdateTask(id int, description string) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed load file tasks - %w", err)
	}

	if id > len(tasks) {
		return fmt.Errorf("incorrect id - %w", err)
	}

	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
			break
		}
	}

	err = MarshalJson(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal - %w", err)
	}
	fmt.Println()
	return nil
}

func MarshalJson(tasks []Tasks) error {
	value, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal tasks: %w", err)
	}
	err = os.WriteFile("tasks.json", value, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}

func MarkProgressTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks - %w", err)
	}
	if id > len(tasks) {
		return fmt.Errorf("failed incorrect id - %w", err)
	}

	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Status = "in-progress"
			break
		}
	}
	err = MarshalJson(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal %w", err)
	}
	fmt.Println()
	return nil
}

func MarkDoneTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks - %w", err)
	}

	for i := range tasks {
		if tasks[i].Id == id {
			tasks[i].Status = "done"
			break
		}
	}
	err = MarshalJson(tasks)
	if err != nil {
		return fmt.Errorf("failed to marshal %w", err)
	}
	fmt.Println()
	return nil
}

func ListByStatus(status string) error {
	tasks, err := loadTasks()
	if err != nil {
		return fmt.Errorf("failed to load tasks - %w", err)
	}

	if status != "done" && status != "todo" && status != "in-progress" {
		return fmt.Errorf("incorrect input: `%s", status)
	}
	for i := range tasks {
		if tasks[i].Status == status {
			fmt.Println(tasks[i])
		}
	}
	fmt.Println()
	return nil
}
