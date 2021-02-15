package middleware

import (
	"go-template/shared/response"
	"go-template/shared/utils"

	"github.com/labstack/echo/v4"
)

// VerifyJWT is a JWT middleware that verify the logged in user and set user context if verified.
// And will set user context to nil if not
func VerifyJWT(signingKey string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			token, err := utils.VerifyTokenFromRequest(c, signingKey)
			if err != nil {
				return next(c)
			}
			c.Set("user", token)

			return next(c)
		}
	}
}

// MustLoggedIn is a JWT middleware that verify the logged in user and set user context if verified.
// And will set user context if not
func MustLoggedIn() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// pass if already logged in
			if c.Get("user") != nil {
				return next(c)
			}
			// else return unauthorized
			return response.Unauthorized()

		}
	}
}
