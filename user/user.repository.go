package user

import (
	"go-todo-app/database"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r *repository) save(user *User) (*User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func NewRepository() *repository {
	db := database.Init()
	repository := repository{db}
	db.AutoMigrate(&User{})

	return &repository
}
