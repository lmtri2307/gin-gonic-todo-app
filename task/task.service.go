package task

type service struct {
	repository *repository
}

func (*service) HelloWorld() string {
	return "Hello world"
}

func (s *service) GetAll() []Task {
	return s.repository.getAll()
}

func NewService() *service {
	repository := NewRepository()
	service := service{repository}

	return &service
}
