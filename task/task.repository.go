package task

import (
	"errors"
	"go-todo-app/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
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

func (r *repository) save(task *Task) (*Task, error) {
	if err := r.db.Save(task).Error; err != nil {
		return nil, errors.New("internal error")
	}

	return task, nil
}

func (r *repository) deleteById(id int) error {
	if err := r.db.Delete(&Task{}, id).Error; err != nil {
		return errors.New("task not found")
	}

	return nil
}

func NewRepository() *repository {
	db := database.Init()
	repository := repository{db}
	db.AutoMigrate(&Task{})

	return &repository
}
