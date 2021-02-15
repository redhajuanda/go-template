package app

import (
	"database/sql"
	"fmt"
	"go-template/config"
	"go-template/infra/logger"
	"go-template/shared/response"
	"go-template/shared/validator"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

// App Errors
var (
	ErrNoRowsAffected = fmt.Errorf("No Rows Affected")
)

// CustomHTTPErrorHandler sets error response for different type of errors and logs
func CustomHTTPErrorHandler(cfg config.Config) echo.HTTPErrorHandler {

	return func(err error, c echo.Context) {
		logEntry := makeLogEntry(c)

		if resp, ok := err.(stackTracer); ok {
			if cfg.Server.ENV == "production" {
				logEntry = logEntry.WithParam("trace", resp.StackTrace())
			} else {
				fmt.Printf("%+v\n", resp.StackTrace())
			}
		}

		err = errors.Cause(err)
		code := http.StatusInternalServerError

		if res, ok := err.(*mysql.MySQLError); ok {
			switch res.Number {
			case 1452: // 1452: Cannot add or update a child row: a foreign key constraint fails
				err = response.NotFound()
			}
		}

		if errors.Is(err, pg.ErrNoRows) {
			err = response.NotFound()
		}

		if val, ok := err.(pg.Error); ok {
			if val.Field('C') == "22P02" {
				err = response.BadRequest("The given ID is not in the right format")
			}
		}
		// Handles bad request error
		if _, ok := err.(validator.ValidationErrors); ok {
			err = response.BadRequest(err.Error())
		}

		if errors.Is(err, sql.ErrNoRows) {
			err = response.NotFound()
		}

		if resp, ok := err.(response.ErrorResponse); ok {
			code = resp.StatusCode()
		}

		if resp, ok := err.(*echo.HTTPError); ok {
			code = resp.Code
			err = response.ErrorResponse{
				Status:  http.StatusText(code),
				Code:    code,
				Message: resp.Message.(string),
			}
		}

		logEntry.WithParam("code", code)

		if code == http.StatusInternalServerError {
			// bugsnag.Notify(err)
			logEntry.Error(err)
			c.JSON(code, response.InternalServerError())
		} else {
			logEntry.Info(err)
			c.JSON(code, response.ErrorResponse{
				Status:  http.StatusText(code),
				Code:    code,
				Message: err.Error(),
			})
		}
	}
}

func makeLogEntry(c echo.Context) logger.LoggerWrapper {
	if c == nil {
		return logger.WithParams(logger.Params{})
	}

	id := c.Request().Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = c.Response().Header().Get(echo.HeaderXRequestID)
	}

	return logger.WithParams(logger.Params{
		// "at":     time.Now().Format("2006-01-02 15:04:05"),
		"method": c.Request().Method,
		"uri":    c.Request().URL.String(),
		"user":   c.Get("user"),
		"ip":     c.Request().RemoteAddr,
		"id":     id,
	})
}
