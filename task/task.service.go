package task

type service struct {
	repository *repository
}

func (*service) HelloWorld() string {
	return "Hello world"
}

func (s *service) GetAll(userId int) ([]Task, error) {
	return s.repository.getAll(userId)
}

func (s *service) GetById(id int, userId int) (*Task, error) {
	task, err := s.repository.getById(id, userId)
	if err != nil {
		return nil, &Errors.NotFound
	}
	return task, nil
}

func (s *service) Create(payload CreateRequest, userId int) (*Task, error) {
	task, err := s.repository.save(&Task{
		Description: payload.Description,
		UserID:      userId,
	})
	return task, err
}

func (s *service) UpdateById(id int, payload UpdateRequest, userId int) (*Task, error) {
	task, err := s.repository.getById(id, userId)
	if err != nil {
		return nil, &Errors.NotFound
	}
	task.Description = payload.Description
	return s.repository.save(task)
}

func (s *service) DeleteById(id int, userId int) error {
	_, err := s.repository.getById(id, userId)
	if err != nil {
		return &Errors.NotFound
	}

	return s.repository.deleteById(id, userId)
}

func NewService() *service {
	repository := NewRepository()
	service := service{repository}

	return &service
}
