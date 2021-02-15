package utils

import (
	"go-template/internal/entity"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// GetLoggedInUser returns logged in user interface
func GetLoggedInUser(c echo.Context) entity.User {
	if c.Get("user") == nil {
		return entity.User{}
	}
	LoggedInUser := c.Get("user").(*jwt.Token)

	claims := LoggedInUser.Claims.(jwt.MapClaims)

	var id string
	if val, ok := claims["id"].(string); ok {
		id = val
	}

	var username string
	if val, ok := claims["username"].(string); ok {
		username = val
	}

	return entity.User{
		ID:    id,
		Email: username,
	}
}
