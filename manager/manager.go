package manager

import "github.com/MishaAC/go-todo-cli/task"

type Manager struct {
	tasks  []task.Task
	nextID int
}

func New() *Manager {
	return &Manager{
		tasks:  []task.Task{},
		nextID: 1,
	}
}

func (m *Manager) AddTask(title string) error {
	return nil
}

func (m *Manager) ListTasks() []task.Task {
	return nil
}

func (m *Manager) CompleteTask(id int) error {
	return nil
}

func (m *Manager) DeleteTask(id int) error {
	return nil
}

func (m *Manager) IsEmpty() bool {
	return false
}