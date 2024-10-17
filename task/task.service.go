package task

import "errors"

type service struct {
	repository *repository
}

func (*service) HelloWorld() string {
	return "Hello world"
}

func (s *service) GetAll() ([]Task, error) {
	return s.repository.getAll()
}

func (s *service) GetById(id int) (*Task, error) {
	return s.repository.getById(id)
}

func (s *service) Create(payload CreateRequest) (*Task, error) {
	task, err := s.repository.save(&Task{Description: payload.Description})
	return task, err
}

func (s *service) UpdateById(id int, payload UpdateRequest) (*Task, error) {
	task, err := s.repository.getById(id)
	if err != nil {
		return nil, err
	}
	task.Description = payload.Description
	task, err = s.repository.save(task)
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (s *service) DeleteById(id int) error {
	_, err := s.repository.getById(id)
	if err != nil {
		return errors.New("task not found")
	}

	return s.repository.deleteById(id)
}

func NewService() *service {
	repository := NewRepository()
	service := service{repository}

	return &service
}
