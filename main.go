package main

import (
	"fmt"
	"go-template/app"
	"go-template/config"
	"go-template/infra/db"
	"go-template/infra/logger"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.LoadDefault()

	logger.New(cfg.Server.NAME, cfg.Server.ENV)

	db := db.NewGoPG(cfg)

	// st, err := storage.New()
	// if err != nil {
	// 	log.Panic(err)
	// }

	// Init application
	application := app.New(cfg, echo.New(), db)

	// Start server
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.Server.PORT),
		Handler: application.BuildHandler(),
	}
	// logger.Infof("server %v is running at %v", Version, address)
	fmt.Printf("server is running at %v\n", cfg.Server.PORT)

	if err := server.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
