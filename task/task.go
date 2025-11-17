package task

type Task struct {
	id          int
	description string
	status      string //todo, in-progress, done
	createdAt   string
	updatedAt   string
}

func CreateTask() Task {
	firstTask := Task{
		id:          1,
		description: "to do something",
		status:      "todo",
		createdAt:   "0000.00.00",
		updatedAt:   "0000.00.00",
	}
	return firstTask
}
