package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

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

func withTaskById(id int, operation func(task *Tasks) error) error {
	return withTask(func(tasks *[]Tasks) error {
		for i := range *tasks {
			if (*tasks)[i].Id == id {
				return operation(&(*tasks)[i])
			}
		}
		return fmt.Errorf("task with ID %d not found", id)
	})
}

func withTask(operation func(task *[]Tasks) error) error {
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		if err := os.WriteFile("tasks.json", []byte("[]"), 0644); err != nil {
			return fmt.Errorf("failed to create tasks file: %w", err)
		}
	}
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if err := operation(&tasks); err != nil {
		return err
	}

	return MarshalJson(tasks)
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

func loadTasks() ([]Tasks, error) {
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var tasks []Tasks
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal - %w", err)
	}
	return tasks, nil
}

func AddTask(description string) error {
	if description == "" {
		return fmt.Errorf("task description cannot be empty")
	}
	return withTask(func(tasks *[]Tasks) error {
		task := Tasks{Id: getNextId(*tasks),
			Description: description,
			Status:      "todo",
			CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt:   "",
		}
		*tasks = append(*tasks, task)
		return nil
	})
}

func ListTask() error {
	return withTask(func(tasks *[]Tasks) error {
		for _, task := range *tasks {
			fmt.Println(task)
		}
		return nil
	})
}

func DeleteTask(id int) error {
	return withTask(func(tasks *[]Tasks) error {
		for i := 0; i < len(*tasks); i++ {
			if (*tasks)[i].Id == id {
				*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("task with ID %d not found", id)
	})
}

func UpdateTask(id int, description string) error {
	return withTaskById(id, func(task *Tasks) error {
		task.Description = description
		task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		return nil
	})
}

func MarkTask(id int, status string) error {
	return withTaskById(id, func(task *Tasks) error {
		task.Status = status
		task.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")
		return nil
	})
}

func ListByStatus(status string) error {
	return withTask(func(tasks *[]Tasks) error {
		for _, task := range *tasks {
			if task.Status == status {
				fmt.Println(task)
			}
		}
		return nil
	})
}
