package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response struct
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Success responses with JSON formatresponseMsg
func Success(c echo.Context, code int, data interface{}, msg ...string) error {

	responseMsg := buildResponseMsg("Success", msg...)

	res := Response{
		Status:  http.StatusText(code),
		Message: responseMsg,
		Data:    data,
	}
	return c.JSON(code, res)
}

// SuccessOK returns code 200
func SuccessOK(c echo.Context, data interface{}, msg ...string) error {
	return Success(c, http.StatusOK, data, msg...)
}

// SuccessCreated returns code 201
func SuccessCreated(c echo.Context, data interface{}, msg ...string) error {
	return Success(c, http.StatusCreated, data, msg...)
}
