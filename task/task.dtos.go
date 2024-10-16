package task

type (
	CreateRequest struct {
		Description string `json:"description" binding:"required"`
	}
	UpdateRequest struct {
		Description string `json:"description" binding:"required"`
	}
)
