package task

type repository struct {
}

var tasks = []Task{
	{ID: 1, Description: "Task 1"},
	{ID: 2, Description: "Task 2"},
	{ID: 3, Description: "Task 3"},
}

func (*repository) getAll() []Task {
	return tasks
}

func NewRepository() *repository {
	repository := repository{}
	return &repository
}
