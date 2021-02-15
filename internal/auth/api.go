package auth

import (
	"go-template/config"
	"go-template/shared/response"

	"github.com/labstack/echo/v4"
)

// RegisterModule registers a new auth module
func RegisterModule(r echo.Group, cfg config.Config, service IService) {
	handler := handler{cfg, service}

	r.POST("/auth/token/refresh", handler.refreshToken)
	r.POST("/auth/login", handler.login)

}

type handler struct {
	cfg     config.Config
	service IService
}

func (h handler) login(c echo.Context) error {
	req := LoginRequest{}
	err := c.Bind(&req)
	if err != nil {
		return response.BadRequest()
	}

	res, err := h.service.Login(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return response.SuccessOK(c, res)
}

func (h handler) refreshToken(c echo.Context) error {
	var req RefreshTokenRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	res, err := h.service.RefreshToken(c.Request().Context(), req)
	if err != nil {
		return err
	}

	return response.SuccessOK(c, res, "token refreshed")
}
