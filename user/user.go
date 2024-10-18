package user

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserName       string `gorm:"unique" json:"username"`
	HashedPassword string
}

func NewUser(username string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	return &User{UserName: username, HashedPassword: string(hashedPassword)}, nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(providedPassword))
	return err
}
