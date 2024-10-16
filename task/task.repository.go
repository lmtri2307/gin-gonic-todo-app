package task

import "errors"

type repository struct {
}

var tasks = []Task{
	{ID: 1, Description: "Task 1"},
	{ID: 2, Description: "Task 2"},
	{ID: 3, Description: "Task 3"},
}

func (*repository) getAll() []Task {
	return tasks
}

func (*repository) getById(id int) (*Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, errors.New("Task not found")
}

func NewRepository() *repository {
	repository := repository{}
	return &repository
}
