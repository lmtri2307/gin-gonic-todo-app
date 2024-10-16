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

func (s *service) GetById(id int) (*Task, error) {
	task, err := s.repository.getById(id)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *service) Create(payload CreateRequest) (*Task, error) {
	task, err := s.repository.saveNew(payload)
	return task, err
}

func NewService() *service {
	repository := NewRepository()
	service := service{repository}

	return &service
}
