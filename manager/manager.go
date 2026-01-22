package manager

import "github.com/MishaAC/go-todo-cli/task"

type Manager struct {
	tasks  []task.Task
	nextID int
}