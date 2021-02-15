package user

import (
	"go-template/config"
	"go-template/middleware"
	"go-template/shared/pagination"
	"go-template/shared/response"
	"go-template/shared/utils"

	"github.com/labstack/echo/v4"
)

// RegisterModule registers a new update request module
func RegisterModule(r echo.Group, cfg config.Config, service IService) {
	handler := handler{cfg, service}

	r.Use(middleware.MustLoggedIn())

	r.GET("/users", handler.query)
	r.GET("/users/:user_id", handler.get)

}

type handler struct {
	cfg     config.Config
	service IService
}

func (h handler) query(c echo.Context) error {

	ctx := c.Request().Context()
	viewerID := utils.GetLoggedInUser(c).GetID()

	pages := pagination.NewFromRequest(c.Request())
	users, totalData, err := h.service.Query(ctx, "latest", viewerID, pages.Offset(), pages.Limit())
	if err != nil {
		return err
	}
	pages.SetData(users, totalData)

	return response.SuccessOK(c, pages)
}

func (h handler) get(c echo.Context) error {
	ctx := c.Request().Context()
	viewerID := utils.GetLoggedInUser(c).GetID()

	userID := c.Param("user_id")
	if userID == "me" {
		if viewerID == "" {
			return response.Unauthorized()
		}
		userID = viewerID
	}

	user, err := h.service.Get(ctx, userID, viewerID)
	if err != nil {
		return err
	}
	return response.SuccessOK(c, user)
}
