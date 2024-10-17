package auth

import (
	"go-todo-app/user"
	"go-todo-app/utils/tokenutil"
)

type service struct {
	userService *user.Service
}

func (s *service) login(request LoginRequest) (*LoginResponse, error) {
	user, err := s.userService.FindByUsername(request.UserName)
	if err != nil {
		return nil, err
	}

	err = user.CheckPassword(request.PassWord)
	if err != nil {
		return nil, &Errors.IncorrectPassword
	}

	accessToken, err := tokenutil.CreateToken(TokenPayload{user.ID}, "secret", 5000)
	return &LoginResponse{AccessToken: accessToken}, err
}

func NewService() *service {
	userService := user.NewService()
	service := service{userService}

	return &service
}
