package task

type service struct {
}

func (*service) HelloWorld() string {
	return "Hello world"
}

func newService() *service {
	service := service{}

	return &service
}
