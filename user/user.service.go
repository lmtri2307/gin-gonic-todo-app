package user

type service struct {
	repository *repository
}

func (s *service) register(payload RegisterUserRequest) (*RegisterUserResponse, error) {
	user, err := NewUser(payload.UserName, payload.PassWord)

	if err != nil {
		return nil, err
	}

	user, err = s.repository.save(user)

	return &RegisterUserResponse{
		UserName: user.UserName,
	}, err
}

func NewService() *service {
	repository := NewRepository()
	service := service{repository}

	return &service
}
