package response

import "net/http"

// ErrorResponse is the response that represents an error.
type ErrorResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"-"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// Error is required by the error interface.
func (e ErrorResponse) Error() string {
	return e.Message
}

// StatusCode is required by CustomHTTPErrorHandler
func (e ErrorResponse) StatusCode() int {
	return e.Code
}

// InternalServerError creates a new error response representing an internal server error (HTTP 500)
func InternalServerError(msg ...string) ErrorResponse {
	responseMsg := buildResponseMsg("We encountered an error while processing your request.", msg...)
	return ErrorResponse{
		Status:  http.StatusText(http.StatusInternalServerError),
		Code:    http.StatusInternalServerError,
		Message: responseMsg,
	}
}

// NotFound creates a new error response representing a resource-not-found error (HTTP 404)
func NotFound(msg ...string) ErrorResponse {
	responseMsg := buildResponseMsg("The requested resource was not found.", msg...)
	return ErrorResponse{
		Status:  http.StatusText(http.StatusNotFound),
		Code:    http.StatusNotFound,
		Message: responseMsg,
	}
}

// Unauthorized creates a new error response representing an authentication/authorization failure (HTTP 401)
func Unauthorized(msg ...string) ErrorResponse {
	responseMsg := buildResponseMsg("You are not authorized to perform the requested action.", msg...)
	return ErrorResponse{
		Status:  http.StatusText(http.StatusUnauthorized),
		Code:    http.StatusUnauthorized,
		Message: responseMsg,
	}
}

// Forbidden creates a new error response representing an authorization failure (HTTP 403)
func Forbidden(msg ...string) ErrorResponse {
	responseMsg := buildResponseMsg("You are not authorized to perform the requested action.", msg...)
	return ErrorResponse{
		Status:  http.StatusText(http.StatusForbidden),
		Code:    http.StatusForbidden,
		Message: responseMsg,
	}
}

// BadRequest creates a new error response representing a bad request (HTTP 400)
func BadRequest(msg ...string) ErrorResponse {
	responseMsg := buildResponseMsg("Your request is in a bad format.", msg...)
	return ErrorResponse{
		Status:  http.StatusText(http.StatusBadRequest),
		Code:    http.StatusBadRequest,
		Message: responseMsg,
	}
}
