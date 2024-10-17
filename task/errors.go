package task

import (
	"go-todo-app/base"
	"net/http"
)

var Errors = struct {
	NotFound             base.ApiError
	InvalidId            base.ApiError
	InvalidUpdatePayload base.ApiError
}{
	NotFound: base.ApiError{
		Status:  http.StatusNotFound,
		Message: "task not found",
	},
	InvalidId: base.ApiError{
		Status:  http.StatusBadRequest,
		Message: "invalid task id",
	},
	InvalidUpdatePayload: base.ApiError{
		Status:  http.StatusBadGateway,
		Message: "invalid update task payload",
	},
}
