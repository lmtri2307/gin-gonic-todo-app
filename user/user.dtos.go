package user

type (
	RegisterUserRequest struct {
		UserName string `json:"userName" binding:"required"`
		PassWord string `json:"passWord" binding:"required"`
	}
	RegisterUserResponse struct {
		UserName string `json:"userName"`
	}
)
