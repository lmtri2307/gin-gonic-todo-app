package task

import (
	"go-todo-app/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) getAll(userId int) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userId).Find(&tasks).Error
	return tasks, err
}

func (r *repository) getById(id int, userId int) (*Task, error) {
	// var task Task
	// if err := r.db.First(&task, id).Error; err != nil {
	// 	return nil, &Errors.NotFound
	// }

	// return &task, nil
	var task Task
	err := r.db.Where("id = ? AND user_id = ?", id, userId).First(&task).Error
	return &task, err
}

func (r *repository) save(task *Task) (*Task, error) {
	err := r.db.Save(task).Error

	return task, err
}

func (r *repository) deleteById(id int, userId int) error {
	err := r.db.Where("id = ? AND user_id = ?", id, userId).Delete(&Task{}).Error

	return err
}

func NewRepository() *repository {
	db := database.Init()
	repository := repository{db}
	db.AutoMigrate(&Task{})

	return &repository
}
