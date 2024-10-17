package task

import (
	"go-todo-app/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) getAll() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *repository) getById(id int) (*Task, error) {
	var task Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, &Errors.NotFound
	}

	return &task, nil
}

func (r *repository) save(task *Task) (*Task, error) {
	err := r.db.Save(task).Error

	return task, err
}

func (r *repository) deleteById(id int) error {
	err := r.db.Delete(&Task{}, id).Error

	return err
}

func NewRepository() *repository {
	db := database.Init()
	repository := repository{db}
	db.AutoMigrate(&Task{})

	return &repository
}
