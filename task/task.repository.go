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

func (*repository) saveNew(payload CreateRequest) (*Task, error) {
	task := Task{len(tasks) + 1, payload.Description}
	tasks = append(tasks, task)
	return &task, nil
}

func (r *repository) save(task *Task) (*Task, error) {
	for index, oldTask := range tasks {
		if oldTask.ID == task.ID {
			tasks[index].Description = task.Description
			return &tasks[index], nil
		}
	}
	return nil, errors.New("task not found")
}

func (r *repository) deleteById(id int) error {
	for index, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			return nil
		}
	}

	return errors.New("Task Not Found")
}

func NewRepository() *repository {
	repository := repository{}
	return &repository
}
