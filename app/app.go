package app

import (
	"go-template/config"
	"go-template/internal/auth"
	"go-template/internal/user"
	customMiddleware "go-template/middleware"
	"net/http"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// App struct
type App struct {
	cfg    config.Config
	router *echo.Echo
	db     *pg.DB
	// storage storage.Client
}

// New inits a new app
func New(
	cfg config.Config,
	router *echo.Echo,
	db *pg.DB,
	// storage storage.Client,
) *App {
	return &App{
		cfg,
		router,
		db,
		// storage,
	}
}

// BuildHandler builds handlers and returns router
func (app App) BuildHandler() *echo.Echo {

	app.configRouter()

	auth.RegisterModule(
		*app.router.Group(""),
		app.cfg,
		auth.NewService(
			app.cfg,
			user.NewRepository(app.db),
		),
	)

	user.RegisterModule(
		*app.router.Group(""),
		app.cfg,
		user.NewService(
			app.cfg,
			user.NewRepository(app.db),
		),
	)

	app.router.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"app": app.cfg.Server.NAME,
			"env": app.cfg.Server.ENV,
		})
	})

	return app.router
}

func (app App) configRouter() {

	app.router.Pre(middleware.RemoveTrailingSlash())
	app.router.Use(middleware.RequestID())

	// Setup custom HTTP error handler
	app.router.HTTPErrorHandler = CustomHTTPErrorHandler(app.cfg)

	// Register middleware recover from panic
	// app.router.Use(customMiddleware.Recovery)

	// Verify JWT
	app.router.Use(customMiddleware.VerifyJWT(app.cfg.JWT.SigningKey))

	// Setup access log
	// app.router.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// 	Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}","host":"${host}",` +
	// 		`"method":"${method}","uri":"${uri}","status":${status},"error":"${error}","latency":${latency},` +
	// 		`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
	// 		`"service":"` + app.cfg.Server.NAME + `",` +
	// 		`"environment":"` + app.cfg.Server.ENV + `",` +
	// 		`"type":"access",` +
	// 		`"bytes_out":${bytes_out}}` + "\n",
	// 	Output: app.logWriters.AccessLog,
	// }))

	// r.router.Logger.SetOutput(panicLog)

	app.router.Use(middleware.Recover())

}
