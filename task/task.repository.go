package task

import (
	"errors"
	"go-todo-app/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var tasks = []Task{
	{ID: 1, Description: "Task 1"},
	{ID: 2, Description: "Task 2"},
	{ID: 3, Description: "Task 3"},
}

func (r *repository) getAll() ([]Task, error) {
	var tasks []Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, errors.New("internal error")
	}
	return tasks, nil
}

func (r *repository) getById(id int) (*Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (r *repository) create(payload CreateRequest) (*Task, error) {
	task := Task{
		Description: payload.Description,
	}

	if err := r.db.Create(&task).Error; err != nil {
		return nil, errors.New("internal error")
	}

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
	db := database.Init()
	repository := repository{db}
	db.AutoMigrate(&Task{})

	return &repository
}
