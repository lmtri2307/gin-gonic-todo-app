package auth

import (
	"go-todo-app/base"
	"net/http"
)

var Errors = struct {
	IncorrectPassword   base.ApiError
	InvalidLoginRequest base.ApiError
}{
	IncorrectPassword: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "wrong password",
	},
	InvalidLoginRequest: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "invalid login request",
	},
}
