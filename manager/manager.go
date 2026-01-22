package manager

import (
	"errors"

	"github.com/MishaAC/go-todo-cli/task"
)

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
	if title == "" {
		return errors.New("task title cannot be empty")
	}
	task := task.Task{
		ID:        m.nextID,
		Title:     title,
		Completed: false,
	}
	m.tasks = append(m.tasks, task)
	m.nextID++
	return nil
}

func (m *Manager) ListTasks() []task.Task {
	tasksCopy := make([]task.Task, len(m.tasks))
	copy(tasksCopy, m.tasks)
	return tasksCopy
}

func (m *Manager) CompleteTask(id int) error {
	return nil
}

func (m *Manager) DeleteTask(id int) error {
	return nil
}

func (m *Manager) IsEmpty() bool {
	return len(m.tasks) == 0
}
