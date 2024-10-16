package task

type (
	CreateRequest struct {
		Description string `json:"description" binding:"required"`
	}
)
