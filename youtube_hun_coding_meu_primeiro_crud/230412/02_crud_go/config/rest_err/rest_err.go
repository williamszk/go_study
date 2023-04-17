package rest_err

import "net/http"

type RestErr struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// To make a Go friendly string error using our RestErr.
// The Error method is a method from an interface.
func (r *RestErr) Error() string {
	return r.Message
}

func NewRestErr(message string, err string, code int, causes []Cause) RestErr {
	return RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Cause) RestErr {
	return RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) RestErr {
	return RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) RestErr {
	return RestErr{
		Message: message,
		Err:     "not found error",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) RestErr {
	return RestErr{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
